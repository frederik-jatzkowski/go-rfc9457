package main

import (
	"errors"
	"fmt"
	"github.com/frederik-jatzkowski/go-rfc9457"
	"net/http"
)

var (
	ErrOutOfCredit = errors.New("out of credit")
)

func main() {
	registry, err := rfc9457.NewRegistry("http://localhost:8080/")
	if err != nil {
		panic(err)
	}

	err = registry.Define(ErrOutOfCredit)
	if err != nil {
		panic(err)
	}

	handler, err := rfc9457.NewHandler(registry)
	if err != nil {
		panic(err)
	}

	fmt.Println(registry.ProblemTypeFor(errors.New("test error")).Instantiate("test detail").String())
	fmt.Println(registry.ProblemTypeFor(ErrOutOfCredit).Instantiate("test detail").String())

	err = http.ListenAndServe("localhost:8080", handler)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
