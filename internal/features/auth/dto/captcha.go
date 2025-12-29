package dto

import "github.com/mojocn/base64Captcha"

// CaptchaRepo defines the data access methods for captcha.
// It embeds the base64Captcha.Store interface.
type CaptchaRepo interface {
	base64Captcha.Store
}
