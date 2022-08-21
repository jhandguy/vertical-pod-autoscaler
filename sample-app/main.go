package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type response struct {
	Node      string `json:"node"`
	Namespace string `json:"namespace"`
	Pod       string `json:"pod"`
}

func handleSuccess(response response) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(response)
	}
}

func handleError(status int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	}
}

func main() {
	httpPort, _ := os.LookupEnv("HTTP_PORT")
	node, _ := os.LookupEnv("KUBERNETES_NODE")
	namespace, _ := os.LookupEnv("KUBERNETES_NAMESPACE")
	pod, _ := os.LookupEnv("KUBERNETES_POD")
	resp := response{
		Node:      node,
		Namespace: namespace,
		Pod:       pod,
	}

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/monitoring/health", func(_ http.ResponseWriter, _ *http.Request) {})
	httpMux.HandleFunc("/success", handleSuccess(resp))
	httpMux.HandleFunc("/error", handleError(http.StatusInternalServerError))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", httpPort), httpMux))
}
