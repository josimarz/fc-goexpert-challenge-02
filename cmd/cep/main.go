package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/josimarz/fc-goexpert-challenge-02/internal/app/cep"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("the search term is empty\nset the search term by command line\nexample: go run main.go 89023-000")
	}
	term := os.Args[1]

	ch := make(chan *cep.Output)

	go func() {
		res, err := cep.Search(cep.ApiCEP, term)
		if err == nil {
			ch <- res
		}
	}()

	go func() {
		res, err := cep.Search(cep.ViaCEP, term)
		if err == nil {
			ch <- res
		}
	}()

	select {
	case res := <-ch:
		fmt.Printf("%s", res.ToJSON())
	case <-time.After(time.Second):
		fmt.Print("timeout")
	}
}
