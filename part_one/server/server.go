package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8081", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		reqData, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("error"))
			return
		}

		log.Println("Request at endpoint :", string(reqData))

		rw.Header().Add("test-header", "test-header-value")
		rw.WriteHeader(http.StatusAccepted)
		rw.Write([]byte("Endpoint called\n"))
	})))

}
