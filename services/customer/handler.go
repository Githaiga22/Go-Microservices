// services/customer/handler.go
package customer

import (
	"encoding/json"
	"net/http"
)

func HandleCustomer(w http.ResponseWriter, r *http.Request) {
	customer := Customer{}
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Welcome Allan Robinson to our hotel services", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}