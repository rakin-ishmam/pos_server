package api

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/rakin-ishmam/pos_server/action/empty"
	"github.com/rakin-ishmam/pos_server/apperr"
)

func jsonDecode(r *http.Request, des interface{}, where, what string) (errAcc *empty.JSON) {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(des); err != nil {
		errAcc = empty.NewJSON(apperr.Parse{Where: where, What: what})
	}

	return
}

func idFetch(r *http.Request, where string) (string, *empty.JSON) {
	vars := mux.Vars(r)
	id := vars["id"]
	if !bson.IsObjectIdHex(id) {
		return "", empty.NewJSON(apperr.NewValidation(where, "id", "invalid"))
	}

	return id, nil
}