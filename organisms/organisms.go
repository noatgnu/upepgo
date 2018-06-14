package organisms

import (
		"encoding/json"
	"net/http"
	"strconv"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"log"
	"github.com/noatgnu/upepgo/models"
)

func GetAllOrganisms(w http.ResponseWriter, vars map[string]string) {
	dbId, err := strconv.Atoi(vars["dbId"])
	if err != nil {
		log.Panicln(err)
	}
	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		limit = 500
	}
	var organisms []*models.UpepOrganism
	if limit == 0 {
		organisms, err = models.UpepOrganismsG(qm.Where("upep_ref_seq_db_id=?", dbId), qm.OrderBy("name")).All()
		if err != nil {
			log.Panicln(err)
		}
	} else {
		organisms, err = models.UpepOrganismsG(qm.Where("upep_ref_seq_db_id=?", dbId), qm.Limit(limit), qm.OrderBy("name")).All()
		if err != nil {
			log.Panicln(err)
		}
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(organisms)
}
