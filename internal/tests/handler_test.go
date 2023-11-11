package tests

import (
	"emailn/internal/controllers"
	"emailn/internal/domain/exceptions"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("should an error with status code 500", func(t *testing.T) {
		assertions := assert.New(t)
		endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			return nil, 0, exceptions.UnkownErrror
		}
		handlerFunc := controllers.HandlerError(endpoint)
		req, _ := http.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		handlerFunc.ServeHTTP(res, req)
		assertions.Equal(http.StatusInternalServerError, res.Code)
	})
	t.Run("should an error with status code 400", func(t *testing.T) {
		assertions := assert.New(t)
		endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			return nil, 000, errors.New("vacilou")
		}
		handlerFunc := controllers.HandlerError(endpoint)
		req, _ := http.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		handlerFunc.ServeHTTP(res, req)
		assertions.Equal(http.StatusBadRequest, res.Code)
	})
	t.Run("should return with no errors", func(t *testing.T) {
		assertions := assert.New(t)
		type TestBody struct {
			Id int
		}
		expected := TestBody{Id: 2}
		endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
			return expected, 200, nil
		}
		handlerFunc := controllers.HandlerError(endpoint)
		req, _ := http.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		handlerFunc.ServeHTTP(res, req)

		obj := TestBody{}
		json.Unmarshal(
			res.Body.Bytes(),
			&obj,
		)
		assertions.Equal(http.StatusOK, res.Code)
		assertions.Equal(expected, obj)
	})
}
