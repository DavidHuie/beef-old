package main

import (
	"flag"
	"fmt"
	"github.com/DavidHuie/beef"
	"net/http"
	"strconv"
)

const (
	// Value response codes
	true_value  = "t"
	false_value = "f"

	// Integer response codes
	success                 = "0"
	invalid_parameter_value = "1"
	invalid_parameters      = "2"
	bf_exists               = "3"
	bf_does_not_exist       = "4"
)

func create_handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name_values, name_provided := r.Form["name"]
	size_values, size_provided := r.Form["size"]

	if !(name_provided && size_provided) {
		fmt.Fprintf(w, invalid_parameters)
		return
	}

	size, error := strconv.ParseInt(size_values[0], 10, 64)

	if error != nil {
		fmt.Fprintf(w, invalid_parameter_value)
		return
	}

	uint_size := uint64(size)

	if error := bf_manager.CreateBF(name_values[0], uint_size); error != nil {
		fmt.Fprintf(w, bf_exists)
		return
	}

	fmt.Printf(
		"Created bloom filter with name %v and size %v\n",
		name_values[0],
		uint_size,
	)

	fmt.Fprintf(w, success)
}

func delete_handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name_values, name_provided := r.Form["name"]

	if !name_provided {
		fmt.Fprintf(w, invalid_parameters)
		return
	}

	if error := bf_manager.DeleteBF(name_values[0]); error != nil {
		fmt.Fprintf(w, bf_does_not_exist)
		return
	}

	fmt.Printf("Deleted bloom filter with name %v\n", name_values[0])
	fmt.Fprintf(w, success)
}

func exists_handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name_values, name_provided := r.Form["name"]

	if !name_provided {
		fmt.Fprintf(w, invalid_parameters)
		return
	}

	value := bf_manager.ExistsBF(name_values[0])

	if value {
		fmt.Fprintf(w, true_value)
	} else {
		fmt.Fprintf(w, false_value)
	}

	fmt.Printf(
		"Checked existence of bloom filter with name %v, %v\n",
		name_values[0],
		value,
	)
}

func insert_handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name_values, name_provided := r.Form["name"]
	value_values, value_provided := r.Form["value"]

	if !(name_provided && value_provided) {
		fmt.Fprintf(w, invalid_parameters)
		return
	}

	error := bf_manager.InsertBF(name_values[0], value_values[0])
	if error != nil {
		fmt.Fprintf(w, bf_does_not_exist)
		return
	}

	fmt.Fprintf(w, success)
	fmt.Printf(
		"Inserted value %v into bloom filter %v\n",
		value_values[0],
		name_values[0],
	)
}

func check_handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name_values, name_provided := r.Form["name"]
	value_values, value_provided := r.Form["value"]

	if !(name_provided && value_provided) {
		fmt.Fprintf(w, invalid_parameters)
		return
	}

	value, error := bf_manager.CheckBF(name_values[0], value_values[0])
	if error != nil {
		fmt.Fprintf(w, bf_does_not_exist)
		return
	}

	if value {
		fmt.Fprintf(w, true_value)
	} else {
		fmt.Fprintf(w, false_value)
	}

	fmt.Printf(
		"Check for existence of %v in bloom filter %v, %v\n",
		value_values[0],
		name_values[0],
		value,
	)
}

var bf_manager = beef.NewManager()

const (
	default_port = 8080
)

func handle_http(port int, listen_on_all_interfaces bool) {
	http.HandleFunc("/create", create_handler)
	http.HandleFunc("/delete", delete_handler)
	http.HandleFunc("/exists", exists_handler)
	http.HandleFunc("/insert", insert_handler)
	http.HandleFunc("/check", check_handler)

	var interface_prefix string

	if listen_on_all_interfaces {
		interface_prefix = "0.0.0.0"
	} else {
		interface_prefix = ""
	}

	error := http.ListenAndServe(interface_prefix+":"+strconv.Itoa(port), nil)
	if error != nil {
		panic(error)
	}
}

func main() {
	var port int
	var listen_on_all_interfaces bool
	flag.IntVar(&port, "port", default_port, "Port to listen on")
	flag.BoolVar(&listen_on_all_interfaces, "all", false, "Listen on all interfaces")
	flag.Parse()

	handle_http(port, listen_on_all_interfaces)
}
