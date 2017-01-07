package response

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rakin-ishmam/pos_server/apperr"
)

// ServeErr server error response
func ServeErr(w http.ResponseWriter, h *http.Request, err error) {
	switch err.(type) {
	case apperr.Notfound:
		w.WriteHeader(http.StatusNotFound)
	case *apperr.Notfound:
		w.WriteHeader(http.StatusNotFound)

	case apperr.Access:
		w.WriteHeader(http.StatusUnauthorized)
	case *apperr.Access:
		w.WriteHeader(http.StatusUnauthorized)

	case apperr.Authentication:
		w.WriteHeader(http.StatusUnauthorized)
	case *apperr.Authentication:
		w.WriteHeader(http.StatusUnauthorized)

	case apperr.Database:
		w.WriteHeader(http.StatusInternalServerError)
	case *apperr.Database:
		w.WriteHeader(http.StatusInternalServerError)

	case apperr.Internal:
		w.WriteHeader(http.StatusInternalServerError)
	case *apperr.Internal:
		w.WriteHeader(http.StatusInternalServerError)

	case apperr.Validation:
		w.WriteHeader(http.StatusBadRequest)
	case *apperr.Validation:
		w.WriteHeader(http.StatusBadRequest)

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	buff := &bytes.Buffer{}

	msg := map[string]string{"error": err.Error()}

	enc := json.NewEncoder(buff)
	if err = enc.Encode(msg); err != nil {
		log.Println("error: ServeErr: ", err.Error())
		return
	}

	JSON(w, h, buff)

}
