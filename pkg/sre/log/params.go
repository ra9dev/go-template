package log

import "github.com/ra9dev/go-template/pkg/sre"

type Params struct {
	Env   sre.Env
	Level Level
}

func NewParams(env sre.Env, lvl Level) Params {
	return Params{
		Env:   env,
		Level: lvl,
	}
}

func (p Params) withDefault() Params {
	if p.Env == "" {
		p.Env = sre.EnvProduction
	}

	if p.Level == "" {
		p.Level = DebugLevel
	}

	return p
}
