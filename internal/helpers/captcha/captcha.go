/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package captcha implements the functions, types, and contracts for the module.
package captcha

import (
	"fmt"
	"net/http"

	"github.com/mojocn/base64Captcha"

	"github.com/origadmin/runtime/errors"

	typespb "origadmin/application/admin/api/v1/services/types"
	confpb "origadmin/application/admin/internal/conf/pb"
)

var (
	// ErrNotFound is returned when a captcha ID is not found in the store.
	ErrNotFound = errors.New(http.StatusBadRequest, typespb.AuthErrorReason_AUTH_ERROR_REASON_CAPTCHA_NOT_FOUND.String(), "captcha not found")
)

// Constants for different captcha types.
const (
	TypeAudio   = "audio"
	TypeString  = "string"
	TypeDigit   = "digit"
	TypeChinese = "chinese"
	TypeMath    = "math"
)

// Constants for MIME types.
const (
	MimeTypeAudio = base64Captcha.MimeTypeAudio
	MimeTypeImage = base64Captcha.MimeTypeImage
)

// Type aliases for base64Captcha types.
type (
	Store         = base64Captcha.Store
	Driver        = base64Captcha.Driver
	DriverAudio   = base64Captcha.DriverAudio
	DriverString  = base64Captcha.DriverString
	DriverChinese = base64Captcha.DriverChinese
	DriverMath    = base64Captcha.DriverMath
	DriverDigit   = base64Captcha.DriverDigit
)

// Captcha provides an interface for generating and verifying captchas.
type Captcha struct {
	store    Store
	captchas map[string]*base64Captcha.Captcha
	drivers  map[string]base64Captcha.Driver
	config   *confpb.Captcha
}

// Config holds the configuration for the captcha service.
type Config struct {
	Store         Store
	DriverAudio   *DriverAudio
	DriverString  *DriverString
	DriverChinese *DriverChinese
	DriverMath    *DriverMath
	DriverDigit   *DriverDigit
	Captcha       *confpb.Captcha
}

// NewCaptcha creates a new Captcha instance with the given configuration.
// It initializes default drivers for various captcha types if they are not provided.
func NewCaptcha(config *Config) *Captcha {
	if config == nil {
		config = &Config{}
	}
	if config.Store == nil {
		config.Store = base64Captcha.DefaultMemStore
	}
	if config.Captcha == nil {
		config.Captcha = &confpb.Captcha{}
	}

	drivers := make(map[string]base64Captcha.Driver)
	captchas := make(map[string]*base64Captcha.Captcha)

	// Initialize Digit Driver
	if config.DriverDigit == nil {
		config.DriverDigit = &base64Captcha.DriverDigit{
			Height:   int(config.Captcha.GetHeight()),
			Width:    int(config.Captcha.GetWidth()),
			Length:   int(config.Captcha.GetLength()),
			MaxSkew:  float64(config.Captcha.GetMaxSkew()),
			DotCount: int(config.Captcha.GetDotCount()),
		}
	}
	drivers[TypeDigit] = config.DriverDigit
	captchas[TypeDigit] = base64Captcha.NewCaptcha(config.DriverDigit, config.Store)

	// Initialize String Driver
	if config.DriverString == nil {
		config.DriverString = &base64Captcha.DriverString{
			Height:          int(config.Captcha.GetHeight()),
			Width:           int(config.Captcha.GetWidth()),
			NoiseCount:      int(config.Captcha.GetDotCount()),
			ShowLineOptions: base64Captcha.OptionShowHollowLine,
			Length:          int(config.Captcha.GetLength()),
			Source:          "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		}
	}
	drivers[TypeString] = config.DriverString
	captchas[TypeString] = base64Captcha.NewCaptcha(config.DriverString, config.Store)

	// Initialize Chinese Driver
	if config.DriverChinese == nil {
		config.DriverChinese = &base64Captcha.DriverChinese{
			Height:          int(config.Captcha.GetHeight()),
			Width:           int(config.Captcha.GetWidth()),
			NoiseCount:      int(config.Captcha.GetDotCount()),
			ShowLineOptions: base64Captcha.OptionShowSlimeLine,
			Length:          int(config.Captcha.GetLength()),
			Source:          "的一是在不了有和人这中大为上个国我以要他时来用们生到作地于出就分对成会可主发年动同工也能下过子说产种面而方后多定行学法所民得经十三之进着等",
		}
	}
	drivers[TypeChinese] = config.DriverChinese
	captchas[TypeChinese] = base64Captcha.NewCaptcha(config.DriverChinese, config.Store)

	// Initialize Math Driver
	if config.DriverMath == nil {
		config.DriverMath = &base64Captcha.DriverMath{
			Height:          int(config.Captcha.GetHeight()),
			Width:           int(config.Captcha.GetWidth()),
			NoiseCount:      int(config.Captcha.GetDotCount()),
			ShowLineOptions: base64Captcha.OptionShowSineLine,
		}
	}
	drivers[TypeMath] = config.DriverMath
	captchas[TypeMath] = base64Captcha.NewCaptcha(config.DriverMath, config.Store)

	// Initialize Audio Driver
	if config.DriverAudio == nil {
		config.DriverAudio = base64Captcha.DefaultDriverAudio
	}
	drivers[TypeAudio] = config.DriverAudio
	captchas[TypeAudio] = base64Captcha.NewCaptcha(config.DriverAudio, config.Store)

	return &Captcha{
		store:    config.Store,
		captchas: captchas,
		drivers:  drivers,
		config:   config.Captcha,
	}
}

// Generate creates a new captcha of the specified type.
// It returns the captcha ID, the base64 encoded image or audio, the answer, and an error if any.
func (c *Captcha) Generate(captchaType string) (id, b64s, answer string, err error) {
	captcha, ok := c.captchas[captchaType]
	if !ok {
		return "", "", "", errors.New(http.StatusBadRequest, "INVALID_CAPTCHA_TYPE", fmt.Sprintf("invalid captcha type: %s", captchaType))
	}
	return captcha.Generate()
}

// GetAudioForID generates an audio representation for an existing captcha ID.
// This is useful for accessibility purposes, allowing users to listen to a captcha.
func (c *Captcha) GetAudioForID(id string) (b64s string, err error) {
	// Get the answer from the store without clearing it.
	answer := c.store.Get(id, false)
	if answer == "" {
		return "", ErrNotFound
	}

	// Get the audio driver.
	audioDriver, ok := c.drivers[TypeAudio].(*DriverAudio)
	if !ok {
		// This should not happen if the captcha service is configured correctly.
		return "", errors.New(http.StatusInternalServerError, "AUDIO_DRIVER_NOT_CONFIGURED", "audio driver is not configured")
	}

	// Generate audio content using the retrieved answer.
	audio, err := audioDriver.DrawCaptcha(answer)
	if err != nil {
		return "", err
	}

	return audio.EncodeB64string(), nil
}

// Verify checks if the provided answer for a given captcha ID is correct.
// The `clear` parameter determines whether to remove the captcha from the store after verification.
func (c *Captcha) Verify(id, answer string, clear bool) bool {
	// If captcha is not enabled in the config, always return true.
	if !c.config.GetEnabled() {
		return true
	}
	return c.store.Verify(id, answer, clear)
}
