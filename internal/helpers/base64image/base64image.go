/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package base64image implements the functions, types, and interfaces for the module.
package base64image

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Image(buffer []byte, typ string) string {
	return fmt.Sprintf("data:%s;base64,%s", imageType(typ), base64.StdEncoding.EncodeToString(buffer))
}

func Buffer(buffer string) (string, []byte, error) {
	// Split strings to extract MIME types and Base64 data
	parts := strings.SplitN(buffer, ";base64,", 2)
	if len(parts) != 2 {
		return "", nil, fmt.Errorf("invalid data URL format")
	}
	// Extract mime types
	imageType := strings.TrimPrefix(parts[0], "data:")

	// Decode base64 data
	bs, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", nil, fmt.Errorf("failed to decode base64: %v", err)
	}
	return imageType, bs, nil
}

func ImageFromBuffer(buffer *bytes.Buffer, typ string) string {
	return Image(buffer.Bytes(), typ)
}

func ImageFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return "", err
	}
	return Image(buffer.Bytes(), filepath.Ext(filename)), nil
}

func imageType(typ string) string {
	switch typ {
	case "png":
		return "image/png"
	case "jpg":
		return "image/jpeg"
	case "jpeg":
		return "image/jpeg"
	case "gif":
		return "image/gif"
	case "webp":
		return "image/webp"
	case "svg":
		return "image/svg+xml"
	case "bmp":
		return "image/bmp"
	case "tiff":
		return "image/tiff"
	}
	return "image/png"
}
