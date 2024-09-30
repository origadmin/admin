package bootstrap

const (
	// Name is the name of the application
	Name = "backend"
)

type Config struct {
	Name   string
	Random bool
	// config read property
	Dir     string
	Configs string
}
