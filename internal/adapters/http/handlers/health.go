package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheck struct {
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

// RegisterHealth register the liveness and readiness probe endpoints
func (h *HealthCheck) RegisterHealth(server *echo.Echo) {
	server.GET("/liveness", h.Liveness)
	server.GET("/readiness", h.Readiness)
}

// Liveness godoc
// @Summary Show the status of http.
// @Description get the status of http.
// @Tags Health
// @Accept */*
// @Produce json
// @Success 200 {string} string "token"
// @Router /liveness [get]
func (h *HealthCheck) Liveness(c echo.Context) error {
	response := make(map[string]bool)
	response["status"] = true
	return c.JSON(http.StatusOK, response)
}

// Readiness godoc
// @Summary Show the status of http.
// @Description get the status of http.
// @Tags Health
// @Accept */*
// @Produce json
// @Success 200 {string} string "OK"
// @Router /readiness [get]
func (h *HealthCheck) Readiness(c echo.Context) error {
	response := make(map[string]bool)
	response["status"] = true
	return c.JSON(http.StatusOK, response)
}
