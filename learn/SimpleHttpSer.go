package learn

import (
	"log"
	"net/http"
	"time"
)

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second*5)
	w.Write([]byte("bye bye ,this is v1 httpServer"))
}

func SimplyHttpSer()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})

	http.HandleFunc("/bye", sayBye)
	log.Println("Starting v1 server ...")
	http.ListenAndServe(":1210", nil)
}
