package sorf

import (
		"github.com/volatiletech/sqlboiler/queries/qm"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/noatgnu/upepgo/models"
)

func GetAllSorf(w http.ResponseWriter, vars map[string]string) {
	parameters := map[string]string{"startingCodon":"starting_codon_id", "stoppingCodon": "ending_codon_id", "organism": "ent.organism_id", "refseqname": "ent.name", "dbId": "ref_seq_db_id"}
	query := []qm.QueryMod{qm.InnerJoin("upep.upep_ref_seq_entries ent on ent.id = upep.upep_sorf_positions.ref_seq_entry_id"), qm.Load("StartingCodon"), qm.Load("EndingCodon"), qm.Load("RefSeqEntry.Organism")}
	startClause := false
	for k, v := range parameters {
		if k == "refseqname" {
			if vars[k] != "0" {
				if !startClause {
					query = append(query, qm.Where(v+"=?", vars[k]))
					startClause = true
				} else {
					query = append(query, qm.And(v+"=?", vars[k]))
				}
			}
		} else {
			value, err := strconv.Atoi(vars[k])
			if err != nil {
				log.Println(err)
			} else {
				if value != 0 {
					if !startClause {
						query = append(query, qm.Where(v+"=?", value))
						startClause = true
					} else {
						query = append(query, qm.And(v+"=?", value))
					}
				}
			}
		}
	}

	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		limit = 20
	}
	query = append(query, qm.Limit(limit))
	offset, err := strconv.Atoi(vars["offset"])
	if err != nil {
		offset = 0
	}
	query = append(query, qm.Offset(offset))

	sorfs, err := models.UpepSorfPositionsG(query...).All()
	if err != nil {
		log.Panicln(err)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(sorfs)
}