package config

var DB = configDB{
	User:     "root",
	Soc:      "",
	Password: "@tcp(127.0.0.1:3306)/",
	Name:     "sistema",
}

type configDB struct {
	User     string
	Password string
	Soc      string
	Name     string
}
