package handlers

import (
	logic "automata_logic"
	"encoding/json"
	"fmt"
	mappers "mappers_automata"
	models "models_automata"
	"net/http"
	"strconv"
)

type AutomatonHandler struct {
	timeLaps logic.ITimeLaps
}

func (h *AutomatonHandler) HandlerAutomaton1D(w http.ResponseWriter, r *http.Request) {
	body_automaton := models.Automaton1D{}

	if err := json.NewDecoder(r.Body).Decode(&body_automaton); err != nil {
		fmt.Fprintf(w, "Error getting Json request: %v", err)

		return

	}

	automaton := mappers.ModelAutomata1D2Logic(&body_automaton)
	timeLaps := logic.ITimeLaps(logic.NewTimeLaps(&automaton))
	h.timeLaps = timeLaps

	fmt.Fprintln(w, automaton.GetBoard())
}

func (h *AutomatonHandler) HandlerNext(w http.ResponseWriter, r *http.Request) {
	automata := h.timeLaps.Next()

	fmt.Fprintln(w, (*automata).GetBoard())
}

func (h *AutomatonHandler) HandlerPrevious(w http.ResponseWriter, r *http.Request) {
	automata := h.timeLaps.Previous()

	fmt.Fprintln(w, (*automata).GetBoard())

}

func (h *AutomatonHandler) HandlerGet(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error getting id value. %v", err)
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		fmt.Fprintf(w, "Id must be integer value. %v", err)
	}

	automata := h.timeLaps.Get(uint(id))
	fmt.Fprintln(w, (*automata).GetBoard())

}

func (h *AutomatonHandler) HandlerAutomaton2D(w http.ResponseWriter, r *http.Request) {
	body_automaton := models.Automaton2D{}

	if err := json.NewDecoder(r.Body).Decode(&body_automaton); err != nil {
		fmt.Fprintf(w, "Error getting Json request: %v", err)

		return
	}

	automaton := mappers.ModelAutomata2D2Logic(&body_automaton)
	timeLaps := logic.ITimeLaps(logic.NewTimeLaps(&automaton))
	h.timeLaps = timeLaps

	fmt.Fprintln(w, automaton.GetBoard())
}
