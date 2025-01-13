package supplier

import (
	"fmt"
	"net/http/httptest"
	"school21_project1/pkg/handler"
	"school21_project1/pkg/service"
	mock_service "school21_project1/pkg/service/mocks"
	"school21_project1/types"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_GetSupplierByID(t *testing.T) {
	type mockBehavior func(s *mock_service.MockSupplier, id int)

	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		id                   string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "ok",
			mockBehavior: func(s *mock_service.MockSupplier, id int) {
				s.EXPECT().GetByID(id).Return(types.SupplierDAO{Name: "Denis", Country: "UK", City: "London", Street: "Rechnaya", PhoneNumber: "89999999999"}, nil)
			},
			id:                   "1",
			expectedStatusCode:   200,
			expectedResponseBody: `{"name":"Denis","country":"UK","city":"London","street":"Rechnaya","phone_number":"89999999999"}`,
		},
		{
			name:                 "invalid id",
			mockBehavior:         func(s *mock_service.MockSupplier, id int) {},
			id:                   "gfg",
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid request params"}`,
		},
		{
			name: "supplier not founded",
			mockBehavior: func(s *mock_service.MockSupplier, id int) {
				s.EXPECT().GetByID(id).Return(types.SupplierDAO{}, fmt.Errorf("supplier not found"))
			},
			id:                   "1000",
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"supplier not found"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			id, _ := strconv.Atoi(testCase.id)

			Supplier := mock_service.NewMockSupplier(c)
			testCase.mockBehavior(Supplier, id)

			services := &service.Service{Supplier: Supplier}
			handler := handler.NewHandler(services)

			r := gin.New()

			api := r.Group("/api/v1/")
			{
				supplier := api.Group("/supplier")
				supplier.GET("/:id", handler.GetSupplierByID)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/supplier/%v", testCase.id), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}
