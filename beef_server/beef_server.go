package main

import (
	"errors"
	"github.com/DavidHuie/beef/bloom_filter"
	zmq "github.com/DavidHuie/beef/vendor/go-zmq"
	"regexp"
	"strconv"
)

const BIT_ARRAY_FACTOR = 12
const NUM_HASHES = 3

type beef_server struct {
	bit_arrays map[string]*bloom_filter.BloomFilter
}

func New() *beef_server {
	bs := new(beef_server)
	return bs
}

func (b *beef_server) CreateBF(name string, size uint64) {
	b.bit_arrays[name] = bloom_filter.New(size, NUM_HASHES)
}

func (b *beef_server) DeleteBF(name string) {
	delete(b.bit_arrays, name)
}

func (b *beef_server) InsertBF(name string, value string) {
	b.bit_arrays[name].Insert(value)
}

func (b *beef_server) CheckBF(name string, value string) bool {
	return b.bit_arrays[name].Check(value)
}

type response struct {
	successful   bool
	return_value string
}

func successful_response(value string) *response {
	response := new(response)
	response.successful = true
	response.return_value = value
	return response
}

var create_regex = regexp.MustCompile(`create (\w+) (\d+)`)
var delete_regex = regexp.MustCompile(`delete (\w+)`)
var insert_regex = regexp.MustCompile(`insert (\w+) (.*)`)
var check_regex = regexp.MustCompile(`check (\w+) (.*)`)

func (b *beef_server) ParseAndExecuteCommand(unparsed_command string) (*response, error) {
	if tokens := create_regex.FindStringSubmatch(unparsed_command); len(tokens[0]) > 1 {
		size, _ := strconv.Atoi(tokens[2])
		b.CreateBF(tokens[1], uint64(size))
		return successful_response("ok"), nil
	} else if tokens := delete_regex.FindStringSubmatch(unparsed_command); len(tokens[0]) > 1 {
		b.DeleteBF(tokens[1])
		return successful_response("ok"), nil
	} else if tokens := insert_regex.FindStringSubmatch(unparsed_command); len(tokens[0]) > 1 {
		b.InsertBF(tokens[1], tokens[2])
		return successful_response("ok"), nil
	} else if tokens := check_regex.FindStringSubmatch(unparsed_command); len(tokens[0]) > 1 {
		value := b.CheckBF(tokens[1], tokens[2])
		if value {
			return successful_response("true"), nil
		} else {
			return successful_response("false"), nil
		}
	} else {
		return nil, errors.New("Unrecognized command")
	}
}

func main() {
	server := new(beef_server)

	// Create ZMQ context
	ctx, err := zmq.NewContext()
	if err != nil {
		panic(err)
	}
	defer ctx.Close()

	// Create ZMQ socket
	sock, err := ctx.Socket(zmq.Rep)
	if err != nil {
		panic(err)
	}
	defer sock.Close()

	// Bind to port
	if err = sock.Bind("tcp://*:5555"); err != nil {
		panic(err)
	}

	for {
		parts, err := sock.Recv()
		if err != nil {
			panic(err)
		}

		request := parts[0]
		response, _ := server.ParseAndExecuteCommand(string(request))

		if err = sock.Send([][]byte{
			[]byte(response.return_value),
		}); err != nil {
			panic(err)
		}
	}
}
