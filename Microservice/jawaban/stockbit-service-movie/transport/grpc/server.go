package grpc

import (
	"context"
	"encoding/json"
	"stockbit-service-movie/action"
	pb "stockbit-service-movie/transport/grpc/proto/movies"

	"github.com/jmoiron/sqlx"
)

// Server: create
type Server struct {
	db *sqlx.DB
}

// NewServer create new Server
func NewServer(conn *sqlx.DB) *Server {
	return &Server{
		db: conn,
	}
}

// Search:  server handler
func (srv *Server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	var response pb.SearchResponse

	databaseConn := srv.db
	request := action.SearchRequest{
		Searchword: req.Searchword,
		Pagination: req.Pagination,
	}

	act := action.NewSearch(databaseConn)
	res, err := act.Handle(ctx, request)

	if err != nil {
		return &response, errorMapping(err)
	}

	bytes, _ := json.Marshal(res)
	_ = json.Unmarshal(bytes, &response)

	return &response, nil
}

// Search:  server handler
func (srv *Server) Detail(ctx context.Context, req *pb.DetailRequest) (*pb.DetailResponse, error) {
	var response pb.DetailResponse

	request := action.DetailRequest{
		ImdbID: req.ImdbID,
	}

	act := action.NewDetail()
	res, err := act.Handle(ctx, request)

	if err != nil {
		return &response, errorMapping(err)
	}

	bytes, _ := json.Marshal(res)
	_ = json.Unmarshal(bytes, &response)

	return &response, nil
}
