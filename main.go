package main

import (
	"log"
	"megaverse/domain"
	handler "megaverse/handler"
	service "megaverse/service"

	"github.com/joho/godotenv"
)


func main() {
	// godotenv package to get .env variables.
	envFile, err := loadEnv()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("Error loading environment variables file")
		return
	}
	// config data.
	port := envFile["PORT"]
	//url := envFile["BASEURL"]
	crossmintApiUrl := envFile["CROSSMINTAPIURL"]
	candidateId := envFile["CANDIDATEID"]
	
	// Initialize the polyanet service paramns.
	polyanetServiceParams := service.PolyanetServiceParams{
		ApiUrl: crossmintApiUrl,
		Candidate: domain.Candidate{
			CandidateId: candidateId,
		},
	}

	// Initialize the polyanet service.
	polyanetService := service.NewPolyanetService(polyanetServiceParams)

	// Initialize the polyanet handler.
	polyanetHandler := handler.NewPolyanetHandler(polyanetService)
	polyanetHandler.PolyanetHandlers()

	// Initialize the server.
	server := NewServer(port)

	// log.
	log.Fatal(server.Start())
}

func loadEnv() (envMap map[string]string, err error) {
	envFile, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}
	return envFile, nil
}