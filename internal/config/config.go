package config

type Config struct {
	SearcherURL string `env:"SEARCHER_URL" envDefault:"https://api.tvmaze.com/search/shows"`
}
