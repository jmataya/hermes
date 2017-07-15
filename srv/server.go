package srv

import "log"

// Run starts the server.
func Run() {
	router := routes()

	log.Println("server started on port 8000")
	log.Fatal(router.Start(":8000"))

	// http.HandleFunc("/edit", handleConnections)
	// go handleMessages()

	// log.Println("server started on port 8000")
	// if err := http.ListenAndServe(":8000", nil); err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
