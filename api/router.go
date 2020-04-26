package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

// example for init the database:
//
//  DB, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/employees?charset=utf8&parseTime=true")
//  if err != nil {
//  	panic("failed to connect database: " + err.Error())
//  }
//  defer db.Close()

var DB *gorm.DB

func ConfigRouter() http.Handler {
	router := httprouter.New()
	configAttributeDefinitionsRouter(router)
	configCareSitesRouter(router)
	configCdmSourcesRouter(router)
	configConceptsRouter(router)
	configConceptAncestorsRouter(router)
	configConceptClassesRouter(router)
	configConceptRelationshipsRouter(router)
	configConceptSynonymsRouter(router)
	configConditionErasRouter(router)
	configConditionOccurrencesRouter(router)
	configCostsRouter(router)
	configDeviceExposuresRouter(router)
	configDomainsRouter(router)
	configDoseErasRouter(router)
	configDrugErasRouter(router)
	configDrugExposuresRouter(router)
	configDrugStrengthsRouter(router)
	configFactRelationshipsRouter(router)
	configLocationsRouter(router)
	configLocationHistoriesRouter(router)
	configMeasurementsRouter(router)
	configMetadataRouter(router)
	configNotesRouter(router)
	configNoteNlpsRouter(router)
	configObservationsRouter(router)
	configObservationPeriodsRouter(router)
	configPayerPlanPeriodsRouter(router)
	configPeopleRouter(router)
	configProcedureOccurrencesRouter(router)
	configProvidersRouter(router)
	configRelationshipsRouter(router)
	configSourceToConceptMapsRouter(router)
	configSpecimenRouter(router)
	configSurveyConductsRouter(router)
	configVisitDetailsRouter(router)
	configVisitOccurrencesRouter(router)
	configVocabulariesRouter(router)

	return router
}

func ConfigGinRouter(router gin.IRoutes) {
	configGinAttributeDefinitionsRouter(router)
	configGinCareSitesRouter(router)
	configGinCdmSourcesRouter(router)
	configGinConceptsRouter(router)
	configGinConceptAncestorsRouter(router)
	configGinConceptClassesRouter(router)
	configGinConceptRelationshipsRouter(router)
	configGinConceptSynonymsRouter(router)
	configGinConditionErasRouter(router)
	configGinConditionOccurrencesRouter(router)
	configGinCostsRouter(router)
	configGinDeviceExposuresRouter(router)
	configGinDomainsRouter(router)
	configGinDoseErasRouter(router)
	configGinDrugErasRouter(router)
	configGinDrugExposuresRouter(router)
	configGinDrugStrengthsRouter(router)
	configGinFactRelationshipsRouter(router)
	configGinLocationsRouter(router)
	configGinLocationHistoriesRouter(router)
	configGinMeasurementsRouter(router)
	configGinMetadataRouter(router)
	configGinNotesRouter(router)
	configGinNoteNlpsRouter(router)
	configGinObservationsRouter(router)
	configGinObservationPeriodsRouter(router)
	configGinPayerPlanPeriodsRouter(router)
	configGinPeopleRouter(router)
	configGinProcedureOccurrencesRouter(router)
	configGinProvidersRouter(router)
	configGinRelationshipsRouter(router)
	configGinSourceToConceptMapsRouter(router)
	configGinSpecimenRouter(router)
	configGinSurveyConductsRouter(router)
	configGinVisitDetailsRouter(router)
	configGinVisitOccurrencesRouter(router)
	configGinVocabulariesRouter(router)

	return
}

func ConverHttprouterToGin(f httprouter.Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params httprouter.Params
		_len := len(c.Params)
		if _len == 0 {
			params = nil
		} else {
			params = ((*[1 << 10]httprouter.Param)(unsafe.Pointer(&c.Params[0])))[:_len]
		}

		f(c.Writer, c.Request, params)
	}
}

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}
