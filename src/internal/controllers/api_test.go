package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
	"main/src/internal/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestApi(t *testing.T) {
	// Initialize the test container database
	ctx := context.Background()
	container := models.SetupTestDatabase(t, ctx)
	defer container.Terminate(ctx)

	// Initialize the API
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Set("user", User{Id: "user", Email: "user@test.com"})

	url, _ := url.Parse("")
	c.Request = &http.Request{
		URL:    url,
		Header: make(http.Header),
	}

	ListResourceApi(c)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}
