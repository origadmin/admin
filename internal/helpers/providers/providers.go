package providers

import (
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	authnv1 "github.com/origadmin/contrib/api/gen/go/security/authn/v1"
	"github.com/origadmin/contrib/security/authn/jwt"
	"github.com/origadmin/contrib/security/credential"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/container"
	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"origadmin/application/admin/internal/conf"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/captcha"
)

// ProvideAuthenticatorOptions creates the JWT options from the application configuration.
func ProvideAuthenticatorOptions(c *conf.Config) (*jwt.Options, error) {
	securityConfig := c.GetBootstrap().GetSecurity()
	if securityConfig == nil {
		return nil, errors.New("security configuration not found in bootstrap config")
	}
	authnConfig := securityConfig.GetAuthn()
	if authnConfig == nil {
		return nil, errors.New("authn configuration not found in security config")
	}
	configs := authnConfig.GetConfigs()

	var jwtConfig *authnv1.Authenticator
	for _, mw := range configs {
		if mw.GetType() == "jwt" {
			jwtConfig = mw
			break
		}
	}
	if jwtConfig == nil || jwtConfig.GetJwt() == nil {
		return nil, errors.New("JWT authenticator configuration not found in bootstrap config")
	}
	return jwt.NewOptions(jwtConfig)
}

// ProvideCredentialCreator creates the JWT authenticator instance.
// It returns a credential.Creator interface, which is implemented by *jwt.Authenticator.
func ProvideCredentialCreator(opts *jwt.Options, logger log.Logger) (credential.Creator, error) {
	return jwt.New(opts, logger)
}

// ProvideAuthenticator creates the JWT authenticator instance.
// It returns a *jwt.Authenticator, which is needed by the AuthService.
func ProvideAuthenticator(opts *jwt.Options, logger log.Logger) (*jwt.Authenticator, error) {
	return jwt.New(opts, logger)
}

func ProvideCache(r *runtime.App) (container.CacheProvider, error) {
	cacheProvider, err := r.CacheProvider()
	if err != nil {
		return nil, err
	}
	return cacheProvider, nil
}

func ProvideCaptcha(p container.CacheProvider, cfg *confpb.Captcha) (*captcha.Captcha, error) {
	if cfg == nil {
		cfg = &confpb.Captcha{
			CacheName: "default",
			Height:    80,
			Width:     240,
			Length:    6,
			Maxskew:   0.7,
			DotCount:  80,
		}
	}
	cache, err := p.Cache(cfg.CacheName)
	if err != nil {
		return nil, err
	}

	c := &captcha.Config{
		Store: captcha.NewStore(cache),
		DriverDigit: &captcha.DriverDigit{
			Height:   int(cfg.GetHeight()),
			Width:    int(cfg.GetWidth()),
			Length:   int(cfg.GetLength()),
			MaxSkew:  float64(cfg.GetMaxskew()),
			DotCount: int(cfg.GetDotCount()),
		},
	}
	return captcha.NewCaptcha(c), nil
}

func ProvideHasher() (hash.Crypto, error) {
	return hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(bcrypt.DefaultCost))
}

func ProvideLogger(app *runtime.App) log.Logger {
	return app.Logger()
}

var ProviderSet = wire.NewSet(
	// Instructions for wire to extract nested configs
	wire.FieldsOf(new(*conf.Config), "Bootstrap"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Servers"),
	wire.FieldsOf(new(*confpb.Bootstrap), "Captcha"),
	ProvideLogger,
	ProvideCache,
	ProvideAuthenticatorOptions,
	ProvideCredentialCreator,
	ProvideAuthenticator,
	ProvideCaptcha,
	ProvideHasher,
)
