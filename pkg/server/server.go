package server

import (
	"fmt"
	"net/http"
	"parking-lot/pkg/manage"
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

		parkSlot, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			WriteFailResponse(w, http.StatusBadRequest, err, "invalid slot number")
		}
		if s.Manager.CreateParkingLot(r.Context(), int(parkSlot)) != nil {
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

func (s *Server) LeaveParkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slot := vars["value"]
		slotID, err := strconv.ParseUint(slot, 10, 64)
		if err != nil {
			WriteFailResponse(w, http.StatusBadRequest, err, "invalid slot number")
			return
		}

		if s.Manager.LeaveParkingLot(r.Context(), int(slotID)) != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err, "failed to process leave request")
			return
		}
		responseMsg := fmt.Sprintf("Slot number %v is free", slotID)
		WriteSuccessResponse(w, responseMsg)
	}
}
