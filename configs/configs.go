package configs

type Configs struct {
	MongoDB MongoDB
	App     App
}

type App struct {
	Port string
}

// Database
type MongoDB struct {
	Connection string
	DbName     string
}
