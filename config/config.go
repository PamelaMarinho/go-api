package config

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	user     string
	Pass     string
	Database string
}

type APIConfig struct {
	Port string
}

type j struct{}

func init() {
	DBConfig.Host
}

//struct: coleção de campos de dados com tipos de dados declarados
//para tornar os identificadores (ex: Database) público para outros pacotes é necessário começar com a letra mauiscula, caso contrário
//fica visível apenas nesse pacote config
