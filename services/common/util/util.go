package util

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fayleenpc/grpc-kitchen/services/common/genproto/orders"
)

func ParseJSON(r *http.Request, ordersRequest *orders.CreateOrderRequest) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, ordersRequest)
	if err != nil {
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	data, err := json.Marshal(v)
	if err != nil {
		WriteError(w, http.StatusNotAcceptable, err)
	}
	w.Write(data)
}
