package models

// SampleResponse is an exmple to a response received by the API
type SampleResponse struct {
	Args    map[string]interface{} `json:"args"`
	Headers struct {
		Accept         string `json:"Accept"`
		AcceptEncoding string `json:"Accept-Encoding"`
		CacheControl   string `json:"Cache-Control"`
		Host           string `json:"Host"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceID   string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

// SampleErrorResponse TODO
type SampleErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
