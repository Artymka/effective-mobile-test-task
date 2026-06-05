package lib

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data,omitempty"`
}

type ErrResponse struct {
	Response
	Err string `json:"error"`
}

func WriteResponse(w http.ResponseWriter, v interface{}) error {
	resp := Response{
		Ok:   true,
		Data: v,
	}

	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(&resp)
}

func WriteError(w http.ResponseWriter, msg string, status int) error {
	resp := ErrResponse{
		Response: Response{false, nil},
		Err:      msg,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(resp)
}
