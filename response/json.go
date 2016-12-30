package response

import (
	"io"
	"net/http"
)

// JSON server json response
func JSON(w http.ResponseWriter, h *http.Request, data io.Reader) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	io.Copy(w, data)
}
