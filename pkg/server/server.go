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
			WriteFailResponse(w, http.StatusInternalServerError, err, "failed to create new parking lot")
			return
		}
		responseMsg := fmt.Sprintf("Created a parking lot with %v slot(s)", parkSlot)
		WriteSuccessResponse(w, responseMsg)
	}
}

func (s *Server) ParkingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		regisNum := vars["value"]
		color := vars["attribute"]

		if regisNum == "" || color == "" {
			errMsg := "plate number or color should not empty"
			err := fmt.Errorf(errMsg)
			WriteFailResponse(w, http.StatusBadRequest, err, errMsg)
			return
		}

		slot, err := s.Manager.AllocateParkingLot(r.Context(), regisNum, color)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err, "failed to allocating parking lot")
			return
		}
		responseMsg := fmt.Sprintf("Allocated slot number: %v", slot)
		WriteSuccessResponse(w, responseMsg)
	}
}

func (s *Server) GetStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parkingLots, err := s.Manager.GetParkingLot(r.Context())
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err, "failed to get parking lot status")
			return
		}
		fmt.Printf("Parking lots: %+v \n", parkingLots)
		responseMsg := fmt.Sprintf("Parking lots: %+v", parkingLots)
		WriteSuccessResponse(w, responseMsg)
	}
}
