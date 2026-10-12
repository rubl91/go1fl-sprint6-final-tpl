package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	l := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	s := server.NewServer(l)

	err := s.Server.ListenAndServe()
	if err != nil {
		l.Fatal(err)
	}
}
