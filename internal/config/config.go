package config

type Config struct {
	App

	SearcherURL string `env:"SEARCHER_URL" envDefault:"https://api.tvmaze.com/search/shows"`
	ImdbURL     string `env:"IMDB_URL" envDefault:"https://www.imdb.com/title/"`
	DatabaseDSN string `env,required:"DATABASE_PASSWORD"`
}

type App struct {
	Name string `env:"APP_NAME" envDefault:"movie-tracker"`
	Host string `env:"APP_HOST" envDefault:"localhost"`
	Port string `env:"APP_PORT" envDefault:"8080"`
}
