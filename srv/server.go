package srv

import (
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/edit", handleConnections)
	go handleMessages()

	log.Println("server started on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
