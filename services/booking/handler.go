// services/booking/handler.go
package booking

import (
	"encoding/json"
	"net/http"
)

func HandleBooking(w http.ResponseWriter, r *http.Request) {
	booking := Booking{}
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Welcome Allan Robinson \n Book hotel services", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}