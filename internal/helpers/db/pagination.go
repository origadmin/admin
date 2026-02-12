/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package db provides common, generic helpers for database operations.
package db

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/origadmin/runtime/errors"
	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/helpers/repo"
)

const (
	PagingModeCursor = "cursor"
	PagingModeOffset = "offset"
	PagingModeNone   = "none"
)

// Identifiable defines the contract for any entity that has a retrievable ID.
type Identifiable interface {
	GetId() int64
}

// TimeTracker defines any type that has GetCreateTime and GetUpdateTime methods.
type TimeTracker interface {
	GetCreateTime() *timestamppb.Timestamp
	GetUpdateTime() *timestamppb.Timestamp
}

// Sequencer defines any type that has a GetSequence method.
type Sequencer interface {
	GetSequence() int32
}

// GetPageSize returns the normalized page size for a query.
func GetPageSize(req interface{}) int {
	opt := repo.QueryOptionFromRequest(req)
	normalizedOpt := normalizeQueryOption(&opt)
	return normalizedOpt.PageSize
}

// CalculatePagination computes all pagination-related values for a list response.
// It returns all necessary fields (page, pageSize, token) in a single call,
// allowing the service layer to construct the final response without redundant logic.
func CalculatePagination[T any](results []T, queryOpt *repo.QueryOption) (page int32, pageSize int32, token string, err error) {
	normalizedOpt := normalizeQueryOption(queryOpt)
	pageSize = int32(normalizedOpt.PageSize)
	token = ""
	page = 0

	if queryOpt.PagingMode == PagingModeCursor {
		// Only generate a next page token if the number of results equals the page size,
		// which implies there might be more data.
		if len(results) > 0 && len(results) == int(pageSize) {
			nextToken, err := generateNextPageToken(results, queryOpt)
			if err != nil {
				return 0, 0, "", errors.InternalServer("TOKEN_GENERATION_FAILED", err.Error())
			}
			token = nextToken
		}
	} else {
		page = int32(normalizedOpt.Page)
	}
	return page, pageSize, token, nil
}

// generateNextPageToken creates a compact, Protobuf-based pagination token.
// It uses type assertions on interfaces to extract field values, avoiding reflection.
func generateNextPageToken[T any](results []T, opt *repo.QueryOption) (string, error) {
	if len(results) == 0 {
		return "", nil // No results, no next page token
	}

	lastItem := any(results[len(results)-1]) // Convert to any for type assertion
	var sortValues []*types.SortValue

	orderByClauses := opt.OrderBy
	if len(orderByClauses) == 0 {
		orderByClauses = []string{"id,desc"} // Default stable order
	}

	for _, orderByClause := range orderByClauses {
		field, isDesc := parseOrderByClause(orderByClause)
		fieldEnum := MapFieldToEnum(field)
		var valueStr string
		var err error

		switch fieldEnum {
		case types.SortField_SORT_FIELD_ID:
			if item, ok := lastItem.(Identifiable); ok {
				valueStr = fmt.Sprintf("%d", item.GetId())
			} else {
				err = fmt.Errorf("sort field 'id' used on a type that does not implement db.Identifiable")
			}
		case types.SortField_SORT_FIELD_CREATE_TIME:
			if item, ok := lastItem.(TimeTracker); ok {
				if t := item.GetCreateTime(); t != nil {
					valueStr = fmt.Sprintf("%d", t.AsTime().UnixNano())
				} else {
					valueStr = "0"
				}
			} else {
				err = fmt.Errorf("sort field 'create_time' used on a type that does not implement db.TimeTracker")
			}
		case types.SortField_SORT_FIELD_UPDATE_TIME:
			if item, ok := lastItem.(TimeTracker); ok {
				if t := item.GetUpdateTime(); t != nil {
					valueStr = fmt.Sprintf("%d", t.AsTime().UnixNano())
				} else {
					valueStr = "0"
				}
			} else {
				err = fmt.Errorf("sort field 'update_time' used on a type that does not implement db.TimeTracker")
			}
		case types.SortField_SORT_FIELD_SEQUENCE:
			if item, ok := lastItem.(Sequencer); ok {
				valueStr = fmt.Sprintf("%d", item.GetSequence())
			} else {
				err = fmt.Errorf("sort field 'sequence' used on a type that does not implement db.Sequencer")
			}
		default:
			err = fmt.Errorf("unsupported sort field enum: %v", fieldEnum)
		}

		if err != nil {
			return "", err
		}
		sortValues = append(sortValues, &types.SortValue{
			Value:  valueStr,
			Field:  fieldEnum,
			IsDesc: isDesc,
		})
	}

	cursor := &types.PageCursor{SortValues: sortValues}
	return EncodeCursor(cursor)
}

// EncodeCursor uses proto.Marshal for compact encoding.
func EncodeCursor(c *types.PageCursor) (string, error) {
	data, err := proto.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("proto marshal cursor: %w", err)
	}
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data), nil
}

// DecodeCursor uses proto.Unmarshal.
func DecodeCursor(token string) (*types.PageCursor, error) {
	var c types.PageCursor
	data, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(token)
	if err != nil {
		return nil, fmt.Errorf("base64 decode token: %w", err)
	}
	if err := proto.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("proto unmarshal cursor: %w", err)
	}
	return &c, nil
}

// parseOrderByClause parses a string like "field,desc" into field name and isDesc.
func parseOrderByClause(orderByClause string) (field string, isDesc bool) {
	parts := strings.Split(orderByClause, ",")
	field = parts[0]
	isDesc = true // Default to DESC
	if len(parts) > 1 && strings.EqualFold(parts[1], "asc") {
		isDesc = false
	}
	return
}

// MapFieldToEnum converts a field name string to its SortField enum.
func MapFieldToEnum(field string) types.SortField {
	switch strings.ToLower(field) {
	case "id":
		return types.SortField_SORT_FIELD_ID
	case "create_time":
		return types.SortField_SORT_FIELD_CREATE_TIME
	case "update_time":
		return types.SortField_SORT_FIELD_UPDATE_TIME
	case "sequence":
		return types.SortField_SORT_FIELD_SEQUENCE
	default:
		return types.SortField_SORT_FIELD_DEFAULT_UNSPECIFIED
	}
}

// MapEnumToField converts a SortField enum to its corresponding database column name.
func MapEnumToField(fieldEnum types.SortField) string {
	switch fieldEnum {
	case types.SortField_SORT_FIELD_ID:
		return "id"
	case types.SortField_SORT_FIELD_CREATE_TIME:
		return "create_time"
	case types.SortField_SORT_FIELD_UPDATE_TIME:
		return "update_time"
	case types.SortField_SORT_FIELD_SEQUENCE:
		return "sequence"
	default:
		return "id" // Safe fallback
	}
}

// BuildCursorWhere is a generic predicate builder for cursor-based pagination.
func BuildCursorWhere[P ~func(*sql.Selector)](cursor *Cursor) P {
	return func(s *sql.Selector) {
		sortValues := cursor.GetSortValues()
		if len(sortValues) == 0 {
			return
		}

		var orPredicates []*sql.Predicate
		for i := 0; i < len(sortValues); i++ {
			var andPredicates []*sql.Predicate
			for j := 0; j < i; j++ {
				sv := sortValues[j]
				andPredicates = append(andPredicates, sql.EQ(MapEnumToField(sv.Field), sv.Value))
			}
			sv := sortValues[i]
			var inequality *sql.Predicate
			if sv.IsDesc {
				inequality = sql.LT(MapEnumToField(sv.Field), sv.Value)
			} else {
				inequality = sql.GT(MapEnumToField(sv.Field), sv.Value)
			}
			andPredicates = append(andPredicates, inequality)
			orPredicates = append(orPredicates, sql.And(andPredicates...))
		}
		if len(orPredicates) > 0 {
			s.Where(sql.Or(orPredicates...))
		}
	}
}

// Pageable, Selector, etc. interfaces remain the same.
type Pageable[T any, W any, O any, R any] interface {
	Counter[T]
	Cloner[T]
	Filterer[T, W]
	Orderer[T, O]
	Fetcher[R]
	Limit(int) T
	Offset(int) T
}
type Selector[S any] interface{ Select(fields ...string) S }
type SourceSelector interface{ ~func(*sql.Selector) }
type Orderer[T any, O any] interface{ Order(...O) T }
type Filterer[T any, W any] interface{ Where(...W) T }
type Cloner[T any] interface{ Clone() T }
type Counter[T any] interface {
	Cloner[T]
	Count(ctx context.Context) (int, error)
}
type Fetcher[T any] interface {
	All(ctx context.Context) ([]T, error)
	Only(ctx context.Context) (T, error)
}

func normalizeQueryOption(opt *repo.QueryOption) *repo.QueryOption {
	if opt == nil {
		return &repo.QueryOption{Page: 1, PageSize: repo.DefaultPageSize}
	}
	if opt.Page <= 0 {
		opt.Page = 1
	}
	if opt.PagingMode == PagingModeNone {
		if opt.PageSize <= 0 || opt.PageSize > repo.HardLimit {
			opt.PageSize = repo.HardLimit
		}
	} else {
		if opt.PageSize <= 0 {
			opt.PageSize = repo.DefaultPageSize
		}
		if opt.PageSize > repo.MaxPageSize {
			opt.PageSize = repo.MaxPageSize
		}
	}
	return opt
}

// Cursor is an alias for the Protobuf-defined PageCursor.
type Cursor = types.PageCursor

// Paginate applies pagination and sorting logic to a query.
func Paginate[R any, W ~func(*sql.Selector), O SourceSelector, P Pageable[P, W, O, R]](query P, opt *repo.QueryOption) P {
	if opt == nil {
		opt = &repo.QueryOption{}
	}
	opt = normalizeQueryOption(opt)
	query = query.Limit(opt.PageSize)

	pagingMode := opt.PagingMode
	if pagingMode == "" {
		pagingMode = PagingModeOffset
	}

	switch pagingMode {
	case PagingModeCursor:
		if opt.PageToken != "" {
			cursor, err := DecodeCursor(opt.PageToken)
			if err != nil || len(cursor.GetSortValues()) == 0 {
				return query.Limit(0) // Invalid token
			}

			var orders []O
			for _, sv := range cursor.GetSortValues() {
				orders = append(orders, OrderByField[O](MapEnumToField(sv.Field), sv.IsDesc))
			}
			query = query.Order(orders...)

			whereClause := BuildCursorWhere[W](cursor)
			query = query.Where(whereClause)

		} else {
			orderBy := opt.OrderBy
			if len(orderBy) == 0 {
				orderBy = []string{"id,desc"}
			}
			query = query.Order(OrderBy[O](orderBy)...)
		}
	case PagingModeOffset:
		if len(opt.OrderBy) > 0 {
			query = query.Order(OrderBy[O](opt.OrderBy)...)
		}
		query = query.Offset((opt.Page - 1) * opt.PageSize)
	case PagingModeNone:
		if len(opt.OrderBy) > 0 {
			query = query.Order(OrderBy[O](opt.OrderBy)...)
		}
	}
	return query
}

// CountTotal remains the same.
func CountTotal[Q Counter[Q]](ctx context.Context, query Q) (int32, error) {
	count, err := query.Clone().Count(ctx)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

// Find executes the query with pagination and returns results.
func Find[R any, W, O SourceSelector, P Pageable[P, W, O, R]](ctx context.Context, query P, o *repo.QueryOption) ([]R, int32, error) {
	var count int32
	var err error
	if o.PagingMode != PagingModeCursor {
		count, err = CountTotal(ctx, query)
		if err != nil {
			return nil, 0, fmt.Errorf("count query failed: %w", err)
		}
		if count == 0 {
			return []R{}, 0, nil
		}
	}
	if !o.SortFromToken {
		if len(o.OrderBy) > 0 {
			orders := OrderBy[O](o.OrderBy)
			if len(orders) > 0 {
				query.Order(orders...)
			}
		}
	}
	query = Paginate[R, W, O, P](query, o)
	result, err := query.All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("data query failed: %w", err)
	}
	return result, count, nil
}
