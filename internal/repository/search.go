package repository

import (
	"encoding/json"
	"fmt"
	entity "movietracker/internal/entities"

	"github.com/valyala/fasthttp"
)

type SearchRepository struct {
	searchURL string
	searcher  fasthttp.Client
}

func NewSearchRepository(url string) *SearchRepository {
	return &SearchRepository{
		searchURL: url,
		searcher:  fasthttp.Client{},
	}
}

func (sr *SearchRepository) SearchShow(params entity.SearchParams) (entity.ShowsSearchResult, error) {
	var res entity.ShowsSearchResult

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	url := fmt.Sprintf("%s?q=%s", sr.searchURL, params.Querry)

	fmt.Printf("request url: %v\n", url)
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.SetContentTypeBytes([]byte("application/json"))

	err := sr.searcher.Do(req, resp)
	if err != nil {
		fmt.Printf("error with request")
		return res, err
	}

	fmt.Printf("response: %s\n", string(resp.Body()))
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		fmt.Printf("error while unmarshalling response data\n")
		return res, err
	}

	return res, nil
}
