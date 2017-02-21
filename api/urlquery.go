package api

import (
	"net/http"
	"strconv"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/db/query"
)

type urlQueryParser func(g *query.GenInfo, r *http.Request) error

func urlQSkip(from string) urlQueryParser {
	return func(g *query.GenInfo, r *http.Request) error {
		urlq := r.URL.Query()
		skip, err := strconv.ParseInt(urlq.Get("skip"), 10, 64)
		if err != nil {
			return apperr.NewValidation(from, "skip", "invalid")
		}

		g.SetSkip(int(skip))
		return nil
	}
}

func urlQLimit(from string) urlQueryParser {

	return func(g *query.GenInfo, r *http.Request) error {
		urlq := r.URL.Query()
		limit, err := strconv.ParseInt(urlq.Get("limit"), 10, 64)
		if err != nil {
			return apperr.NewValidation(from, "limit", "invalid")
		}

		g.SetLimit(int(limit))
		return nil
	}
}

func urlQOrder(from string) urlQueryParser {
	return func(g *query.GenInfo, r *http.Request) error {
		urlq := r.URL.Query()
		g.SetOrder(urlq.Get("order"))
		return nil
	}
}

func urlquery(g *query.GenInfo, r *http.Request, params ...urlQueryParser) error {
	// urlq := r.URL.Query()
	//
	// skip, _ := strconv.ParseInt(urlq.Get("skip"), 10, 64)
	// limit, _ := strconv.ParseInt(urlq.Get("limit"), 10, 64)
	//
	// g.SetSkip(int(skip))
	// g.SetLimit(int(limit))

	for _, v := range params {
		if err := v(g, r); err != nil {
			return err
		}
	}

	return nil
}
