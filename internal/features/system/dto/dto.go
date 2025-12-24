package dto

import (
	"origadmin/application/admin/internal/data/entity/ent/user"
)

//go:generate go run github.com/origadmin/abgen/cmd/abgen -debug go run ./cmd/abgen -debug .

//go:abgen:package:path=origadmin/application/admin/internal/features/system/data/ent,alias=ent
//go:abgen:package:path=origadmin/application/admin/api/v1/services/types,alias=types
//go:abgen:pair:packages="ent,types"
//go:abgen:convert:direction="both"
//go:abgen:convert:source:suffix=""
//go:abgen:convert:target:suffix="PB"

// ConvertGenderToString is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertGenderToString(from user.Gender) string {
	// TODO: Implement this custom conversion
	panic("stub! not implemented")
}

// ConvertStringToGender is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertStringToGender(from string) user.Gender {
	// TODO: Implement this custom conversion
	panic("stub! not implemented")
}

type Pagination struct {
	Page     int
	PageSize int
}

type QueryOption struct {
	Pagination *Pagination
	OrderBy    []string
}
