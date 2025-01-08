/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package helpers implements the functions, types, and interfaces for the module.
package base64image

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestBase64Image(t *testing.T) {
	type args struct {
		buffer *bytes.Buffer
		typ    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				buffer: bytes.NewBufferString("test"),
				typ:    filepath.Ext("test.png"),
			},
			want: "data:image/png;base64,dGVzdA==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ImageFromBuffer(tt.args.buffer, tt.args.typ); got != tt.want {
				t.Errorf("Base64Image() = %v, want %v", got, tt.want)
			}
		})
	}
}
