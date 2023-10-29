package controllers

import (
	"emailn/internal/domain/exceptions"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(handle Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := handle(w, r)
		if err == nil {
			if err != nil {
				if errors.Is(err, exceptions.DbError) {
					render.Status(r, 500)
				} else {
					render.Status(r, 400)
				}

				render.JSON(w, r, map[string]string{"error": err.Error()})
				return
			}
		}
		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
