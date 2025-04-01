/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/types"

	"origadmin/application/admin/internal/configs"
)

func Setup(bootstrap *configs.Bootstrap) error {
	// add init action here
	crypto := bootstrap.GetCryptoType()
	cryptoType := types.TypeArgon2
	if crypto != "" {
		cryptoType = types.ParseType(crypto)
	}
	err := hash.UseCrypto(cryptoType)
	if err != nil {
		return err
	}
	return nil
}
