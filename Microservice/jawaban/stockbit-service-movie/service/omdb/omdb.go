package omdb

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/parnurzeal/gorequest"
)

type Omdb struct {
	Host   string
	ApiKey string
}

func (o *Omdb) Search(word string, page int64) SearchResponse {
	url := o.Host
	_, body, _ := gorequest.New().
		Get(url).
		Query("apikey=" + o.ApiKey).
		Query("s=" + word).
		Query("page=" + strconv.Itoa(int(page))).
		End()

	var result SearchResponse
	json.Unmarshal([]byte(body), &result)

	return result
}

func (o *Omdb) Detail(imdbID string) DetailResponse {
	url := o.Host
	_, body, _ := gorequest.New().
		Get(url).
		Query("apikey=" + o.ApiKey).
		Query("i=" + imdbID).
		End()

	var result DetailResponse
	json.Unmarshal([]byte(body), &result)
	return result
}

func NewOmdb() *Omdb {
	return &Omdb{
		Host:   os.Getenv("omdb_host"),
		ApiKey: os.Getenv("omdb_key"),
	}
}
