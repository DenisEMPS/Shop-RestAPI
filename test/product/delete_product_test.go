package product

import (
	"fmt"
	"net/http/httptest"
	"school21_project1/pkg/handler"
	"school21_project1/pkg/service"
	mock_service "school21_project1/pkg/service/mocks"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func Test_DeleteProductByID(t *testing.T) {
	type mockBehavior func(m *mock_service.MockProduct, id int)

	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		id                   string
		expectetStatusCode   int
		expectetResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(m *mock_service.MockProduct, id int) {
				m.EXPECT().Delete(id).Return(nil)
			},
			id:                   "1",
			expectetStatusCode:   200,
			expectetResponseBody: `{"status":"ok"}`,
		},
		{
			name:                 "invalid id",
			mockBehavior:         func(m *mock_service.MockProduct, id int) {},
			id:                   "asdad",
			expectetStatusCode:   400,
			expectetResponseBody: `{"message":"invalid request params"}`,
		},
		{
			name: "product was not find",
			mockBehavior: func(m *mock_service.MockProduct, id int) {
				m.EXPECT().Delete(id).Return(fmt.Errorf("error"))
			},
			id:                   "100",
			expectetStatusCode:   404,
			expectetResponseBody: `{"message":"product was not find"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			id, _ := strconv.Atoi(testCase.id)

			product := mock_service.NewMockProduct(c)
			testCase.mockBehavior(product, id)

			service := &service.Service{Product: product}
			handler := handler.NewHandler(service)

			r := gin.New()

			api := r.Group("/api/v1/")
			{
				product := api.Group("/product")
				{
					product.DELETE("/:id", handler.DeleteProductByID)
				}
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/product/%v", testCase.id), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectetResponseBody, w.Body.String())
			assert.Equal(t, testCase.expectetStatusCode, w.Code)
		})
	}
}
