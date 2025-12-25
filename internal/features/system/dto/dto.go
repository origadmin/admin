package dto

import (
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/enums"
)

//go:generate abgen -debug .

//go:abgen:package:path=origadmin/application/admin/internal/data/entity/ent,alias=ent
//go:abgen:package:path=origadmin/application/admin/api/v1/services/types,alias=types
//go:abgen:pair:packages="ent,types"
//go:abgen:convert:direction="both"
//go:abgen:convert:source:suffix=""
//go:abgen:convert:target:suffix="PB"

// ConvertGenderToString is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertGenderToString(from user.Gender) string {
	switch from {
	case user.GenderFemale:
		return "female"
	default:
		return "male"
	}
}

// ConvertStringToGender is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertStringToGender(from string) user.Gender {
	switch from {
	case "female":
		return user.GenderFemale
	default:
		return user.GenderMale
	}
}
