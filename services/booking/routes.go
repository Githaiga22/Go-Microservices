// services/booking/routes.go
package booking

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/booking", HandleBooking)
}