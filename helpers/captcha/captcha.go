/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package captcha implements the functions, types, and interfaces for the module.
package captcha

import (
	"errors"
	"net/http"

	"github.com/mojocn/base64Captcha"
)

var ErrNotFound = errors.New("captcha not found")

const (
	TypeAudio   = "audio"
	TypeString  = "string"
	TypeDigit   = "digit"
	TypeChinese = "chinese"
	TypeMath    = "math"
)

const (
	MimeTypeAudio = base64Captcha.MimeTypeAudio
	MimeTypeImage = base64Captcha.MimeTypeImage
)

type (
	Store         = base64Captcha.Store
	Driver        = base64Captcha.Driver
	DriverAudio   = base64Captcha.DriverAudio
	DriverString  = base64Captcha.DriverString
	DriverChinese = base64Captcha.DriverChinese
	DriverMath    = base64Captcha.DriverMath
	DriverDigit   = base64Captcha.DriverDigit
)

//Captcha json request body.
type Captcha struct {
	Store         Store
	DriverAudio   *base64Captcha.Captcha
	DriverString  *base64Captcha.Captcha
	DriverChinese *base64Captcha.Captcha
	DriverMath    *base64Captcha.Captcha
	DriverDigit   *base64Captcha.Captcha
}

type Config struct {
	Store         Store
	DriverAudio   *DriverAudio
	DriverString  *DriverString
	DriverChinese *DriverChinese
	DriverMath    *DriverMath
	DriverDigit   *DriverDigit
}

func NewCaptcha(config *Config) *Captcha {
	if config == nil {
		config = &Config{}
	}
	if config.Store == nil {
		config.Store = base64Captcha.DefaultMemStore
	}
	if config.DriverAudio == nil {
		config.DriverAudio = base64Captcha.DefaultDriverAudio
	}
	//if config.DriverString == nil {
	//	config.DriverString = base64Captcha.DefaultDriverString
	//}
	if config.DriverDigit == nil {
		config.DriverDigit = base64Captcha.DefaultDriverDigit
	}
	if config.DriverChinese == nil {
		//config.DriverChinese = base64Captcha.NewDriverChinese()
	}
	if config.DriverMath == nil {
		//config.DriverMath = base64Captcha.NewDriverMath()
	}
	return &Captcha{
		DriverAudio: base64Captcha.NewCaptcha(config.DriverAudio, config.Store),
		DriverDigit: base64Captcha.NewCaptcha(config.DriverDigit, config.Store),
		Store:       config.Store,
	}
}

func (c *Captcha) GenerateDigit() (string, string, string, error) {
	return c.DriverDigit.Generate()
}

func (c *Captcha) GenerateString() (string, string, string, error) {
	return c.DriverString.Generate()
}

func (c *Captcha) GenerateAudio() (string, string, string, error) {
	return c.DriverAudio.Generate()
}

func (c *Captcha) GenerateChinese() (string, string, string, error) {
	return c.DriverChinese.Generate()
}

func (c *Captcha) ServerHTTP(w http.ResponseWriter, r *http.Request) {

}

func (c *Captcha) Verify(id, answer string, clear bool) bool {
	return c.Store.Verify(id, answer, clear)
}

func (c *Captcha) Reload(typ string, id string) bool {
	if typ == TypeAudio {
		return c.reload(c.DriverAudio.Driver, id)
	}
	return c.reload(c.DriverDigit.Driver, id)
}

func (c *Captcha) reload(d base64Captcha.Driver, id string) bool {
	old := c.Store.Get(id, false)
	if old == "" {
		return false
	}
	_, _, old = d.GenerateIdQuestionAnswer()
	err := c.Store.Set(id, old)
	if err != nil {
		return false
	}
	return true
}
