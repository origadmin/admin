package providers

import (
	"github.com/google/wire"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/captcha"
)

// ProviderCommonSet provides common dependencies that are safe for all modules.
var ProviderCommonSet = wire.NewSet(
	ProvideLogger,
	ProvideCache,
	ProvideHasher,
	ProvideCaptcha,
	wire.FieldsOf(new(*conf.Config), "Bootstrap"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Auth"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Security"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Servers"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Captcha"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Brokers"),
)

// ProvideLogger provides a logger instance from the runtime App.
func ProvideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

// ProvideCache provides a cache provider from the runtime App.
func ProvideCache(app *runtime.App) (container.CacheProvider, error) {
	return app.CacheProvider()
}

// ProvideHasher provides a password hasher.
func ProvideHasher() (hash.Crypto, error) {
	return hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(bcrypt.DefaultCost))
}

// ProvideCaptcha provides a captcha instance.
func ProvideCaptcha(app *runtime.App, p container.CacheProvider, cfg *confpb.Bootstrap) (*captcha.Captcha, error) {
	captchaConfig := cfg.GetCaptcha()
	if captchaConfig == nil {
		captchaConfig = &confpb.Captcha{}
	}
	if captchaConfig.CacheName == "" {
		captchaConfig.CacheName = "default"
	}
	if captchaConfig.Height == 0 {
		captchaConfig.Height = 80
	}
	if captchaConfig.Width == 0 {
		captchaConfig.Width = 240
	}
	if captchaConfig.Length == 0 {
		captchaConfig.Length = 6
	}
	if captchaConfig.MaxSkew == 0 {
		captchaConfig.MaxSkew = 0.7
	}
	if captchaConfig.DotCount == 0 {
		captchaConfig.DotCount = 80
	}

	cache, err := p.Cache(captchaConfig.CacheName)
	if err != nil {
		return nil, err
	}

	c := &captcha.Config{
		Store:   captcha.NewStore(cache),
		Captcha: captchaConfig,
	}
	return captcha.NewCaptcha(c), nil
}
