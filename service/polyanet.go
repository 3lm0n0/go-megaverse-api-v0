package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	domain "megaverse/domain"
	"net/http"
	"strings"
	"time"
)

type PolyanetService interface {
	GeneratePolyantesCross(ctx context.Context, request *http.Request) (*domain.Matrix, error)
	CreatePolyanet(ctx context.Context, request *http.Request, polyanet *PolyanetApiRequest) (*domain.Polyanet, error)
	DeletePolyanet(ctx context.Context, request *http.Request) (*domain.Polyanet, error)
}

type PolyanetServiceParams struct {
	ApiUrl string
	CandidateId string
}

type PolyanetApiRequest  struct {
	candidateId string
	row int
	column int
}

func NewPolyanetService(psp PolyanetServiceParams) PolyanetService {
	return &PolyanetServiceParams{
		ApiUrl: psp.ApiUrl,
		CandidateId: psp.CandidateId,
	}
}

func (p *PolyanetServiceParams) GeneratePolyantesCross(ctx context.Context, request *http.Request) (*domain.Matrix, error) {
	var mc *domain.MatrixConfiguration 
	// read request body.
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil  {
		return nil, err
	}
	// unmarshal request body.
	err = json.Unmarshal(requestBody, &mc)
	if err != nil  {
		return nil, err
	}
	// TODO: delete prints.
	fmt.Println(p.ApiUrl)
	fmt.Println("(" + time.Now().Format("2006-01-02 15:04:05") + ") received params to generate polyanet cross: ", mc)
	// validate request body.
	if err := validateConfiguration(mc.Rows, mc.Columns, mc.Offset); err != nil {
		return nil, err
	}
	// generates the polyantes matrix (cross).
	data := make([][]int, mc.Rows)
	for i := 0; i < mc.Rows; i++ {
		data[i] = make([]int, mc.Columns)
		for j := 0; j < mc.Columns; j++ {
			if i > mc.Offset-1 && i < mc.Rows-mc.Offset && j > mc.Offset-1 && j < mc.Columns-mc.Offset {
				if j == i || j == mc.Columns-1-i {
					data[i][j] = 1
					p.CreatePolyanet(ctx, request, &PolyanetApiRequest{
						candidateId: p.CandidateId,
						row: i,
						column: j,
					})
				}
			}
		}
	}
	m := domain.Matrix{Data: data}
	print(m)
	return &m, nil
}

func (p *PolyanetServiceParams) DeletePolyanet(ctx context.Context, request *http.Request) (*domain.Polyanet, error) {
	var polyanet *PolyanetApiRequest 
	// read request body.
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil  {
		return nil, err
	}
	// unmarshal request body.
	err = json.Unmarshal(requestBody, &polyanet)
	if err != nil  {
		return nil, err
	}
	// crossmint candidate id
	polyanet.candidateId = p.CandidateId
	// TODO: delete prints.
	fmt.Println("(" + time.Now().Format("2006-01-02 15:04:05") + ") received request to delete polyanet: ", polyanet)
	
	return &domain.Polyanet{
		CandidateId: polyanet.candidateId,
		Row: polyanet.row,
		Column: polyanet.column,
	}, nil
}

func (p *PolyanetServiceParams) CreatePolyanet(ctx context.Context, request *http.Request, polyanet *PolyanetApiRequest) (*domain.Polyanet, error) {
	if polyanet == nil {
		// read request body.
		requestBody, err := ioutil.ReadAll(request.Body)
		if err != nil  {
			return nil, err
		}
		// unmarshal request body
		err = json.Unmarshal(requestBody, &polyanet)
		if err != nil  {
			return nil, err
		}
		// crossmint candidate id.
		polyanet.candidateId = p.CandidateId
		// TODO: delete prints.
		fmt.Println("(" + time.Now().Format("2006-01-02 15:04:05") + ") received request to create a polyanet: ", polyanet)
	}
	// polyanet received from generatePolyanetsCross method.
	if polyanet != nil {
		// TODO: delete prints.
		fmt.Println("(" + time.Now().Format("2006-01-02 15:04:05") + ") received request to generate polyanets cross: ", polyanet)
		// Calling Sleep method
	    time.Sleep(1 * time.Second) // TODO: refactor this.
	}
	// creates a new http request body.
	requestBody := strings.NewReader(
		fmt.Sprintf(
			`{"candidateId": "%s", "row": %d, "column": %d}`, 
			polyanet.candidateId, 
			polyanet.row, 
			polyanet.column,
		),
	)
	// create a new http request.
	request, err := http.NewRequest(http.MethodPost, p.ApiUrl + "/polyanets", requestBody)
	if err != nil {
		return nil, err
	}
	// add content type to header.
	request.Header.Add("Content-Type", "application/json")
	// send http request.
	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}
	// A defer statement defers the execution of a function until the surrounding function returns.
	defer response.Body.Close()
	// TODO: delete prints.
	fmt.Println("response status code: ",response.StatusCode)
	fmt.Println("response request method: ",response.Request.Method)

	// TODO: retry when requests fails.

	return &domain.Polyanet{
		CandidateId: polyanet.candidateId,
		Row: polyanet.row,
		Column: polyanet.column,
	}, nil
}