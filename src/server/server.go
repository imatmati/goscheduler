package server

import (
	"encoding/json"
	"fmt"
	"heap"
	"net/http"
	"node"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"
)

//Options carry values to set the server
type Options struct {
	Addr              string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}

//Run launches the server with passed options.
func Run(options Options) {

	manager := heap.New()

	router := mux.NewRouter()
	// Example of call : curl -X POST -d '{"Priority":1,"Load":{"TTL":1203,"user":"kfk56kgl3S"}}' http://127.0.0.1:8080/jobs
	router.HandleFunc("/jobs", func(rw http.ResponseWriter, request *http.Request) {
		var node node.Node
		decoder := json.NewDecoder(request.Body)
		if err := decoder.Decode(&node); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		if err := manager.Push(&node); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			//TODO add a message
			return
		}
		rw.WriteHeader(http.StatusCreated)
	}).Methods("POST")

	srv := http.Server{Handler: router}

	copier.Copy(&srv, &options)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
