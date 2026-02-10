package dto

import (
	"encoding/json"

	"origadmin/application/admin/internal/data/entity/ent/userprofile"
	"origadmin/application/admin/internal/data/entity/ent/view"
)

//go:generate abgen -debug .

//go:abgen:package:path=origadmin/application/admin/internal/data/entity/ent,alias=ent
//go:abgen:package:path=origadmin/application/admin/api/v1/services/types,alias=types
//go:abgen:pair:packages="ent,types"
//go:abgen:convert:direction="both"
//go:abgen:convert:source:suffix=""
//go:abgen:convert:target:suffix="PB"

const (
	StatusEnabled  = 1
	StatusDisabled = 0
)

// ConvertStringToStringMapToString is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertStringToStringMapToString(from map[string]string) string {
	bytes, err := json.Marshal(from)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// ConvertStringToStringToStringMap is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertStringToStringToStringMap(from string) map[string]string {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(from), &m)
	if err != nil {
		return nil
	}
	return m
}

// ConvertGenderToString is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertGenderToString(from userprofile.Gender) string {
	switch from {
	case userprofile.GenderFemale:
		return "female"
	default:
		return "male"
	}
}

// ConvertStringToGender is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertStringToGender(from string) userprofile.Gender {
	switch from {
	case "female":
		return userprofile.GenderFemale
	default:
		return userprofile.GenderMale
	}
}

// ConvertStringToType is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertStringToType(from string) view.Type {
	return ViewTypeCode(from)
}

// ConvertTypeToString is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertTypeToString(from view.Type) string {
	return ViewTypeName(from)
}

// NonNilStr returns the value of the string pointer if it is not nil, otherwise returns an empty string.
func NonNilStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// StrPtr returns a pointer to the given string.
func StrPtr(s string) *string {
	return &s
}
