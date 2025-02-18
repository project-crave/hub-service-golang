package configuration

type Database struct {
	Uri      string
	Username string
	Password string
}

type Secret struct {
	Controller string
}

type Dependency struct {
}

type Variable struct {
	Databese   *Database
	Secret     *Secret
	Dependency *Dependency
	HubApiIp   string
	HubApiPort uint16
}

func NewVariable() *Variable {
return &Variable{
		Databese: &Database{
			Uri:      "bolt://localhost:7687",
			Username: "neo4j",
			Password: "password",
		},
		Secret: &Secret{
			Controller: "secret variable",
		},
		HubApiIp: "localhost",
	}
}
