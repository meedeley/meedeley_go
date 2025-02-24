package pkg

type Response struct {
	Status  int `json:"status"`
	Message any `json:"message,omitempty"`
	Data    any `json:"data,omitempty"`
}
