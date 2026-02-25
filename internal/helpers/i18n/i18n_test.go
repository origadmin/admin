/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package i18n implements the functions, types, and contracts for the module.
package i18n

import (
	"testing"
)

func TestLocaleText(t *testing.T) {
	type args struct {
		locale string
		key    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				locale: "en_US",
				key:    "test",
			},
			want: "en_US test",
		},
		{
			name: "test2",
			args: args{
				locale: "zh_CN",
				key:    "test",
			},
			want: "中文测试",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TextFor(tt.args.locale, tt.args.key); got != tt.want {
				t.Errorf("LocaleText() = %v, want %v", got, tt.want)
			}
		})
	}
}
