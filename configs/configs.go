package configs

type Configs struct {
	MongoDB MongoDB
	App     Echo
}

type Echo struct {
	Port string
}

// Database
type MongoDB struct {
	Connection string
	DbName     string
}
