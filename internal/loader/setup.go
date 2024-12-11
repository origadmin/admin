/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/errors"

	"origadmin/application/admin/internal/configs"
)

func Setup(bootstrap *configs.Bootstrap) error {
	// add init action here
	crypto := bootstrap.GetCryptoType()
	switch hash.Type(crypto) {
	case hash.TypeArgon2:
		hash.UseArgon2()
	case hash.TypeMD5:
		hash.UseMD5()
	case hash.TypeSHA1:
		hash.UseSHA1()
	case hash.TypeSHA256:
		hash.UseSHA256()
	case hash.TypeScrypt:
		hash.UseScrypt()
	case hash.TypeHMAC256:

	default:
		return errors.New(fmt.Sprintf("unsupported crypto type: %s", crypto))
	}

	return nil
}
