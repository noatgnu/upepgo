package wrapper

import (
	"os/exec"
	"log"
	"errors"
	"fmt"
	"encoding/xml"
	"os"
	"strconv"
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

type TBlastXHit struct {
	ID string
	Def string
	Accession string
	Length int
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
		commandArray = append(commandArray, "-d", t.DB)
	} else {
		return nil, errors.New("text: need db type")
	}
	if t.In != "" {
		commandArray = append(commandArray, "-i", t.In)
	} else {
		return nil, errors.New("text: need input file")
	}
	if t.Filter {
		commandArray = append(commandArray, "-F")
	}
	commandArray = append(commandArray, "-e", fmt.Sprintf("%v", t.Evalue), "-m", fmt.Sprintf("%v", t.Outfmt))
	if t.Out != "" {
		commandArray = append(commandArray, "-o", t.Out)
	} else {
		return nil, errors.New("text: need output file")
	}

	return commandArray, nil
}

func ReadTBlastXOutput(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	var node TBlastXNode
	decoder := xml.NewDecoder(f)
	err = decoder.Decode(&node)
	if err != nil {
		log.Panicln(err)
	}
	var hit TBlastXHit
	var hsp TBlastXHsp
	walk([]TBlastXNode{node}, func(node TBlastXNode) bool {
		switch node.XMLName.Local {
		case "Hit":
			if hit.Hsps != nil {

			}
			hit = TBlastXHit{}
			hit.Hsps = []TBlastXHsp{}
		case "Hit_id":
			hit.ID = string(node.Content)
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
}

func walk(nodes []TBlastXNode, f func(TBlastXNode) bool) {
	for _, n := range nodes {
		if f(n) {
			walk(n.Nodes, f)
		}
	}
}