package sre

const (
	EnvLocal      Env = "local"
	EnvDev        Env = "dev"
	EnvStaging    Env = "stg"
	EnvProduction Env = "prod"
)

type Env string

func (e Env) String() string {
	return string(e)
}
