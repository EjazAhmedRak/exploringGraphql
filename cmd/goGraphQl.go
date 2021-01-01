package main

import (
	"encoding/json"
	"fmt"
	"gographql/pkg/graphqlexecutor"
	"net/http"
)

func main() {

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphqlexecutor.ExecuteQuery(r.URL.Query().Get("query"), graphqlexecutor.Schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'")
	http.ListenAndServe(":8080", nil)
}
