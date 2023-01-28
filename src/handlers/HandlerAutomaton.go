package handlers

import (
	logic "automata_logic"
	"encoding/json"
	mappers "mappers_automata"
	models "models_automata"
	"net/http"
	"strconv"
)

func enableCors(w *http.ResponseWriter, requestMethod string) bool {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).WriteHeader(http.StatusAccepted)

	return requestMethod == "OPTIONS"
}

type AutomatonHandler struct {
	timeLaps logic.ITimeLaps
}

func (h *AutomatonHandler) HandlerAutomaton1D(w http.ResponseWriter, r *http.Request) {
	skip := enableCors(&w, r.Method)
	if skip {
		return
	}

	body_automaton := models.Automaton1D{}

	if err := json.NewDecoder(r.Body).Decode(&body_automaton); err != nil {

		json.NewEncoder(w).Encode(err)

		return

	}

	automaton := mappers.ModelAutomata1D2Logic(&body_automaton)
	timeLaps := logic.ITimeLaps(logic.NewTimeLaps(&automaton))
	h.timeLaps = timeLaps

	json.NewEncoder(w).Encode(automaton.GetBoard())

}

func (h *AutomatonHandler) HandlerNext(w http.ResponseWriter, r *http.Request) {
	skip := enableCors(&w, r.Method)
	if skip {
		return
	}

	if h.timeLaps == nil {
		json.NewEncoder(w).Encode("Automata not sendend.")
		return
	}

	automata := h.timeLaps.Next()
	json.NewEncoder(w).Encode((*automata).GetBoard())
}

func (h *AutomatonHandler) HandlerPrevious(w http.ResponseWriter, r *http.Request) {
	skip := enableCors(&w, r.Method)
	if skip {
		return
	}

	if h.timeLaps == nil {
		json.NewEncoder(w).Encode("Automata not sendend.")
		return
	}

	automata := h.timeLaps.Previous()
	json.NewEncoder(w).Encode((*automata).GetBoard())

}

func (h *AutomatonHandler) HandlerGet(w http.ResponseWriter, r *http.Request) {
	skip := enableCors(&w, r.Method)
	if skip {
		return
	}

	if h.timeLaps == nil {
		json.NewEncoder(w).Encode("Automata not sendend.")
		return
	}

	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode(err)
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	automata := h.timeLaps.Get(uint(id))

	json.NewEncoder(w).Encode((*automata).GetBoard())
}

func (h *AutomatonHandler) HandlerAutomaton2D(w http.ResponseWriter, r *http.Request) {
	skip := enableCors(&w, r.Method)
	if skip {
		return
	}

	body_automaton := models.Automaton2D{}

	if err := json.NewDecoder(r.Body).Decode(&body_automaton); err != nil {
		json.NewEncoder(w).Encode(err)

		return
	}

	automaton := mappers.ModelAutomata2D2Logic(&body_automaton)
	timeLaps := logic.ITimeLaps(logic.NewTimeLaps(&automaton))
	h.timeLaps = timeLaps

	json.NewEncoder(w).Encode(automaton.GetBoard())
}
