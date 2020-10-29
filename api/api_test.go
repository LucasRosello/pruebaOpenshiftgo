package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	router := gin.Default()
	SetupRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/example/1", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
