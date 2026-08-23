package auth

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
)

const (
	AuthCookieName = "Auth"
	AuthContextKey = "auth"
)

type Service struct {
	Register map[string]string
}

func New() *Service {
	return &Service{
		Register: make(map[string]string),
	}
}

func (s *Service) Add(name string) http.Cookie {
	secret := rand.Text()

	s.Register[name] = secret

	return http.Cookie{
		Name:     AuthCookieName,
		Path:     "/",
		Value:    name + " " + secret,
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

func (s *Service) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(AuthCookieName)
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		var name, token string
		_, err = fmt.Sscan(c.Value, &name, &token)
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		if s.Register[name] == token {
			r = r.WithContext(context.WithValue(r.Context(), AuthContextKey, name))
		}

		h.ServeHTTP(w, r)
	})
}
