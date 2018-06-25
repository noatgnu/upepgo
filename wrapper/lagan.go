package wrapper

import (
	"os/exec"
	"log"
	"errors"
	"os"
	"bufio"
	"io"
	"strings"
	"github.com/noatgnu/upepgo/helper"
)

type LaganDBCommandline struct {
	Command string
	In1      string
	In2 string
	Out     string
	CommandLine
}

func (m *LaganDBCommandline) Execute() (err error) {
	commandArray, err := m.CommandBuild()
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

func (m *LaganDBCommandline) CommandBuild() (commandArray []string, err error) {
	if m.Command != "" {
		commandArray = append(commandArray, m.Command)
	} else {
		commandArray = append(commandArray, "perl", "lagan")
	}
	if m.In1 != "" {
		commandArray = append(commandArray, m.In1)
	} else {
		return nil, errors.New("text: need input file")
	}
	if m.In2 != "" {
		commandArray = append(commandArray, m.In2)
	} else {
		return nil, errors.New("text: need second input file")
	}
	if m.Out != "" {
		commandArray = append(commandArray, "-out", m.Out)
	} else {
		return nil, errors.New("text: need output file")
	}

	return commandArray, nil
}

func ParseLaganCMDOutPut(filep string) helper.LaganAlignment {
	f, err := os.Open(filep)
	if err != nil {
		log.Panicln(err)
	}
	reader := bufio.NewReader(f)
	inSeq := false
	upep := ""
	midLine := ""
	query := ""
	for {
		r, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicln(err)
		}

		if r != "" {
			if strings.HasPrefix(r, "seq1") {
				inSeq = true
				upep += r[9:]
			} else if strings.HasPrefix(r, "seq2") {
				inSeq = false
				query += r[9:]
			} else {
				if inSeq {
					midLine += r[9:]
				}
			}
		}
	}
	f.Close()
	return helper.LaganAlignment{Target: upep, MidLine: midLine, Query: query}
}
