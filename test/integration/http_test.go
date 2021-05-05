// Package http provides ...
package integration_test

import (
   "net/http"
   "net/http/httptest"
   "testing"
   "github.com/stretchr/testify/assert"
   bookRoute "github.com/VictorEspiritu/goAPIRest/cmd/http"
)

func TestPingRoute(t *testing.T) {
   router := bookRoute.SetupRouter()
   w := httptest.NewRecorder()
   req, _ := http.NewRequest("GET", "/ping", nil)
   router.ServeHTTP(w, req)

   assert.Equal(t, 200, w.Code)
   assert.JSONEq(t, `{"message": "pong"}`, w.Body.String())

}
