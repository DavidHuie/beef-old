package main

import "fmt"
import "os"

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
	}

	return nil, errors.New("Unrecognized command")
}

// import (
// 	zmq "github.com/DavidHuie/beef/vendor/go-zmq"
// )

func main() {
	for {
		fmt.Printf("beef_client> ")
		var request string = os.Stdin.Read()
		fmt.Printf("You entered: %v\n", request)
	}
}
