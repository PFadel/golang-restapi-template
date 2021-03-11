package models

// HealthCheckResponse TODO
type HealthCheckResponse struct {
	ApplicationVersion string `json:"applicationVersion"`
	Application        struct {
		Goroutines int    `json:"goroutines"`
		HeapAlloc  string `json:"heapAlloc"`
	} `json:"application"`
}
