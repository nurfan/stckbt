package action

import (
	"context"
	"log"
	"math"
	"stockbit-service-movie/service/omdb"
	"strconv"

	"stockbit-service-movie/util/errors"

	"github.com/jmoiron/sqlx"
)

type Search struct {
	db      *sqlx.DB
	e       errors.UniError
	service *omdb.Omdb
}

func (s *Search) Handle(ctx context.Context, req SearchRequest) (*SearchResponse, *errors.UniError) {
	var result SearchResponse

	if req.Searchword == "" {
		return &result, s.e.BadRequest("searchword")
	}

	if req.Pagination == 0 {
		req.Pagination = 1
	}

	log.Println("Handle", req)
	res := s.service.Search(req.Searchword, req.Pagination)

	result = s.mappingSearchResponse(res, req)
	return &result, nil
}

func (s *Search) mappingSearchResponse(omdbData omdb.SearchResponse, req SearchRequest) SearchResponse {
	var response SearchResponse

	for _, v := range omdbData.Search {
		m := Movie{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		}
		response.Data = append(response.Data, m)
	}

	tr, _ := strconv.ParseInt(omdbData.TotalResults, 10, 64)
	ps := int64(len(omdbData.Search))
	tp := math.Ceil(float64(tr) / float64(ps))

	response.Pagination = Pagination{
		PageSize:    ps,
		CurrentPage: req.Pagination,
		TotalPage:   int64(tp),
		TotalResult: tr,
	}
	return response
}

func NewSearch(conn *sqlx.DB) *Search {
	return &Search{
		db:      conn,
		service: omdb.NewOmdb(),
	}
}
