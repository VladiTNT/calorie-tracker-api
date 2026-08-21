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
		Value:    fmt.Sprintf("%s:%s", name, secret),
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

func (s *Service) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Defer makes sure the rest of our routing stack runs after this code is executed
		defer h.ServeHTTP(w, r)

		c, err := r.Cookie(AuthCookieName)
		if err != nil {
			return
		}

		var name, pass string
		_, err = fmt.Sscanf(c.Value, "%s:%s", &name, &pass)
		if err != nil {
			return
		}

		if s.Register[name] == pass {
			r = r.WithContext(context.WithValue(r.Context(), AuthContextKey, name))
		}
	})
}
