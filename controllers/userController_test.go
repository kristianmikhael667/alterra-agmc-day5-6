package controllers

import (
	"fmt"
	"main/config"

	"main/middleware"
	"main/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	//setup
	config.InitDB()
	e := echo.New()
	return e
}

var (
	echoMock = mocks.EchoMock{E: echo.New()}
)

func TestGetUserController(t *testing.T) {
	var testCase = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Berhasil",
			path:                 "/users",
			expectBodyStartsWith: "{\"status\":\"success\",\"users\":[}",
			expectStatus:         http.StatusOK,
		},
	}
	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCase {
		c.SetPath(testCase.path)

		// assertions
		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			// assert.Equal(t, userJSON, rec.Body.String())
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

// Error 400
func TestUsersGetInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := middleware.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/v1/users")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(GetUsersController(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

// Error 200 and 401
func TestGetUsersSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := middleware.CreateToken(1)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/v1/users")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)

	if assert.NoError(t, GetUsersController(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "ID")
		asserts.Contains(body, "name")
		asserts.Contains(body, "email")
	}
}

// Test response 201
func TestResponse201CreateUsers(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPost, "/", nil)
	c.SetPath("/users")
	// testing
	asserts := assert.New(t)
	if asserts.NoError(CreateUserController(c)) {
		asserts.Equal(201, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "Status 201 Created")
	}
}

// Error 500
func TestError500(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/v1/users")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(GetUsersController(c)) {
		asserts.Equal(500, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "Server Error 500")
	}
}
