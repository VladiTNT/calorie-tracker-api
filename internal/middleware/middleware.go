package middleware

import (
	"net/http"
	"slices"
)

type MiddlewareFunc func(http.Handler) http.Handler

func Stack(mws ...MiddlewareFunc) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		for _, mw := range slices.Backward(mws) {
			h = mw(h)
		}
		return h
	}
}
