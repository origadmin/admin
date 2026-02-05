package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"origadmin/application/admin/internal/helpers/captcha"
)

// CaptchaUseCase is a captcha use case.
type CaptchaUseCase struct {
	captcha *captcha.Captcha
	log     *log.Helper
}

// NewCaptchaUseCase new a captcha use case.
func NewCaptchaUseCase(c *captcha.Captcha, logger log.Logger) *CaptchaUseCase {
	return &CaptchaUseCase{
		captcha: c,
		log:     log.NewHelper(logger),
	}
}

// GenerateCaptcha generates a new captcha.
func (uc *CaptchaUseCase) GenerateCaptcha(ctx context.Context, captchaType string) (id, b64s, answer string, err error) {
	return uc.captcha.Generate(captchaType)
}

// GetCaptchaAudio generates audio for a given captcha ID.
func (uc *CaptchaUseCase) GetCaptchaAudio(ctx context.Context, id string) (string, error) {
	return uc.captcha.GetAudioForID(id)
}

// VerifyCaptcha verifies a user's answer for a given captcha ID.
func (uc *CaptchaUseCase) VerifyCaptcha(ctx context.Context, id, answer string) bool {
	return uc.captcha.Verify(id, answer, true)
}
