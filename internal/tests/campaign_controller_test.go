package tests

import (
	"bytes"
	"emailn/internal/contract"
	"emailn/internal/controllers"
	"emailn/internal/mocks"
	"encoding/json"
	"fmt"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController(t *testing.T) {
	var (
		fake = faker.New()
		body = contract.NewCampaign{
			Name:    fake.Lorem().Text(9),
			Content: fake.Lorem().Text(20),
			Email:   []string{},
		}
	)
	t.Run("should save new campaign", func(t *testing.T) {

		assertions := assert.New(t)
		service := new(mocks.ServiceMock)
		service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
			if request.Name == body.Name && request.Content == body.Content {
				return true
			} else {
				return false
			}
		})).Return("34x", nil)

		handler := controllers.CampaignHandler{
			CampaignService: service,
		}

		var buffer bytes.Buffer
		json.NewEncoder(&buffer).Encode(body)

		req, _ := http.NewRequest("POST", "/", &buffer)
		rr := httptest.NewRecorder()

		_, status, err := handler.PostCampaign(rr, req)
		assertions.Equal(201, status)
		assertions.Nil(err)
	})
	t.Run("should throw error when exists", func(t *testing.T) {
		assertions := assert.New(t)
		service := new(mocks.ServiceMock)
		service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))

		handler := controllers.CampaignHandler{
			CampaignService: service,
		}

		var buffer bytes.Buffer
		json.NewEncoder(&buffer).Encode(body)

		req, _ := http.NewRequest("POST", "/", &buffer)
		rr := httptest.NewRecorder()

		_, _, err := handler.PostCampaign(rr, req)

		assertions.NotNil(err)
	})
}
