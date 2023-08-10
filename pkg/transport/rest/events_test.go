package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/salesforceanton/events-api/domain"
	"github.com/salesforceanton/events-api/pkg/service"
	service_mocks "github.com/salesforceanton/events-api/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
)

var testEvent = domain.Event{
	Id:            1,
	Title:         "go to golang",
	StartDatetime: time.Now().Local().String(),
	TimezoneId:    "America/Los_Angeles",
	OrganizerId:   1,
	Description:   "Free meeting",
}

var testSaveRequest = domain.SaveEventRequest{
	Title:         "go to golang",
	StartDatetime: time.Now().Local().String(),
	TimezoneId:    "America/Los_Angeles",
	Description:   "Free meeting",
}

var invalidTestSaveRequest = domain.SaveEventRequest{
	Title:       "go to golang",
	Description: "Free meeting",
}

func TestHandler_getAll(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockEvents, userId int)

	responseBody, _ := json.Marshal(EventsResponse{
		[]domain.Event{testEvent},
	})

	tests := []struct {
		name                 string
		userId               int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "Ok",
			userId: 1,
			mockBehavior: func(r *service_mocks.MockEvents, userId int) {
				r.EXPECT().GetAll(userId).Return([]domain.Event{testEvent}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: string(responseBody),
		},
		{
			name:   "Service Error",
			userId: 1,
			mockBehavior: func(r *service_mocks.MockEvents, userId int) {
				r.EXPECT().GetAll(userId).Return(nil, errors.New("Something went wrong"))
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

			eventsService := service_mocks.NewMockEvents(c)
			test.mockBehavior(eventsService, test.userId)

			services := &service.Service{Events: eventsService}
			handler := Handler{services}

			// Init Endpoint
			gin.SetMode(gin.TestMode)

			// Create mock context with user-id
			resp := httptest.NewRecorder()
			ctx, r := gin.CreateTestContext(resp)

			r.Use(func(ctx *gin.Context) {
				ctx.Set(USER_CTX, 1)
			})

			// Configure routet
			r.GET("/events", handler.GetAll)

			// Do request
			ctx.Request, _ = http.NewRequest(http.MethodGet, "/events", nil)
			r.ServeHTTP(resp, ctx.Request)

			// Assert
			assert.Equal(t, test.expectedStatusCode, resp.Code)
			assert.Equal(t, test.expectedResponseBody, resp.Body.String())
		})
	}
}

func TestHandler_Create(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockEvents, userId int, request domain.SaveEventRequest)

	tests := []struct {
		name                 string
		userId               int
		saveRequest          domain.SaveEventRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			userId:      1,
			saveRequest: testSaveRequest,
			mockBehavior: func(r *service_mocks.MockEvents, userId int, request domain.SaveEventRequest) {
				r.EXPECT().Create(userId, testSaveRequest).Return(1, nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"Status":"Event record [id]:1 has been saved successfully"}`,
		},
		{
			name:                 "Invalid Request",
			userId:               1,
			saveRequest:          invalidTestSaveRequest,
			mockBehavior:         func(r *service_mocks.MockEvents, userId int, request domain.SaveEventRequest) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Request is invalid type"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			eventsService := service_mocks.NewMockEvents(c)
			test.mockBehavior(eventsService, test.userId, test.saveRequest)

			services := &service.Service{Events: eventsService}
			handler := Handler{services}

			// Init Endpoint
			gin.SetMode(gin.TestMode)

			// Create mock context with user-id
			resp := httptest.NewRecorder()
			ctx, r := gin.CreateTestContext(resp)

			r.Use(func(ctx *gin.Context) {
				ctx.Set(USER_CTX, 1)
			})

			// Configure routet
			r.GET("/events", handler.Create)

			// Do request

			reqBody, _ := json.Marshal(test.saveRequest)
			ctx.Request, _ = http.NewRequest(http.MethodGet, "/events", bytes.NewBuffer(reqBody))
			r.ServeHTTP(resp, ctx.Request)

			// Assert
			assert.Equal(t, test.expectedStatusCode, resp.Code)
			assert.Equal(t, test.expectedResponseBody, resp.Body.String())
		})
	}
}
