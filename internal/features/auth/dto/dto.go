/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the auth module.
package dto

import (
	"origadmin/application/admin/internal/data/entity/ent/resource"
	"origadmin/application/admin/internal/data/entity/ent/user"
	"origadmin/application/admin/internal/data/entity/ent/view"
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

// ConvertInt32ToStatus is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertInt32ToStatus(from int32) resource.Status {
	switch from {
	case 1:
		return resource.StatusEnabled
	default:
		return resource.StatusDisabled
	}
}

// ConvertStatusToInt32 is a custom conversion function stub.
// Please implement this function to complete the conversion.
func ConvertStatusToInt32(from resource.Status) int32 {
	switch from {
	case resource.StatusEnabled:
		return 1
	default:
		return 0
	}
}
