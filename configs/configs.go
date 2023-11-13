package configs

type Configs struct {
	App     App
	MongoDB MongoDB
	Redis   Redis
}

type App struct {
	Port string
}

// Database
type MongoDB struct {
	Connection string
	DbName     string
}

// Redis
type Redis struct {
	Host       string
	Pass       string
	ShortCache int
	LongCache  int
}
