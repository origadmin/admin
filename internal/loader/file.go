/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"os"
	"strings"

	"github.com/goexts/generic/settings"
	"github.com/origadmin/toolkits/codec"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// SaveOption represents an option for saving configuration data.
type SaveOption = func(*protojson.MarshalOptions)

// SaveConfig saves the configuration data to the specified file path.
func SaveConfig(path string, data any, opts ...SaveOption) error {
	if v, ok := data.(proto.Message); ok && strings.HasSuffix(path, ".json") {
		opt := settings.Apply(&protojson.MarshalOptions{
			Indent: " ",
		}, opts)
		bytes, err := opt.Marshal(v)
		if err != nil {
			return err
		}
		if err := os.WriteFile(path, bytes, 0644); err != nil {
			return err
		}
		return nil
	}
	if err := codec.EncodeToFile(path, data); err != nil {
		return err
	}
	return nil
}
