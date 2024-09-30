package config

type Settings struct {
	ServiceName string `default:"KasaAdmin"`
	Version     string `default:"v1.0.0"`

	PprofAddr      string
	DisableStatic  bool
	DisableSwagger bool

	CryptoType      string `default:"argon2"`
	DefaultLoginPwd string `default:"06f684620c2e8f7caf9bb5a4fcba2ff2"` // CryptoType(KasaAdmin@123456)
	MenuFile        string // From schema.Menus (JSON/YAML/YML/TOML)
	LockMenuTables  bool
	HTTP            struct {
		Addr            string `default:":28080"`
		ShutdownTimeout int    `default:"10"` // unit:seconds
		ReadTimeout     int    `default:"30"` // unit:seconds
		WriteTimeout    int    `default:"60"` // unit:seconds
		IdleTimeout     int    `default:"10"` // unit:seconds
		UseTLS          bool
		CertFile        string
		KeyFile         string
	}
	Root struct {
		Random   bool
		ID       string `default:"root"`
		Username string `default:"kasaadmin"`
		Password string
		Name     string `default:"KasaAdmin"`
	}
}
