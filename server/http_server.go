package server

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
)

var nodesH map[string]string
var edgesH map[string][]string

func Run(nodes *map[string]string, edges *map[string][]string) {
	mux := http.NewServeMux()

	nodesH = *nodes
	edgesH = *edges
	mux.HandleFunc("/", ShowGraph)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3007")
}

func ShowGraph(rw http.ResponseWriter, r *http.Request) {

	rw.Write([]byte(fmt.Sprintf("nodes %v\n", nodesH)))
	rw.Write([]byte(fmt.Sprintf("edges %v\n", edgesH)))
}
