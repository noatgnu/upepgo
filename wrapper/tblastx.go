package wrapper

import (
	"os/exec"
	"log"
	"errors"
	"fmt"
	"encoding/xml"
	"os"
	"strconv"
	"io"
	"upepgo/models"
	"database/sql"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"upepgo/helper"
)

type TBlastXCommandline struct {
	Command string
	In      string
	DB      string
	Filter bool
	Out     string
	Evalue float64
	Outfmt int
	CommandLine
}

type TBlastXNode struct {
	XMLName xml.Name
	Content []byte `xml:",innerxml"`
	Nodes   []TBlastXNode `xml:",any"`
}

type TBlastXQuery struct {
	QueryID string
	Seq string
	Hits []TBlastXHit
}

type TBlastXHit struct {
	QueryID string
	HitID int64
	Hit *models.UpepSorfPosition
	Def string
	Accession string
	Length int
	Seq string
	Hsps []TBlastXHsp
}

type TBlastXHsp struct {
	BitScore float64
	Score int
	Evalue float64
	QueryStartPosition int
	QueryEndPosition int
	HitStartPosition int
	HitEndPosition int
	QueryFrame int
	HitFrame int
	Identity int
	Positive int
	Gaps int
	AlignLength int
	QuerySeq string
	HitSeq string
	Midline string
}

func (t *TBlastXCommandline) Execute() (err error) {
	commandArray, err := t.CommandBuild()
	if err != nil {
		return err
	}
	if commandArray != nil {
		cmd := exec.Command(commandArray[0], commandArray[1:]...)
		log.Println(cmd.Args)
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else {
		return errors.New("text: need parameters")
	}
	return err
}

func (t *TBlastXCommandline) CommandBuild() (commandArray []string, err error) {
	if t.Command != "" {
		commandArray = append(commandArray, t.Command)
	} else {
		commandArray = append(commandArray, "tblastx")
	}
	if t.DB != "" {
		commandArray = append(commandArray, "-db", t.DB)
	} else {
		return nil, errors.New("text: need db type")
	}
	if t.In != "" {
		commandArray = append(commandArray, "-query", t.In)
	} else {
		return nil, errors.New("text: need input file")
	}
	if t.Filter {
		commandArray = append(commandArray, "-F")
	}
	commandArray = append(commandArray, "-evalue", fmt.Sprintf("%v", t.Evalue), "-outfmt", fmt.Sprintf("%v", t.Outfmt))
	if t.Out != "" {
		commandArray = append(commandArray, "-out", t.Out)
	} else {
		return nil, errors.New("text: need output file")
	}

	return commandArray, nil
}

func ParseHitTBlastXOutputXML(path string, querySeqMap map[string]*helper.Sequence, db *sql.DB) []TBlastXQuery {
	f, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	c := make(chan TBlastXHit)
	m := make(map[int64]*models.UpepSorfPosition)
	queries := make([]TBlastXQuery, len(querySeqMap))
	queriesMapToArray := make(map[string]int)
	count := 0
	for k := range querySeqMap {
		queries[count] = TBlastXQuery{
			QueryID: k,
			Seq:     querySeqMap[k].Seq,
			Hits:    []TBlastXHit{},
		}
		queriesMapToArray[k] = count
		count++
	}
	go ParseFromXMLReader(f, c)
	for h := range c {
		if _, ok := m[h.HitID]; !ok {
			sorf, err := models.UpepSorfPositions(db, qm.Where("id=?", h.HitID), qm.Load("RefSeqEntry")).One()
			if err != nil {
				log.Panicln(err)
			}
			m[h.HitID] = sorf
		}
		h.Seq = m[h.HitID].R.RefSeqEntry.RefSeqSequence[m[h.HitID].StartingPosition-1:m[h.HitID].EndingPosition]
		if val, ok := queriesMapToArray[h.QueryID]; ok {
			queries[val].Hits = append(queries[val].Hits, h)
		}
	}
	f.Close()
	return queries[:]
}



func ParseFromXMLReader(f io.Reader,c chan TBlastXHit) {
	var node TBlastXNode
	decoder := xml.NewDecoder(f)
	err := decoder.Decode(&node)
	if err != nil {
		log.Panicln(err)
	}
	var hit TBlastXHit
	var hsp TBlastXHsp
	var currentQuery string
	walk([]TBlastXNode{node}, func(node TBlastXNode) bool {
		switch node.XMLName.Local {
		case "Iteration_query-def":
			currentQuery = string(node.Content)
		case "Hit":
			if hit.Hsps != nil {
				c <- hit
			}
			hit = TBlastXHit{}
			hit.QueryID = currentQuery
			hit.Hsps = []TBlastXHsp{}
		case "Hit_id":
			hit.HitID, err = strconv.ParseInt(string(node.Content), 10, 64)
			if err != nil {
				log.Panicln(err)
			}
		case "Hit_def":
			hit.Def = string(node.Content)
		case "Hit_accession":
			hit.Accession = string(node.Content)
		case "Hit_len":
			l, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hit.Length = l
		case "Hit_hsps":
			hit.Hsps = []TBlastXHsp{}
		case "Hsp":
			hsp = TBlastXHsp{}
		case "Hsp_bit-score":
			bs, err := strconv.ParseFloat(string(node.Content), 64)
			if err != nil {
				log.Panicln(err)
			}
			hsp.BitScore = bs
		case "Hsp_score":
			s, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.Score = s
		case "Hsp_evalue":
			e, err := strconv.ParseFloat(string(node.Content), 64)
			if err != nil {
				log.Panicln(err)
			}
			hsp.Evalue = e
		case "Hsp_query-from":
			e, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.QueryStartPosition = e
		case "Hsp_query-to":
			e, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.QueryEndPosition = e
		case "Hsp_hit-from":
			e, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.HitStartPosition = e
		case "Hsp_hit-to":
			e, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.HitEndPosition = e
		case "Hsp_query-frame":
			e, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.QueryFrame = e
		case "Hsp_hit-frame":
			e, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.HitFrame = e
		case "Hsp_identity":
			e, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.Identity = e
		case "Hsp_positive":
			e, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.Positive = e
		case "Hsp_gaps":
			e, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.Gaps = e
		case "Hsp_align-len":
			e, err := strconv.Atoi(string(node.Content))
			if err != nil {
				log.Panicln(err)
			}
			hsp.AlignLength = e
		case "Hsp_qseq":
			hsp.QuerySeq = string(node.Content)
		case "Hsp_hseq":
			hsp.HitSeq = string(node.Content)
		case "Hsp_midline":
			hsp.Midline = string(node.Content)
			hit.Hsps = append(hit.Hsps, hsp)
		}
		return true
	})
	close(c)
}

func walk(nodes []TBlastXNode, f func(TBlastXNode) bool) {
	for _, n := range nodes {
		if f(n) {
			walk(n.Nodes, f)
		}
	}
}