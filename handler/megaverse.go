package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"megaverse/domain"
	service "megaverse/service"
	"net/http"
)

type PolyanetHandlerInterface interface {
	PolyanetHandlers()
}

type Handler struct {
	service service.PolyanetService
}

func NewPolyanetHandler(s service.PolyanetService) *Handler {
	return &Handler{
		service: s,
	}
}

func(h *Handler) PolyanetHandlers() {
	http.HandleFunc("/megaverse/polyanet", h.handlePolyanets)
	http.HandleFunc("/megaverse/polyanet/cross", h.handlePolyanetsCross)
	http.HandleFunc("/megaverse/logo", h.handleCreateLogo)
}

func(h *Handler) Handlers() {
	http.HandleFunc("/polyanets", h.handlePolyanets)
	http.HandleFunc("/generatePolyanetsCross", h.handlePolyanetsCross)
}

func(h *Handler) handlePolyanets(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodDelete:
		fmt.Println("http request method: ", request.Method)
		h.handleDeletePolyanet(response, request)

	case http.MethodPost:
		fmt.Println("http request method: ", request.Method)
		h.handleCreatePolyanet(response, request)

	default:
		writeJSONResponse(response, http.StatusMethodNotAllowed, nil)
	}
}

func(h *Handler) handleDeletePolyanet(response http.ResponseWriter, request *http.Request) {
	polyanet, err := h.service.DeletePolyanet(context.Background(), request)
	if err != nil {
		writeJSONResponse(response, http.StatusInternalServerError, err)
		return	
	}

	writeJSONResponse(response, http.StatusOK, polyanet)
}

func(h *Handler) handleCreatePolyanet(response http.ResponseWriter, request *http.Request) {
	polyanet, err := h.service.CreatePolyanet(context.Background(), request, nil)
	if err != nil {
		writeJSONResponse(response, http.StatusInternalServerError, err)
		return
	}

	writeJSONResponse(response, http.StatusCreated, polyanet)
}

func(h *Handler) handlePolyanetsCross(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost { 
		writeJSONResponse(response, http.StatusMethodNotAllowed, nil)
	}
	matrix, err := h.service.CreatePolyantesCross(context.Background(), request)
	if err != nil {
		writeJSONResponse(response, http.StatusInternalServerError, err)
		return	
	}

	writeJSONResponse(response, http.StatusOK, matrix)
}

func(h *Handler) handleCreateLogo(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost { 
		writeJSONResponse(response, http.StatusMethodNotAllowed, nil)
	}
	// read request body.
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil  {
		writeJSONResponse(response, http.StatusInternalServerError, err)
	}
	// unmarshal request body.
	var matrix *domain.MegaverseGoal
	err = json.Unmarshal(requestBody, &matrix)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// create logo.
	m, err := h.service.CreateLogo(context.Background(), matrix)
	if err != nil {
		writeJSONResponse(response, http.StatusInternalServerError, err)
		return	
	}

	writeJSONResponse(response, http.StatusCreated, &m.Goal)
}

func writeJSONResponse(response http.ResponseWriter, status int, value any) error {
	response.WriteHeader(status)
	response.Header().Add("content-type", "application/json; charset=UTF-8")

	return json.NewEncoder(response).Encode(value)
}