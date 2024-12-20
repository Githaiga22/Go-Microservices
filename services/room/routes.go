// services/room/routes.go
package room

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/room", HandleRoom)
}