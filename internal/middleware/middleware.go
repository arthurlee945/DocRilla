package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateStack(mw ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(mw) - 1; i >= 0; i-- {
			mwFunc := mw[i]
			next = mwFunc(next)
		}

		return next
	}
}
