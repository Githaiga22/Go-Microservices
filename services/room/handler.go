// services/room/handler.go
package room

import (
	"encoding/json"
	"net/http"
)

func HandleRoom(w http.ResponseWriter, r *http.Request) {
	room := Room{}
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, "Welcome Allan Robinson \n Book a room", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}