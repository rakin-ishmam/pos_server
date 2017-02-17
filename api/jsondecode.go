package api

import (
	"encoding/json"
	"net/http"

	"github.com/rakin-ishmam/pos_server/action/empty"
	"github.com/rakin-ishmam/pos_server/apperr"
)

func jsonDecode(r *http.Request, des interface{}, where, what string) (errAcc *empty.JSON) {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(des); err != nil {
		errAcc = &empty.JSON{Err: apperr.Parse{Where: where, What: what}}
	}

	return
}
