package dtos

import "time"

type HealthResponse struct {
	Message   string            `json:"message"`
	Status    string            `json:"status"`
	Services  map[string]string `json:"services"`
	Timestamp time.Time         `json:"timestamp"`
}
