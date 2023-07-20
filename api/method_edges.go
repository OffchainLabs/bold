package api

import (
	"net/http"
	"sort"
)

func (s *Server) listEdgesHandler(w http.ResponseWriter, r *http.Request) {
	e, err := convertSpecEdgeEdgesToEdges(r.Context(), s.data.GetEdges())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO: Allow params to sort by other fields
	sort.Slice(e, func(i, j int) bool {
		return e[i].CreatedAtBlock < e[j].CreatedAtBlock
	})

	if err := writeJSONResponse(w, 200, e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (s *Server) getEdgeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}
