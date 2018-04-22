package refseqdb

import (
	"testing"
	"strings"
)

func TestReadGZText(t *testing.T) {
	rsChan := ReadGZText("./test.gbff.gz")
	count := 0
	for r := range rsChan {
		if strings.HasPrefix(r, "LOCUS") {
			count++
		}
	}
	if count !=2 {
		t.Errorf("Expected %d, got %d", 2, count)
	}
}

func TestParseRefSeqDB(t *testing.T) {
	p := ParseRefSeqDB("test.gbff.gz", 1)
	acc := []string{"XM_006453678", "XM_006453679"}
	for r := range p {
		 result := func(ref RefSInf) bool {
			for _, v := range acc {
				if r.Accession.Accession == v {
					return true
				}
			}
			return false
		}(r)
		if !result {
			t.Error("Expected XM_006453678 or XM_006453679, got None")
		}

	}
}