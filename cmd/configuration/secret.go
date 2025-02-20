package configuration

import (
	craveConfiguration "crave/shared/configuration"
)

type Secret struct {
	Controller string
}

type Dependency struct {
}

type Variable struct {
	Databese   *craveConfiguration.Database
	Secret     *Secret
	Dependency *Dependency
	HubApiIp   string
	HubApiPort uint16
}

func NewVariable() *Variable {
	return &Variable{

		Secret: &Secret{
			Controller: "secret variable",
		},
		HubApiIp:   "localhost",
		HubApiPort: 17000,
	}
}
