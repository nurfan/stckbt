package action

import (
	"context"
	"encoding/json"
	"stockbit-service-movie/service/omdb"
	"stockbit-service-movie/util/errors"
)

type Detail struct {
	e       errors.UniError
	service *omdb.Omdb
}

func (s *Detail) Handle(ctx context.Context, req DetailRequest) (*DetailResponse, *errors.UniError) {
	var result DetailResponse

	if req.ImdbID == "" {
		return &result, s.e.BadRequest("IMDB ID")
	}

	movie := s.service.Detail(req.ImdbID)

	bytes, _ := json.Marshal(movie)
	_ = json.Unmarshal(bytes, &result.Movie)

	return &result, nil
}

func NewDetail() *Detail {
	return &Detail{
		service: omdb.NewOmdb(),
	}
}
