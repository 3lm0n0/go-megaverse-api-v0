package handler

import (
	"context"
	"encoding/json"
	"fmt"
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
	http.HandleFunc("/polyanet", h.handlePolyanets)
	http.HandleFunc("/generatePolyanetsCross", h.handlePolyanetsDiagonal)
}

func(h *Handler) Handlers() {
	http.HandleFunc("/polyanets", h.handlePolyanets)
	http.HandleFunc("/generatePolyanetsCross", h.handlePolyanetsDiagonal)
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

func(h *Handler) handlePolyanetsDiagonal(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost { 
		writeJSONResponse(response, http.StatusMethodNotAllowed, nil)
	}
	matrix, err := h.service.GeneratePolyantesCross(context.Background(), request)
	if err != nil {
		writeJSONResponse(response, http.StatusInternalServerError, err)
		return	
	}

	writeJSONResponse(response, http.StatusOK, matrix)
}

func writeJSONResponse(response http.ResponseWriter, status int, value any) error {
	response.WriteHeader(status)
	response.Header().Add("content-type", "application/json; charset=UTF-8")

	return json.NewEncoder(response).Encode(value)
}