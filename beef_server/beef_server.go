package main

import (
	"fmt"
	"github.com/DavidHuie/beef"
	"net/http"
	"strconv"
)

func create_handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name_values, name_provided := r.Form["name"]
	size_values, size_provided := r.Form["size"]

	if !(name_provided && size_provided) {
		fmt.Fprintf(w, "invalid_parameters")
		return
	}

	size, error := strconv.ParseInt(size_values[0], 10, 64)

	if error != nil {
		fmt.Fprintf(w, "invalid_size")
		return
	}

	uint_size := uint64(size)

	if error := bf_manager.CreateBF(name_values[0], uint_size); error != nil {
		fmt.Fprintf(w, "bf_exists")
		return
	}

	fmt.Printf(
		"Created bloom filter with name %v and size %v\n",
		name_values[0],
		uint_size,
	)

	fmt.Fprintf(w, "OK")
}

func delete_handler(w http.ResponseWriter, r *http.Request) {
}

func exists_handler(w http.ResponseWriter, r *http.Request) {
}

func insert_handler(w http.ResponseWriter, r *http.Request) {
}

func check_handler(w http.ResponseWriter, r *http.Request) {
}

var bf_manager = beef.NewManager()

func main() {
	http.HandleFunc("/create_bf", create_handler)
	http.HandleFunc("/delete_bf", delete_handler)
	http.HandleFunc("/exists_bf", exists_handler)
	http.HandleFunc("/insert_bf", insert_handler)
	http.HandleFunc("/check_bf", check_handler)

	http.ListenAndServe(":8080", nil)
}
