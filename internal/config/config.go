package config

type Config struct {
	SearcherURL      string `env:"SEARCHER_URL" envDefault:"https://api.tvmaze.com/search/shows"`
	ImdbURL          string `env:"IMDB_URL" envDefault:"https://www.imdb.com/title/"`
	DatabaseURL      string `env,required:"DATABASE_URL"`
	DatabaseUser     string `env,required:"DATABASE_USER"`
	DatabasePassword string `env,required:"DATABASE_PASSWORD"`
}
