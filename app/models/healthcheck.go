package models

// HealthCheckResponse representa o retorno da rota de health-check da aplicação
type HealthCheckResponse struct {
	ApplicationVersion string `json:"applicationVersion"`
	Application        struct {
		Goroutines int    `json:"goroutines"`
		HeapAlloc  string `json:"heapAlloc"`
	} `json:"application"`
}
