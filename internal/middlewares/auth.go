package middlewares

import (
	"context"
	"emailn/internal/config"
	"github.com/coreos/go-oidc/v3/oidc"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Authorization") == "" {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "authorization not provided"})
			return
		}

		token := strings.Replace(r.Header.Get("Authorization"), "Bearer", "", 1)

		provider, err := oidc.NewProvider(r.Context(), config.LoadEnv().AuthenticationURL)

		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": "error to connect to provider"})
			return
		}
		_, err = provider.Verifier(&oidc.Config{SkipClientIDCheck: true}).Verify(r.Context(), token)
		if err != nil {
			render.Status(r, 401)
			render.JSON(w, r, map[string]string{"error": "Invalid auth token"})
			return
		}

		decoded, _ := jwtGo.Parse(token, nil)

		email := decoded.Claims.(jwtGo.MapClaims)["email"]

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "email", email)))

	})
}
