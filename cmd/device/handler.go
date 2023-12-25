package device

import (
	"net/http"
)

// DeviceAPIHandler обработчик для Device Service API.
type DeviceAPIHandler struct{}

// NewDeviceHandler создает новый экземпляр DeviceAPIHandler.
func NewDeviceHandler() *DeviceAPIHandler {
	return &DeviceAPIHandler{}
}

// GetJWTTokenHandler обработчик для получения JWT-токена по API токену.
func (h *DeviceAPIHandler) GetJWTTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("/tokens/jwt"))
}

// SendEventHandler обработчик для отправки событий.
func (h *DeviceAPIHandler) SendEventHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("/events"))
}
