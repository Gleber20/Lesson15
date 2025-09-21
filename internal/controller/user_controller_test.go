package controller

import (
	mock_contracts "Lesson15/internal/contracts/mocks"
	"Lesson15/internal/errs"
	"Lesson15/internal/models"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_CreateUser(t *testing.T) {
	type mockBehaviour func(s *mock_contracts.MockServiceI, u models.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"John","email":"john@example.com","age":25}`,
			inputUser: models.User{Name: "John", Email: "john@example.com", Age: 25},
			mockBehaviour: func(s *mock_contracts.MockServiceI, u models.User) {
				s.EXPECT().CreateUser(gomock.Any(), u).Return(nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"message":"User created successfully!"}`,
		},
		{
			name:                 "Empty fields",
			inputBody:            `{}`,
			mockBehaviour:        func(s *mock_contracts.MockServiceI, u models.User) {},
			expectedStatusCode:   http.StatusUnprocessableEntity,
			expectedResponseBody: `{"error":"invalid field value"}`,
		},
		{
			name:      "Invalid fields",
			inputBody: `{"name":"","email":"", "age":-1}`,
			inputUser: models.User{},
			mockBehaviour: func(s *mock_contracts.MockServiceI, u models.User) {
			},
			expectedStatusCode:   http.StatusUnprocessableEntity,
			expectedResponseBody: `{"error":"invalid field value"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"name":"John","email":"john@example.com","age":25}`,
			inputUser: models.User{Name: "John", Email: "john@example.com", Age: 25},
			mockBehaviour: func(s *mock_contracts.MockServiceI, u models.User) {
				s.EXPECT().CreateUser(gomock.Any(), u).Return(errs.ErrInvalidFieldValue)
			},
			expectedStatusCode:   http.StatusUnprocessableEntity,
			expectedResponseBody: `{"error":"invalid field value"}`,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			// Init dependencies
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock_contracts.NewMockServiceI(ctrl)
			tc.mockBehaviour(mockService, tc.inputUser)

			handler := NewUserController(mockService)

			// Test server
			r := gin.New()
			r.POST("/users", handler.Create)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(tc.inputBody))
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			r.ServeHTTP(w, req)

			// Assert response
			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.JSONEq(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}
