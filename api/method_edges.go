package api

import (
	"net/http"
	"sort"

	"github.com/gorilla/mux"

	"github.com/ethereum/go-ethereum/common"
)

func (s *Server) listHonestEdgesHandler(w http.ResponseWriter, r *http.Request) {
	e, err := convertSpecEdgeEdgesToEdges(r.Context(), s.edges.GetHonestEdges())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: Allow params to sort by other fields
	sort.Slice(e, func(i, j int) bool {
		return e[i].CreatedAtBlock < e[j].CreatedAtBlock
	})

	if err := writeJSONResponse(w, 200, e); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}

func (s *Server) listEdgesHandler(w http.ResponseWriter, r *http.Request) {
	specEdges, err := s.edges.GetEdges(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	e, err := convertSpecEdgeEdgesToEdges(r.Context(), specEdges)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: Allow params to sort by other fields
	sort.Slice(e, func(i, j int) bool {
		return e[i].CreatedAtBlock < e[j].CreatedAtBlock
	})

	if err := writeJSONResponse(w, 200, e); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}

func (s *Server) listMiniStakesHandler(w http.ResponseWriter, r *http.Request) {
	specEdges, err := s.edges.GetEdges(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	e, err := convertSpecEdgeEdgesToEdges(r.Context(), specEdges)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	stakeInfoMap := make(map[common.Hash]map[string]*StakeInfo)
	for _, edge := range e {
		if (edge.MiniStaker != common.Address{}) {
			if stakeInfoMap[edge.AssertionHash] == nil {
				stakeInfoMap[edge.AssertionHash] = make(map[string]*StakeInfo)
			}
			if stakeInfoMap[edge.AssertionHash][edge.Type] == nil {
				stakeInfoMap[edge.AssertionHash][edge.Type] = &StakeInfo{
					StakerAddresses:       []common.Address{},
					NumberOfMinistakes:    0,
					StartCommitmentHeight: edge.StartCommitment.Height,
					EndCommitmentHeight:   edge.EndCommitment.Height,
				}
			}
			stakeInfoMap[edge.AssertionHash][edge.Type].StakerAddresses = append(stakeInfoMap[edge.AssertionHash][edge.Type].StakerAddresses, edge.MiniStaker)
			stakeInfoMap[edge.AssertionHash][edge.Type].NumberOfMinistakes++
		}
	}
	ministakesList := make([]Ministakes, 0)
	for assertionHash, stakeInfoPerAssertion := range stakeInfoMap {
		for level, stakeInfo := range stakeInfoPerAssertion {
			ministakesList = append(ministakesList, Ministakes{
				AssertionHash: assertionHash,
				Level:         level,
				StakeInfo:     stakeInfo,
			})
		}
	}

	if err := writeJSONResponse(w, 200, ministakesList); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}

func (s *Server) getEdgeHandler(w http.ResponseWriter, r *http.Request) {
	edgeId := mux.Vars(r)["id"]
	specEdge, err := s.edges.GetEdge(r.Context(), common.HexToHash(edgeId))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	edge, err := convertSpecEdgeEdgeToEdge(r.Context(), specEdge)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	if err := writeJSONResponse(w, 200, edge); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
}
