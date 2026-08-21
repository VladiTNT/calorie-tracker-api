package handlers

import (
	"encoding/json"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"Message": "Hello, world!",
	})
}
