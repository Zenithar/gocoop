package protocols

import "time"

type Status struct {
	Status string `json:"status"`
}

// CoopResponse is the response for coop.
type CoopResponse struct {
	OpeningCondition map[string]string `json:"opening_condition"`
	ClosingCondition map[string]string `json:"closing_condition"`
	Latitude         float64           `json:"latitude"`
	Longitude        float64           `json:"longitude"`
	Status           string            `json:"status"`
	IsAutomatic      bool              `json:"is_automatic"`
	OpeningTime      time.Time         `json:"opening_time"`
	ClosingTime      time.Time         `json:"closing_time"`
}

// APIControllerResponse is the response for API.
type APIControllerResponse struct {
	ErrorID          string `json:"error_id,omitempty"`
	ErrorMessage     string `json:"error_message,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// JWToken ...
type JWToken struct {
	Token string `json:"token,omitempty"`
}

// JWTokenRequest ...
type JWTokenRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// HelloResponse contains the result of the Hello request.
type HelloResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}
