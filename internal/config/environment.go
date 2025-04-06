package config

import "fmt"

type Environment string

const (
	EnvironmentLocal Environment = "local"
	EnvironmentProd  Environment = "prod"
)

func ShowEnvironments() string {
	return fmt.Sprintf("%s, %s", EnvironmentLocal, EnvironmentProd)
}

func ToEnvironment(value string) Environment {
	switch value {
	case string(EnvironmentLocal):
		return EnvironmentLocal
	case string(EnvironmentProd):
		return EnvironmentProd
	default:
		return EnvironmentLocal
	}
}
