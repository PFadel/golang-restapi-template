package models

// HealthCheckResponse is the struct that contains the application's health info
type HealthCheckResponse struct {
	ApplicationVersion string `json:"applicationVersion"`
	Application        struct {
		Goroutines int    `json:"goroutines"`
		HeapAlloc  string `json:"heapAlloc"`
	} `json:"application"`
}
