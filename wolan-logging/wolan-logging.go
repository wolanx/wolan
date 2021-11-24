package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/_bulk", _bulk)
	err := http.ListenAndServe(":20100", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// {
//  "name" : "dep-elasticsearch-75f55c6b49-8tfjj",
//  "cluster_name" : "docker-cluster",
//  "cluster_uuid" : "1LRsCfEsS-O0DU1nYwOXsg",
//  "version" : {
//    "number" : "7.12.0",
//    "build_flavor" : "default",
//    "build_type" : "docker",
//    "build_hash" : "78722783c38caa25a70982b5b042074cde5d3b3a",
//    "build_date" : "2021-03-18T06:17:15.410153305Z",
//    "build_snapshot" : false,
//    "lucene_version" : "8.8.0",
//    "minimum_wire_compatibility_version" : "6.8.0",
//    "minimum_index_compatibility_version" : "6.0.0-beta1"
//  },
//  "tagline" : "You Know, for Search"
//}
func sayHello(w http.ResponseWriter, r *http.Request) {
	l := r.ContentLength
	body := make([]byte, l)

	_, _ = r.Body.Read(body)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"version":{"number":"7.12.0"}}`))
	fmt.Println(r.Method, r.URL, string(body))
}

func _bulk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"took":0,"errors":false}`))
	//_, _ = w.Write([]byte(`{"took":0,"errors":false,"items":[]}`))
	fmt.Println(r.Method, r.URL)
}
