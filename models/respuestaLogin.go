package models

// RespuestaLogin es la respuesta que le damos en el jwt
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
