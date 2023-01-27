package main

import (
	handlers "automaton_handler"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Automata started :)")
	handler := handlers.AutomatonHandler{}

	http.HandleFunc("/init_automaton1D", handler.HandlerAutomaton1D)
	http.HandleFunc("/init_automaton2D", handler.HandlerAutomaton2D)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
