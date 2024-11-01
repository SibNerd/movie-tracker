package entity

type SearchParams struct {
	Querry string `json:"q"`
	IMDbTT string `json:"tt"`
}

type ShowsSearchResult []struct {
	Show Show `json:"show"`
}

type Show struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Externals ShowExternals `json:"externals"`
	Image     ShowImages    `json:"image"`
	Summary   string        `json:"summary"`
}

type ShowExternals struct {
	Imdb string `json:"imdb"`
}

type ShowImages struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}
