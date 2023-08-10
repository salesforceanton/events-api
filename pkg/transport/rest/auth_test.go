package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/salesforceanton/events-api/domain"
	"github.com/salesforceanton/events-api/pkg/service"
	service_mocks "github.com/salesforceanton/events-api/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockAuthorization, user domain.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            domain.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "username", "email": "Test@mockmail.com", "password": "qwerty"}`,
			inputUser: domain.User{
				Username: "username",
				Email:    "Test@mockmail.com",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user domain.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"username": "username"}`,
			inputUser:            domain.User{},
			mockBehavior:         func(r *service_mocks.MockAuthorization, user domain.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Request is invalid type"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"username": "username", "email": "Test@mockmail.com", "password": "qwerty"}`,
			inputUser: domain.User{
				Username: "username",
				Email:    "Test@mockmail.com",
				Password: "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user domain.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("Something went wrong"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"Something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			authService := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(authService, test.inputUser)

			services := &service.Service{Authorization: authService}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-up", handler.SignUp)

			// Create Request and empty Response
			resp := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(resp, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, resp.Code)
			assert.Equal(t, test.expectedResponseBody, resp.Body.String())
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockAuthorization, username, password string)

	tests := []struct {
		name                 string
		username             string
		password             string
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			username:  "username",
			password:  "qwerty",
			inputBody: `{"username": "username", "password": "qwerty"}`,
			mockBehavior: func(r *service_mocks.MockAuthorization, username, password string) {
				r.EXPECT().GenerateToken(username, password).Return("test_token", nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"token":"test_token"}`,
		},
		{
			name:                 "Invalid request",
			username:             "username",
			password:             "password",
			inputBody:            `{"username": "username", "password": ""}`,
			mockBehavior:         func(r *service_mocks.MockAuthorization, username, password string) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Request is invalid type"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			authService := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(authService, test.username, test.password)

			services := &service.Service{Authorization: authService}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-in", handler.SignIn)

			// Create Request and empty Response
			resp := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(resp, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, resp.Code)
			assert.Equal(t, test.expectedResponseBody, resp.Body.String())
		})
	}
}
