package converter

import (
	"bytes"
	"encoding/json"
)

// JSONtoBuff convert json interface to bytes Buffere
func JSONtoBuff(obj interface{}) (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	enc := json.NewEncoder(b)
	if err := enc.Encode(obj); err != nil {
		return nil, err
	}

	return b, nil
}
