package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"golang.org/x/exp/constraints" // for Integer
)

type IDType interface {
	constraints.Integer | string
}

type GenericID[T IDType] struct {
	mixin.Schema
	Key                  string
	CommentKey           string
	Optional             bool
	Positive             bool
	Unique               bool
	Immutable            bool
	UseDefault           bool
	DefaultFunc          func() int64
	UseCustomIDGenerator bool
}

func (g GenericID[T]) Fields() []ent.Field {
	var pkField ent.Field
	var t T

	switch any(t).(type) {
	case int64, int, int32:
		pkField = field.Int64("id")
	case string:
		pkField = field.String("id")
	}
	return []ent.Field{pkField}
}

// ... 实现其他泛型方法 FK, OptionalFK 等

// 在 field.go 中可以这样使用
// var innerID = GenericID[int64]{} // 使用int64
// or
// var innerID = GenericID[string]{} // 使用string
