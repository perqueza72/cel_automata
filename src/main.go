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

	http.HandleFunc("/automaton/init_automaton1D", handler.HandlerAutomaton1D)
	http.HandleFunc("/automaton/init_automaton2D", handler.HandlerAutomaton2D)
	http.HandleFunc("/automaton/get", handler.HandlerGet)
	http.HandleFunc("/automaton/previous", handler.HandlerPrevious)
	http.HandleFunc("/automaton/next", handler.HandlerNext)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
