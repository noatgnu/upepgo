package codons

import (
		"log"
	"net/http"
	"encoding/json"
	"upepgo/models"
)

func GetAllCodons(w http.ResponseWriter) {
	codons, err := models.UpepCodonsG().All()
	if err != nil {
		log.Println(err)
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(codons)
	if err != nil {
		log.Println(err)
	}
}
