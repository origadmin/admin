/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and interfaces for the module.
package dto

// ResourceType 资源类型
type ResourceType int8

const (
	ResourceTypeMenu   ResourceType = 1 // 目录
	ResourceTypePage   ResourceType = 2 // 页面
	ResourceTypeButton ResourceType = 3 // 按钮
	ResourceTypeAPI    ResourceType = 4 // API接口
)

// String 返回资源类型的字符表示
func (t ResourceType) String() string {
	return string(t.Code())
}

// Code 返回资源类型的字符码
func (t ResourceType) Code() rune {
	switch t {
	case ResourceTypeMenu:
		return 'M'
	case ResourceTypePage:
		return 'P'
	case ResourceTypeButton:
		return 'B'
	case ResourceTypeAPI:
		return 'A'
	default:
		return 'U' // Unknown
	}
}

// Name 返回资源类型的名称
func (t ResourceType) Name() string {
	switch t {
	case ResourceTypeMenu:
		return "menu"
	case ResourceTypePage:
		return "page"
	case ResourceTypeButton:
		return "button"
	case ResourceTypeAPI:
		return "api"
	default:
		return "unknown"
	}
}
