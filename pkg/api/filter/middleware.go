package filter

import (
	"context"
	"net/http"
	"strconv"
)

const (
	OptionsContextKey = "filter_options"
)

func Middleware(h http.HandlerFunc, defaultLimit int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitFromQuery := r.URL.Query().Get("limit")
		offsetFromQuery := r.URL.Query().Get("offset")

		limit := defaultLimit
		offset := 0
		var limitParseErr error
		var offsetParseErr error
		if limitFromQuery != "" {
			if limit, limitParseErr = strconv.Atoi(limitFromQuery); limitParseErr != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("bad limit"))
			}
		}
		if offsetFromQuery != "" {
			if offset, offsetParseErr = strconv.Atoi(offsetFromQuery); offsetParseErr != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("bad offset"))
			}
		}

		opts := NewOptions(limit, offset)
		ctx := context.WithValue(r.Context(), OptionsContextKey, opts)
		r = r.WithContext(ctx)

		h(w, r)

	}
}
