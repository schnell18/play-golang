// Package route
package route

import (
	"net/http"

	"github.com/schnell18/play-golang/restapi/internal/handler"
)

func SetupRoutes(mux *http.ServeMux) {
	handler := handler.NewHandler()
	mux.Handle("GET /health", handler.HealthHandler())
}
