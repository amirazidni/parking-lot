package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"parking-lot/pkg/manage"
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
		response, status, err := s.create(r, value)
		if err != nil {
			WriteFailResponse(w, status, err, response)
			return
		}
		WriteSuccessResponse(w, response)
		return
	}
}

func (s *Server) ParkingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		regisNum := strings.TrimSpace(vars["value"])
		color := strings.TrimSpace(vars["attribute"])
		response, status, err := s.park(r, regisNum, color)
		if err != nil {
			WriteFailResponse(w, status, err, response)
			return
		}
		WriteSuccessResponse(w, response)
		return
	}
}

func (s *Server) GetStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, status, err := s.getStatus(r)
		if err != nil {
			WriteFailResponse(w, status, err, response)
			return
		}
		WriteSuccessResponse(w, response)
		return
	}
}

func (s *Server) LeaveParkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slot := vars["value"]
		response, status, err := s.leavePark(r, slot)
		if err != nil {
			WriteFailResponse(w, status, err, response)
			return
		}
		WriteSuccessResponse(w, response)
		return
	}
}

func (s *Server) GetCarsPlateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		value := strings.TrimSpace(vars["value"])
		response, status, err := s.getCarsPlate(r, value)
		if err != nil {
			WriteFailResponse(w, status, err, response)
			return
		}
		WriteSuccessResponse(w, response)
		return
	}
}

func (s *Server) GetCarsSlotHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		value := strings.TrimSpace(vars["value"])
		response, status, err := s.getCarsSlot(r, value)
		if err != nil {
			WriteFailResponse(w, status, err, response)
			return
		}
		WriteSuccessResponse(w, response)
		return
	}
}

func (s *Server) GetSlotNumberHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		value := strings.TrimSpace(vars["value"])
		response, status, err := s.getSlotNumber(r, value)
		if err != nil {
			WriteFailResponse(w, status, err, response)
			return
		}
		WriteSuccessResponse(w, response)
		return
	}
}

func (s *Server) BulkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteFailResponse(w, http.StatusBadRequest, err, "failed to read body request")
			return
		}
		lines := strings.Split(string(body), "\n")
		var status int
		for _, line := range lines {
			var response string
			var err error
			words := strings.Split(line, " ")
			switch words[0] {
			case "create_parking_lot":
				response, status, err = s.create(r, words[1])
			case "park":
				response, status, err = s.park(r, words[1], words[2])
			case "leave":
				response, status, err = s.leavePark(r, words[1])
			case "status":
				response, status, err = s.getStatus(r)
			case "registration_numbers_for_cars_with_colour":
				response, status, err = s.getCarsPlate(r, words[1])
			case "slot_numbers_for_cars_with_colour":
				response, status, err = s.getCarsSlot(r, words[1])
			case "slot_number_for_registration_number":
				response, status, err = s.getSlotNumber(r, words[1])
			default:
				errMsg := "command not recognize"
				WriteBufferResponse(w, http.StatusBadRequest, fmt.Errorf(errMsg), errMsg)
				return
			}
			if err != nil {
				WriteBufferResponse(w, status, err, response)
				continue
			}
			WriteBufferResponse(w, http.StatusOK, nil, response)
			continue
		}
	}
}
