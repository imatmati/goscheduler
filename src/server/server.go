package server

import (
	"encoding/json"
	"fmt"
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

	router := mux.NewRouter()
	router.HandleFunc("/jobs", func(rw http.ResponseWriter, request *http.Request) {
		var node node.Node
		decoder := json.NewDecoder(request.Body)
		if err := decoder.Decode(&node); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(node)
	})

	srv := http.Server{Handler: router}

	copier.Copy(&srv, &options)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
