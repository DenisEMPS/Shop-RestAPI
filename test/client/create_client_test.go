package client

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"school21_project1/pkg/handler"
	"school21_project1/pkg/service"
	mock_service "school21_project1/pkg/service/mocks"
	"school21_project1/types"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_CreateClient(t *testing.T) {
	type mockBehavior func(s *mock_service.MockClient, client types.CreateClient)

	testTable := []struct {
		name                 string
		inputBody            string
		intputClient         types.CreateClient
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Denis", "surname":"Solosenko", "birthday":"1999-11-11", "gender":true, "country":"UK", "city":"London", "street":"Pawden 21/2"}`,
			intputClient: types.CreateClient{
				Name:     "Denis",
				Surname:  "Solosenko",
				Birthday: "1999-11-11",
				Gender:   true,
				Country:  "UK",
				City:     "London",
				Street:   "Pawden 21/2",
			},
			mockBehavior: func(s *mock_service.MockClient, client types.CreateClient) {
				s.EXPECT().Create(client).Return(1, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "input fail",
			inputBody:            `{"name":"Denis", "surname":"Solosenko", "gender":"false", "country":"UK", "city":"London", "street":"Pawden 21/2"}`,
			mockBehavior:         func(s *mock_service.MockClient, client types.CreateClient) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid request params"}`,
		},
		{
			name:      "internal server error",
			inputBody: `{"name":"Denis", "surname":"Solosenko", "birthday":"1999-11-11", "gender":true, "country":"UK", "city":"London", "street":"Pawden 21/2"}`,
			intputClient: types.CreateClient{
				Name:     "Denis",
				Surname:  "Solosenko",
				Birthday: "1999-11-11",
				Gender:   true,
				Country:  "UK",
				City:     "London",
				Street:   "Pawden 21/2",
			},
			mockBehavior: func(s *mock_service.MockClient, client types.CreateClient) {
				s.EXPECT().Create(client).Return(1, fmt.Errorf("server error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"internal server error"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			client := mock_service.NewMockClient(c)
			testCase.mockBehavior(client, testCase.intputClient)

			service := &service.Service{Client: client}
			handler := handler.NewHandler(service)

			r := gin.New()

			api := r.Group("/api/v1/")
			{
				client := api.Group("/client")
				{
					client.POST("/", handler.CreateClient)
				}
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/client/", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
