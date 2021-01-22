package http

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"stockbit-service-movie/action"
	"strconv"
)

type Adapter struct {
	db *sqlx.DB
}

func (adp *Adapter) SearchMovie(c echo.Context) error {
	ctx := context.Background()

	pageTmp := c.QueryParam("page")
	page, _ := strconv.ParseInt(pageTmp, 10, 64)

	req := action.SearchRequest{
		Searchword: c.QueryParam("searchword"),
		Pagination: page,
	}

	act := action.NewSearch(adp.db)
	result , _ := act.Handle(ctx, req)
	var resp Response

	data := map[string]interface{}{
		"Movie": result.Data,
		"Pagination": result.Pagination,
	}

	resp.SetSuccessResponse(http.StatusOK, data)
	return c.JSON(resp.Code, resp)
}

func (adp *Adapter) DetailMovie(c echo.Context) error {
	ctx := context.Background()

	req := action.DetailRequest{
		ImdbID: c.Param("imdbid"),
	}

	act := action.NewDetail()
	result , _ := act.Handle(ctx, req)
	var resp Response

	data := map[string]interface{}{
		"Movie": result.Movie,
	}

	resp.SetSuccessResponse(http.StatusOK, data)
	return c.JSON(resp.Code, resp)
}

func NewAdapter(conn *sqlx.DB) *Adapter {
	return &Adapter{
		db: conn,
	}
}