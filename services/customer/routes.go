// services/customer/routes.go
package customer

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/customer", HandleCustomer)
}