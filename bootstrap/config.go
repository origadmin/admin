package bootstrap

const (
	// Name is the name of the application
	Name = "backend"
)

type Config struct {
	Name   string
	Random bool
	// config read property
	WorkDir string
	Configs []string
}
