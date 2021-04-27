package examples

import (
	"fmt"
	"net/http"
	"time"
)

func _hello(w http.ResponseWriter, req *http.Request) {

	/* A context.Context is created for each request by the net/http machinery, and is available with the Context() method. */
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		_, _ = fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func init() {

	http.HandleFunc("/hello", _hello)
	_ = http.ListenAndServe(":8090", nil)
}
