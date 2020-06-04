package main

import (
	"fmt"
	"io"
	"log"

	"encoding/csv"
	"os"

	"github.com/E-Health/gocdm/api"
	"github.com/E-Health/gocdm/model"
	"github.com/jinzhu/gorm"

	// "github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db, err := gorm.Open("sqlite3", "cdmtest.db")
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Init completed")
	}
	defer db.Close()

	//Pass it to api
	api.DB = db

	// Migrate the schema
	db.AutoMigrate(&model.AttributeDefinition{})
	db.AutoMigrate(&model.CareSite{})
	db.AutoMigrate(&model.CdmSource{})
	db.AutoMigrate(&model.ConceptAncestor{})
	db.AutoMigrate(&model.ConceptClass{})
	db.AutoMigrate(&model.Concept{})
	db.AutoMigrate(&model.ConceptRelationship{})
	db.AutoMigrate(&model.ConceptSynonym{})
	db.AutoMigrate(&model.ConditionEra{})
	db.AutoMigrate(&model.ConditionOccurrence{})
	db.AutoMigrate(&model.Cost{})
	db.AutoMigrate(&model.DeviceExposure{})
	db.AutoMigrate(&model.Domain{})
	db.AutoMigrate(&model.DoseEra{})
	db.AutoMigrate(&model.DrugEra{})
	db.AutoMigrate(&model.DrugExposure{})
	db.AutoMigrate(&model.DrugStrength{})
	db.AutoMigrate(&model.FactRelationship{})
	db.AutoMigrate(&model.Location{})
	db.AutoMigrate(&model.LocationHistory{})
	db.AutoMigrate(&model.Measurement{})
	db.AutoMigrate(&model.Metadatum{})
	db.AutoMigrate(&model.Note{})
	db.AutoMigrate(&model.NoteNlp{})
	db.AutoMigrate(&model.Observation{})
	db.AutoMigrate(&model.ObservationPeriod{})
	db.AutoMigrate(&model.PayerPlanPeriod{})
	db.AutoMigrate(&model.Person{})
	db.AutoMigrate(&model.ProcedureOccurrence{})
	db.AutoMigrate(&model.Provider{})
	db.AutoMigrate(&model.Relationship{})
	db.AutoMigrate(&model.SourceToConceptMap{})
	db.AutoMigrate(&model.Speciman{})
	db.AutoMigrate(&model.SurveyConduct{})
	db.AutoMigrate(&model.VisitDetail{})
	db.AutoMigrate(&model.VisitOccurrence{})
	db.AutoMigrate(&model.Vocabulary{})
}

func main() {

	// Create
	db.Create(&model.Person{YearOfBirth: 2010})
	yob := model.Person{YearOfBirth: 2010}
	name := api.GetNameForTest()
	fmt.Println(yob.TableName(), name)
}

func csvToMap(reader io.Reader) []map[string]string {
	r := csv.NewReader(reader)
	r.Comma = '\t' // Use tab-delimited instead of comma <---- here!
	rows := []map[string]string{}
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}
	return rows
}

func vocab(tsvPath string) {

	csvFile, err := os.Open(tsvPath)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	vocabAll := csvToMap(csvFile)

	for v := range vocabAll {
		fmt.Println(v)
	}

}
