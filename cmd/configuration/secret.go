package configuration

import (
	craveConfiguration "crave/shared/configuration"
)

type Secret struct {
	Controller string
}

type Dependency struct {
	MinerGrpc *craveConfiguration.Api
}

type Miner struct {
	Api *craveConfiguration.Api
}

type Variable struct {
	Database   *craveConfiguration.Database
	Secret     *Secret
	Dependency *Dependency
	Api        *craveConfiguration.Api
	GrpcApi    *craveConfiguration.Api
}

func NewVariable() *Variable {
	return &Variable{
		Database: &craveConfiguration.Database{
			Uri:      "127.0.0.1:18000",
			Username: "root",
			Password: "root",
		},
		Secret: &Secret{
			Controller: "secretValue",
		},
		Dependency: &Dependency{
			MinerGrpc: &craveConfiguration.Api{
				Ip:   "localhost",
				Port: 3001,
			},
		},
		Api: &craveConfiguration.Api{
			Ip:   "localhost",
			Port: 3000,
		},
		GrpcApi: &craveConfiguration.Api{
			Ip:   "localhost",
			Port: 3002,
		},
	}
}
