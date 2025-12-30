package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/mojocn/base64Captcha"

	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/features/auth/dto"
)

// CaptchaUseCase is a captcha use case.
type CaptchaUseCase struct {
	repo   dto.CaptchaRepo
	config *confpb.Captcha
	log    *log.Helper
}

// NewCaptchaUseCase new a captcha use case.
func NewCaptchaUseCase(repo dto.CaptchaRepo, c *confpb.Captcha, logger log.Logger) *CaptchaUseCase {
	return &CaptchaUseCase{
		repo:   repo,
		config: c,
		log:    log.NewHelper(logger),
	}
}

// GenerateCaptcha generates a new captcha.
func (uc *CaptchaUseCase) GenerateCaptcha(ctx context.Context) (id, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(
		int(uc.config.GetHeight()),
		int(uc.config.GetWidth()),
		int(uc.config.GetLength()),
		float64(uc.config.GetMaxskew()),
		uc.config.GetDotCount(),
	)
	c := base64Captcha.NewCaptcha(driver, uc.repo)
	id, content, err := c.Generate()
	return id, content, err
}
