package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
	"sync"
	"strings"
	"errors"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var (
    users = make(map[string]string)
    // for comment 2
    usersMutex sync.RWMutex
)

func createUser(name string) {
    // for comment 2
    usersMutex.Lock()
    defer usersMutex.Unlock()
    // for comment 4
	users[name] = time.Now().Format(time.RFC1123Z)
}

func validateRequest(r *http.Request) (string, error) {
    name := strings.TrimSpace(r.URL.Query().Get("name"))
    if name == "" {
		return "", errors.New("name parameter is required")
    }
    if len(name) > 50 {
		return "", errors.New("name too long")
    }
	return name, nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    // for comment 1
	name, err := validateRequest(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	ctx := r.Context()

	// A channel to communicate error from goroutine
	errChan := make(chan error, 1)
	go func() {
		// Recover from panic inside goroutine
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Recovered from panic in goroutine: %v", rec)
				errChan <- fmt.Errorf("internal error")
			}
		}()
		select {
		case <-ctx.Done():
			log.Println("Context cancelled.")
			errChan <- ctx.Err()
		default:
			createUser(name)
			// bubble up errors when createUser could be wrong
			errChan <- nil
		}
    }()

	select {
	case err := <-errChan:
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	    w.WriteHeader(http.StatusOK)
	case <-ctx.Done():
		http.Error(w, "Request cancelled", http.StatusRequestTimeout)
	}
}
