package service

import (
	"fmt"
	domain "megaverse/domain"
	"net/http"
	"strings"
)

var Client HTTPClient

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// init() is a very special function meant to execute prior to the main() function
func init() {
	Client = &http.Client{}
}

func print(m domain.Matrix) {
	for _, row := range m.Data {
		fmt.Println(row)
	}
}

func validateConfiguration(rows, columns, offset int) error {
	var errorString []string

	if rows <= 0 {
		errorRows := fmt.Errorf("invalid rows number: %d", rows)
		errorString = append(errorString, errorRows.Error())
	}

	if columns <= 0 {
		errorColumns := fmt.Errorf("invalid columns number: %d", columns)
		errorString = append(errorString, errorColumns.Error())
	}

	if offset < 0 {
		errorOffset := fmt.Errorf("invalid offset number: %d", offset)
		errorString = append(errorString, errorOffset.Error())
	}

	if len(errorString) > 0 {
		return fmt.Errorf(strings.Join(errorString, ", "))
	}

	return nil
}

func megaverseMap() map[string]string {
	megaverse := make(map[string]string)
	megaverse["space"] = "space"
	megaverse["polyanet"] = "polyanet"
	megaverse["right_cometh"] = "right_cometh"
	megaverse["left_cometh"] = "left_cometh"
	megaverse["up_cometh"] = "up_cometh"
	megaverse["down_cometh"] = "down_cometh"
	megaverse["white_soloon"] = "white_soloon"
	megaverse["blue_soloon"] = "blue_soloon"
	megaverse["purple_soloon"] = "purple_soloon"
	megaverse["red_soloon"] = "red_soloon"
	return megaverse
}

func megaverseParametersMap(celestialBody string) map[string]string{
	megaverseParameters := make(map[string]string)
	if celestialBody == "polyanet" {
		megaverseParameters["name"] = "polyanets"
	}
	if celestialBody == "right_cometh" {
		megaverseParameters["name"] = "comeths"
		megaverseParameters["parameter"] = "direction"
		megaverseParameters["direction"] = "right"
	}
	if celestialBody == "left_cometh" {
		megaverseParameters["name"] = "comeths"
		megaverseParameters["parameter"] = "direction"
		megaverseParameters["direction"] = "left"
	}
	if celestialBody == "up_cometh" {
		megaverseParameters["name"] = "comeths"
		megaverseParameters["parameter"] = "direction"
		megaverseParameters["direction"] = "up"
	}
	if celestialBody == "down_cometh" {
		megaverseParameters["name"] = "comeths"
		megaverseParameters["parameter"] = "direction"
		megaverseParameters["direction"] = "down"
	}
	if celestialBody == "white_soloon" {
		megaverseParameters["name"] = "soloons"
		megaverseParameters["parameter"] = "color"
		megaverseParameters["color"] = "white"
	}
	if celestialBody == "blue_soloon" {
		megaverseParameters["name"] = "soloons"
		megaverseParameters["parameter"] = "color"
		megaverseParameters["color"] = "blue"
	}
	if celestialBody == "purple_soloon" {
		megaverseParameters["name"] = "soloons"
		megaverseParameters["parameter"] = "color"
		megaverseParameters["color"] = "purple"
	}
	if celestialBody == "red_soloon" {
		megaverseParameters["name"] = "soloons"
		megaverseParameters["parameter"] = "color"
		megaverseParameters["color"] = "red"
	}
	return megaverseParameters
}

func (p *PolyanetServiceParams) requestBody(parameterType, parameterValue string, row, column int) *strings.Reader {
	// creates a new http request body.
	requestBody := strings.NewReader(
		fmt.Sprintf(
			`{"candidateId": "%s", "%s": "%s", "row": %d, "column": %d}`, 
			p.Candidate.CandidateId,
			parameterType,
			parameterValue,
			row, 
			column,
		),
	)
	if len(parameterType) == 0 {
		requestBody = strings.NewReader(
			fmt.Sprintf(
				`{"candidateId": "%s", "row": %d, "column": %d}`, 
				p.Candidate.CandidateId,
				row, 
				column,
			),
		)
	}

	return requestBody
}