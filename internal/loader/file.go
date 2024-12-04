/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package loader implements the functions, types, and interfaces for the module.
package loader

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/goexts/generic/settings"
	"github.com/origadmin/contrib/replacer"
	"github.com/origadmin/toolkits/codec"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// SaveOption represents an option for saving configuration data.
type SaveOption = func(*protojson.MarshalOptions)

var (
	r = replacer.New(replacer.WithStart("${"), replacer.WithEnd("}"), replacer.WithSeparator(":"))
)

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

func Replace(s []byte, envs map[string]string) []byte {
	return r.Replace(s, envs)
}

func ReplaceObject(s any, envs map[string]string) error {
	marshal, err := json.Marshal(s)
	if err != nil {
		return err
	}
	marshal = Replace(marshal, envs)
	return json.Unmarshal(marshal, s)
}
