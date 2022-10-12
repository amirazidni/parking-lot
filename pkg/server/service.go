package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) create(r *http.Request, qty string) (string, int, error) {
	parkSlot, err := strconv.ParseUint(qty, 10, 64)
	if err != nil {
		return "invalid slot number", http.StatusBadRequest, err
	}

	if s.Manager.CreateParkingLot(r.Context(), int(parkSlot)) != nil {
		return "failed to create new parking lot", http.StatusInternalServerError, err
	}

	response := fmt.Sprintf("Created a parking lot with %v slot(s)", parkSlot)
	return response, http.StatusOK, nil
}

func (s *Server) park(r *http.Request, regisNum, color string) (string, int, error) {
	if regisNum == "" || color == "" {
		errMsg := "plate number or color should not empty"
		return errMsg, http.StatusBadRequest, fmt.Errorf(errMsg)
	}

	slot, err := s.Manager.AllocateParkingLot(r.Context(), regisNum, color)
	if slot < 0 {
		return err.Error(), http.StatusOK, err
	}
	if err != nil {
		return err.Error(), http.StatusInternalServerError, err
	}

	response := fmt.Sprintf("Allocated slot number: %v", slot)
	return response, http.StatusOK, nil
}
func (s *Server) getStatus(r *http.Request) (string, int, error) {
	parkingLots, err := s.Manager.GetParkingLot(r.Context())
	if err != nil {
		return "failed to get parking lot status", http.StatusInternalServerError, err
	}

	var messages []string
	messages = append(messages, "Slot No. Registration No Colour")
	for i, carSlot := range *parkingLots {
		if carSlot.ID != 0 {
			item := fmt.Sprintf("%v %s %s", i+1, carSlot.PlateNumber, carSlot.Color)
			messages = append(messages, item)
		}
	}
	response := strings.Join(messages, "\n")
	return response, http.StatusOK, nil
}

func (s *Server) leavePark(r *http.Request, slot string) (string, int, error) {
	slotID, err := strconv.ParseUint(slot, 10, 64)
	if err != nil {
		return "invalid slot number", http.StatusBadRequest, err
	}

	if s.Manager.LeaveParkingLot(r.Context(), int(slotID)) != nil {
		return "failed to process leave request", http.StatusInternalServerError, err
	}

	response := fmt.Sprintf("Slot number %v is free", slotID)
	return response, http.StatusOK, nil
}
func (s *Server) getCarsByColor(r *http.Request, colour, returnCase string) (string, int, error) {
	if colour == "" {
		errMsg := "colour value should not empty"
		return errMsg, http.StatusBadRequest, fmt.Errorf(errMsg)
	}

	parkingLots, err := s.Manager.GetParkingLot(r.Context())
	if err != nil {
		return "failed to get parking lot status", http.StatusInternalServerError, err
	}

	var cars []string
	for _, carSlot := range *parkingLots {
		if carSlot.Color == colour {
			switch returnCase {
			case "ID":
				cars = append(cars, fmt.Sprintf("%v", carSlot.ID))
			case "PlateNumber":
				cars = append(cars, carSlot.PlateNumber)
			default:
				errMsg := "not recognize"
				return errMsg, http.StatusBadRequest, fmt.Errorf(errMsg)
			}
		}
	}
	response := strings.Join(cars, ", ")
	return response, http.StatusOK, nil
}

func (s *Server) getCarsPlate(r *http.Request, colour string) (string, int, error) {
	return s.getCarsByColor(r, colour, "PlateNumber")
}

func (s *Server) getCarsSlot(r *http.Request, colour string) (string, int, error) {
	return s.getCarsByColor(r, colour, "ID")
}

func (s *Server) getSlotNumber(r *http.Request, regisNum string) (string, int, error) {
	if regisNum == "" {
		errMsg := "registration number should not empty"
		return errMsg, http.StatusBadRequest, fmt.Errorf(errMsg)
	}

	parkingLots, err := s.Manager.GetParkingLot(r.Context())
	if err != nil {
		return "failed to get parking lot status", http.StatusInternalServerError, err
	}

	for _, carSlot := range *parkingLots {
		if carSlot.PlateNumber == regisNum && carSlot.ID != 0 {
			response := fmt.Sprintf("%v", carSlot.ID)
			return response, http.StatusOK, nil
		}
	}

	errMsg := "Not found"
	return errMsg, http.StatusOK, nil
}
