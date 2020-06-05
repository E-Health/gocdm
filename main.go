package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/E-Health/gocdm/api"
	"github.com/E-Health/gocdm/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	db, err := gorm.Open("sqlite3", "cdmtest.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

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

	// Create
	db.Create(&model.Person{YearOfBirth: 2010})
	fmt.Println("Database created")

	// Read
	var person model.Person
	db.First(&person, 1)                         // find person with id 1
	db.First(&person, "year_of_birth = ?", 2010) // find person with YearOfBirth 2010
	fmt.Println("Read worked")

	// Update - update person's YearOfBirth to 2000
	db.Model(&person).Update("YearOfBirth", 2000)
	fmt.Println("Update worked")

	// Delete - delete person
	db.Delete(&person)
	fmt.Println("Delete worked")
	// Refer https://gorm.io/docs/ for details

	// Expose API
	router := api.ConfigRouter()
	fmt.Println("Starting server on port 8085")
	log.Fatal(http.ListenAndServe("0.0.0.0:8085", router))
	// see api/router.go to see how to use Gin instead of http
}
