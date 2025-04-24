package message

import (
	"context"
	"encoding/json"
	"event-generator/iternal/model"
	"fmt"
	"net/http"
)

func (i *Implementation) Post(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	oreder := model.OrderRequest{}

	err := json.NewDecoder(r.Body).Decode(&oreder)
	r.Body.Close() // посмотреть для чего
	if err != nil {
		http.Error(w, "can not decode json", http.StatusBadRequest)
		return
	}

	err = i.messageService.SendOrderData(ctx, oreder)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
