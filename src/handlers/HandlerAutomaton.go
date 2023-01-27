package handlers

import (
	logic "automata_logic"
	"encoding/json"
	"fmt"
	mappers "mappers_automata"
	models "models_automata"
	"net/http"
)

type AutomatonHandler struct {
	timeLaps logic.TimeLaps
}

func (h *AutomatonHandler) HandlerAutomaton1D(w http.ResponseWriter, r *http.Request) {
	body_automaton := models.Automaton1D{}

	if err := json.NewDecoder(r.Body).Decode(&body_automaton); err != nil {
		fmt.Fprintf(w, "Error getting Json request: %v", err)

		return

	}

	automaton := mappers.ModelAutomata1D2Logic(&body_automaton)
	h.timeLaps = *logic.NewTimeLaps(&automaton)

}

func (h *AutomatonHandler) HandlerAutomaton2D(w http.ResponseWriter, r *http.Request) {
	body_automaton := models.Automaton2D{}

	if err := json.NewDecoder(r.Body).Decode(&body_automaton); err != nil {
		fmt.Fprintf(w, "Error getting Json request: %v", err)

		return
	}

	automaton := mappers.ModelAutomata2D2Logic(&body_automaton)
	h.timeLaps = *logic.NewTimeLaps(&automaton)
}
