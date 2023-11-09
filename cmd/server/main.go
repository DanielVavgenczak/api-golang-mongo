package main

import (
	"context"
	"fmt"
	"time"
)

type Director struct {
	Name string `json:"name"`
}

type Writer struct {
	Name string `json:"name"`
}

type Movie struct {
	Name string `json:"name"`
	Year string `json:"year"`
	Directors []Director `json:"directors"`
	Writers []Writer `json:"writers"`
}

type BoxOffice struct {
	Budget uint64 `json:"budget"`
	Gross uint64 `json:"gross"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("Build...")
}