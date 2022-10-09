package server

import (
	"fmt"
	"net/http"
	"parking-lot/pkg/manage"
	"parking-lot/pkg/util"
	"strconv"

	"github.com/gorilla/mux"
)

func NewServer(m manage.ManagerService) *Server {
	return &Server{
		Manager: m,
	}
}

type Server struct {
	Manager manage.ManagerService
}

func (s *Server) CreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		value := vars["value"]

		parkSlot, err := strconv.Atoi(value)
		if err != nil {
			util.ErrorHandlerFatal(err, "invalid slot number")
		}
		err = s.Manager.CreateParkingLot(r.Context(), uint8(parkSlot))
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, "failed to create new parking lot")
			return
		}
		responseMsg := fmt.Sprintf("Created a parking lot with %v slot(s)", parkSlot)
		WriteSuccessResponse(w, responseMsg)
	}
}

func (s *Server) GetStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parkingLotQty, err := s.Manager.GetParkingLotStatus(r.Context())
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, "failed to get parking lot status")
			return
		}
		responseMsg := fmt.Sprintf("Len: %v slots", parkingLotQty)
		WriteSuccessResponse(w, responseMsg)
	}
}
