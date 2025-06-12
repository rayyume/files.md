package main

import (
	"syscall/js"
)

func readFile(path string) (string, error) {
	resultChan := make(chan string, 1)
	errorChan := make(chan error, 1)

	callAsync("read", func(result js.Value, err error) {
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result.String()
	}, path)

	select {
	case result := <-resultChan:
		sendToJS(result)
		return result, nil
	case err := <-errorChan:
		return "", err
	}
}

func exists(path string) (bool, error) {
	resultChan := make(chan bool, 1)
	errorChan := make(chan error, 1)

	callAsync("read", func(result js.Value, err error) {
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result.Bool()
	}, path)

	select {
	case result := <-resultChan:
		sendToJS(result)
		return result, nil
	case err := <-errorChan:
		return false, err
	}
}
