package main

import (
	"io"
	"log"
	"net/http"
	"nhl-data/nhlapi"
)

func main() {
	teams, err := nhlapi.GetAllTeams()
	if err != nil {
		log.Fatalf("error while getting all teams: %v", err)
	}

	nhlhandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "NHL Team Names")
		for _, team := range teams {
			io.WriteString(w, team.Teamname)
		}
	}

	http.HandleFunc("/nhl/teams", nhlhandler)
	log.Println("Listening for requests at http://localhost:8000/nhl/teams")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
