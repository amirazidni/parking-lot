package server

import (
	"fmt"
	"net/http"
	"parking-lot/pkg/manage"
	"strconv"
	"strings"

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
		return
	}
}

func (s *Server) ParkingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		regisNum := strings.TrimSpace(vars["value"])
		color := strings.TrimSpace(vars["attribute"])

		if regisNum == "" || color == "" {
			errMsg := "plate number or color should not empty"
			err := fmt.Errorf(errMsg)
			WriteFailResponse(w, http.StatusBadRequest, err, errMsg)
			return
		}

		slot, err := s.Manager.AllocateParkingLot(r.Context(), regisNum, color)
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err, err.Error())
			return
		}
		responseMsg := fmt.Sprintf("Allocated slot number: %v", slot)
		WriteSuccessResponse(w, responseMsg)
		return
	}
}

func (s *Server) GetStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parkingLots, err := s.Manager.GetParkingLot(r.Context())
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err, "failed to get parking lot status")
			return
		}
		responseMsg := fmt.Sprintf("Slot No. Registration No Colour\n")
		for i, carSlot := range *parkingLots {
			if carSlot.ID != 0 {
				item := fmt.Sprintf("%v %s %s\n", i+1, carSlot.PlateNumber, carSlot.Color)
				responseMsg = fmt.Sprintf("%s%v", responseMsg, item)
			}
		}
		WriteSuccessResponse(w, responseMsg)
		return
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
		return
	}
}

func (s *Server) GetCarsPlateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		value := strings.TrimSpace(vars["value"])

		if value == "" {
			errMsg := "colour value should not empty"
			err := fmt.Errorf(errMsg)
			WriteFailResponse(w, http.StatusBadRequest, err, errMsg)
			return
		}

		parkingLots, err := s.Manager.GetParkingLot(r.Context())
		if err != nil {
			WriteFailResponse(w, http.StatusInternalServerError, err, "failed to get parking lot status")
			return
		}
		var responseMsg string
		var carsPlates []string
		for _, carSlot := range *parkingLots {
			if carSlot.Color == value {
				carsPlates = append(carsPlates, carSlot.PlateNumber)
			}
		}
		responseMsg = strings.Join(carsPlates, ", ")
		WriteSuccessResponse(w, responseMsg)
		return
	}
}
