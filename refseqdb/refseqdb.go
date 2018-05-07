package refseqdb

import (
	"path/filepath"
	"os"
	"log"
	"compress/gzip"
	"bufio"
	"strings"
	"regexp"
	"io"
	"strconv"
	"github.com/jlaffaye/ftp"
	"database/sql"
	"upepgo/models"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"gopkg.in/volatiletech/null.v6"
)

var reLocus = regexp.MustCompile(`([\w\-]+)`)
var reFeature = regexp.MustCompile(`(\w+)\s+(<*\d+)\.\.(>*\d+)`)
var reAccession = regexp.MustCompile(`(\w+)`)
var host = "ftp.ncbi.nlm.nih.gov:21"
var base = "/refseq/release/"
var releaseMap = map[string]string{
	"Complete": "complete",
	"Fungi": "fungi",
	"Invertebrate": "invertebrate",
	"Plant": "plant",
	"Mammalian": "vertebrate_mammalian",
	"Vertebrate Other": "vertebrate_other",
	"Bacteria": "bacteria",
}

var initCodons = map[string][]string{
	"start": []string{"AUG", "CUG", "UUG", "GUG", "ACG", "AUA", "AUU", "AAG", "AGG"},
	"end": []string{"TAG", "TAA", "TGA"},
}

type RefSInf struct {
	RefSeq models.UpepRefSeqEntry
	Features []models.UpepFeature
	MolecularType models.UpepMolecularType
	Organism models.UpepOrganism
	Accession models.UpepAccession
	GI models.UpepGeneIdentifier
}


func Truncate(db *sql.DB) {
	r := `TRUNCATE TABLE upep.upep_accessions
    RESTART IDENTITY
    CASCADE;
TRUNCATE TABLE upep.upep_blast_db
    RESTART IDENTITY
    CASCADE;
TRUNCATE TABLE upep.upep_codon
    RESTART IDENTITY
    CASCADE;
TRUNCATE TABLE upep.upep_features
    RESTART IDENTITY
    CASCADE;
TRUNCATE TABLE upep.upep_gene_identifiers
    RESTART IDENTITY
    CASCADE;
TRUNCATE TABLE upep.upep_molecular_types
    RESTART IDENTITY
    CASCADE;
TRUNCATE TABLE upep.upep_organisms
    RESTART IDENTITY
    CASCADE;
TRUNCATE TABLE upep.upep_ref_seq_db
    RESTART IDENTITY
    CASCADE;
TRUNCATE TABLE upep.upep_ref_seq_entries
    RESTART IDENTITY
    CASCADE;
TRUNCATE TABLE upep.upep_sorf_positions
    RESTART IDENTITY
    CASCADE;`
    db.Exec(r)
}

func DownloadRefSeqDB(dbList []string, db *sql.DB) chan *ftp.Entry{
	EnChan := make(chan *ftp.Entry)
	client, err := ftp.Dial(host)
	if err != nil {
		log.Panicln(err)
	}
	if err := client.Login("anonymous", "toan.phung@uq.net.au"); err != nil {
		log.Panicln(err)
	}
	version := GetRemoteReleaseVersion(client)
	log.Println(version)
	go func() {
		localDBs := GetLocalDBVersion(db, dbList)
		for i := range localDBs {
			if localDBs[i].Version < version {
				fileList, err := client.List(base + releaseMap[localDBs[i].Name])
				if err != nil {
					log.Panicln(err)
				}
				for f := range fileList {
					if strings.HasSuffix(fileList[f].Name, "rna.gbff.gz") {
						EnChan <- fileList[f]
					}
				}
			}
		}
		client.Logout()
		client.Quit()
		close(EnChan)
	}()
	return EnChan
}

func InitCodons(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		log.Panicln(err)
		return err
	}
	for _, v := range initCodons["start"] {
		c := models.UpepCodon{Sequence:v, StartingCodon:true}
		err := c.Insert(tx)
		if err != nil {
			tx.Rollback()
			log.Panicln(err)
			return err
		}
	}
	for _, v := range initCodons["end"] {
		c := models.UpepCodon{Sequence:v, EndingCodon:true}
		err := c.Insert(tx)
		if err != nil {
			tx.Rollback()
			log.Panicln(err)
			return err
		}
	}
	tx.Commit()
	return err
}

func GetAllCodons(db *sql.DB) ([]*models.UpepCodon, []*models.UpepCodon) {
	startingCods, err := models.UpepCodons(db, qm.Where("starting_codon = ?", true)).All()
	if err != nil {
		log.Panicln(err)
	}
	endingCods, err := models.UpepCodons(db, qm.Where("ending_codon = ?", true)).All()
	if err != nil {
		log.Panicln(err)
	}
	return startingCods, endingCods
}

func GetDownloadedRefSeqDB (rootPath string) []*models.UpepRefSeqDB {
	var RefSeqList []*models.UpepRefSeqDB
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "rna.gbff.gz") {
			RefSeqList = append(RefSeqList, &models.UpepRefSeqDB{})
		}
		return nil
	})
	if err != nil {
		log.Panicln(err)
	}
	return RefSeqList
}

func ReadGZText(path string) chan string {
	c := make(chan string)
	f, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}

	gzReader, err := gzip.NewReader(f)
	if err != nil {
		log.Panicln(err)
	}
	reader := bufio.NewReader(gzReader)

	go func(){
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Panicln(err)
			}
			c <- strings.TrimRight(line, "\n")
		}
		close(c)
		f.Close()
	}()
	return c
}

func ParseRefSeqDB(path string, source int64) chan RefSInf {
	sqlChan := make(chan RefSInf)
	go func() {
		seqFlag := false
		featureFlag := false
		rsChan := ReadGZText(path)

		var res RefSInf
		res.RefSeq = models.UpepRefSeqEntry{}
		for r := range rsChan {
			if strings.TrimSpace(r) == "//" {
				seqFlag = false

				if res.MolecularType.Name == "mRNA" {
					res.RefSeq.RefSeqSequence = strings.ToUpper(res.RefSeq.RefSeqSequence)
					res.RefSeq.RefSeqDBID = null.NewInt64(source, true)
				}
				sqlChan <- res
				res = RefSInf{}
				res.RefSeq = models.UpepRefSeqEntry{}
			} else if seqFlag {
				res.RefSeq.RefSeqSequence += strings.Replace(r[10:], " ", "", -1)
			} else {
				if strings.HasPrefix(r, "LOCUS") {
					locus := reLocus.FindAllString(r, -1)
					res.RefSeq.Name = null.NewString(locus[1], true)
					res.MolecularType.Name = locus[4]
				} else if strings.HasPrefix(r, "ACCESSION") {
					acc := reAccession.FindAllString(r, -1)
					res.Accession.Accession = acc[1]
				} else if strings.HasPrefix(r, "VERSION") {
					if strings.Contains(r, "GI") {
						s := strings.Index(r, "GI:")
						GI, err := strconv.ParseInt(r[s:s+3], 10, 64)
						if err != nil {
							log.Panicln(err)
						}
						res.GI.Gi = GI
					}
				} else if strings.HasPrefix(r, "  ORGANISM") {
					res.Organism.Name = strings.TrimSpace(r[12:])

				} else if strings.HasPrefix(r, "ORIGIN") {

					seqFlag = true
					featureFlag = false
				} else if strings.HasPrefix(r, "FEATURES") {
					featureFlag = true
				} else if featureFlag {
					if !strings.HasPrefix(r, "      ") {
						feature := reFeature.FindAllStringSubmatch(r[5:], -1)
						if len(feature) > 0 {
							f := models.UpepFeature{Name: strings.ToUpper(feature[0][1])}
							if strings.Contains(feature[0][2], "<") {
								start, err := strconv.Atoi(strings.TrimPrefix(feature[0][2], "<"))
								if err != nil {
									log.Panicln(err)
								}
								f.FeatureStart = start
								f.PartialStart = true
							} else {
								start, err := strconv.Atoi(feature[0][2])
								if err != nil {
									log.Panicln(err)
								}
								f.FeatureStart = start
							}
							if strings.Contains(feature[0][3], ">") {
								end, err := strconv.Atoi(strings.TrimPrefix(feature[0][3], ">"))
								if err != nil {
									log.Panicln(err)
								}
								f.FeatureEnd = end
								f.PartialEnd = true
							} else {
								end, err := strconv.Atoi(feature[0][3])
								if err != nil {
									log.Panicln(err)
								}
								f.FeatureEnd = end
							}
							res.Features = append(res.Features, f)
						}
					}
				}
			}
		}
		close(sqlChan)
	}()
	return sqlChan
}

func ReadRefSeqDB(path string, source int64, db *sql.DB, startingCodons map[string]*models.UpepCodon, endingCodons map[string]*models.UpepCodon) {
	sqlChan := ParseRefSeqDB(path, source)

	molecularTypeMap := make(map[string]models.UpepMolecularType)
	organismMap := make(map[string]models.UpepOrganism)
	count := 0
	for res := range sqlChan {
		tx, err := db.Begin()
		if err != nil {
			log.Panicln(err)
		}
		if val, ok := molecularTypeMap[res.MolecularType.Name]; ok {
			res.RefSeq.MolecularTypeID = null.NewInt64(val.ID, true)
		} else {
			err := res.MolecularType.Insert(db)
			if err != nil {
				tx.Rollback()
				log.Panicln(res.MolecularType)
			}
			res.RefSeq.MolecularTypeID = null.NewInt64(res.MolecularType.ID, true)
			molecularTypeMap[res.MolecularType.Name] = res.MolecularType
		}
		if val, ok := organismMap[res.Organism.Name]; ok {
			res.RefSeq.OrganismID = null.NewInt64(val.ID, true)
		} else {
			err := res.Organism.Insert(db)
			if err != nil {
				tx.Rollback()
				log.Panicln(res.Organism)
			}
			res.RefSeq.OrganismID = null.NewInt64(res.Organism.ID, true)
			organismMap[res.Organism.Name] = res.Organism
		}
		if res.Accession.Accession != "" {
			err = res.Accession.Insert(db)
			if err != nil {
				tx.Rollback()
				log.Panicln(res.Accession)
			}
			res.RefSeq.AccessionID = null.NewInt64(res.Accession.ID, true)
		}

		if res.GI.Gi != 0 {
			err = res.GI.Insert(db)
			if err != nil {
				tx.Rollback()
				log.Panicln(res.GI)
			}
			res.RefSeq.GiID = null.NewInt64(res.GI.ID, true)
		}
		res.RefSeq.RefSeqDBID = null.NewInt64(source, true)

		err = res.RefSeq.Insert(db)
		if err != nil {
			tx.Rollback()
			log.Panicln(res.RefSeq)
		}
		var CDS *models.UpepFeature
		for i := range res.Features {
			res.Features[i].RefSeqEntryID = res.RefSeq.ID
			if res.Features[i].Name == "CDS" {
				CDS = &res.Features[i]
			}
			res.Features[i].Insert(tx)
		}
		tx.Commit()
		if CDS.Name == "CDS" {
			cUpep := ParseUpep(res.RefSeq, CDS.FeatureEnd, startingCodons, endingCodons, db)
			tx, err := db.Begin()
			if err != nil {
				tx.Rollback()
				log.Panicln(res.RefSeq)
			}
			for c := range cUpep {
				sorfLen := c.EndingPosition-c.StartingPosition
				if (c.StartingPosition != CDS.FeatureStart) && (15<=sorfLen) && (sorfLen <=600){
					err := c.Insert(tx)
					if err != nil {
						tx.Rollback()
						log.Println(c.RefSeqEntryID)
						log.Println(c)
						log.Panicln(err)

					}
				}
			}
			tx.Commit()
		}
		count ++
	}
}


func GetLocalDBVersion(db *sql.DB, checkSelect []string) (localDBs []*models.UpepRefSeqDB) {
	tx, err := db.Begin()
	if err != nil {
		log.Panicln(err)
	}

	for i := range checkSelect {
		udb, err := models.UpepRefSeqDBS(tx, qm.Where("name=?", checkSelect[i])).One()
		if err != nil {
			log.Panicln(err)
		}
		localDBs = append(localDBs, udb)
	}

	return localDBs
}

func GetRemoteReleaseVersion(client *ftp.ServerConn) (ver int) {
	DownloadToTemp(client, "RELEASE_NUMBER")

	o, err := os.Open("temp/RELEASE_NUMBER")
	if err != nil {
		log.Panicln(err)
	}
	scanner := bufio.NewScanner(o)

	for scanner.Scan() {
		ver, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			log.Panicln(err)
		}
		break
	}
	return ver
}

func DownloadToTemp(client *ftp.ServerConn, fileName string) {
	log.Printf("Downloading %v", fileName)
	rn, err := client.Retr(base + fileName)
	if err != nil {
		log.Panicln(err)
	}
	defer rn.Close()
	f, err := os.Create("temp/" + fileName)
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()
	n, err := io.Copy(f, rn)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("Downloaded %v (%v bytes)", fileName,n)
}

func ParseUpep(refseq models.UpepRefSeqEntry, searchEnd int, startingCodons map[string]*models.UpepCodon, endingCodons map[string]*models.UpepCodon, db *sql.DB) chan *models.UpepSorfPosition{
	c := make(chan *models.UpepSorfPosition)
	go func() {
		frames := make(map[int]models.UpepSorfPosition)
		for i:= 0; i < 3; i++ {
			frames[i] = models.UpepSorfPosition{}
		}
		for i := 0; i < (searchEnd-3); i+=3 {
			for k, _ := range frames {
				start := i+k
				end := i+k+3
				if end < searchEnd {
					if val, ok := startingCodons[refseq.RefSeqSequence[start:end]]; ok {
						frames[k] =  models.UpepSorfPosition{
							StartingPosition: start+1,
							RefSeqEntryID:    refseq.ID,
							StartingCodonID:  val.ID,
						}
					} else if val, ok := endingCodons[refseq.RefSeqSequence[start:end]]; ok {
						if frames[k].StartingPosition != 0 {
							f := frames[k]
							f.EndingPosition = end
							f.EndingCodonID = val.ID
							c <- &f
							frames[k] = models.UpepSorfPosition{}
						}
					}
				} else {
					break
				}
			}
		}
		close(c)
	}()
	return c
}