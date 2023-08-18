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