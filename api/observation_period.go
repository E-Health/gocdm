package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configObservationPeriodsRouter(router *httprouter.Router) {
	router.GET("/observationperiods", GetAllObservationPeriods)
	router.POST("/observationperiods", AddObservationPeriod)
	router.GET("/observationperiods/:id", GetObservationPeriod)
	router.PUT("/observationperiods/:id", UpdateObservationPeriod)
	router.DELETE("/observationperiods/:id", DeleteObservationPeriod)
}

func configGinObservationPeriodsRouter(router gin.IRoutes) {
	router.GET("/observationperiods", ConverHttprouterToGin(GetAllObservationPeriods))
	router.POST("/observationperiods", ConverHttprouterToGin(AddObservationPeriod))
	router.GET("/observationperiods/:id", ConverHttprouterToGin(GetObservationPeriod))
	router.PUT("/observationperiods/:id", ConverHttprouterToGin(UpdateObservationPeriod))
	router.DELETE("/observationperiods/:id", ConverHttprouterToGin(DeleteObservationPeriod))
}

func GetAllObservationPeriods(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	observationperiods := []*model.ObservationPeriod{}

	observationperiods_orm := DB.Model(&model.ObservationPeriod{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		observationperiods_orm = observationperiods_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		observationperiods_orm = observationperiods_orm.Order(order)
	}

	if err = observationperiods_orm.Find(&observationperiods).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &observationperiods)
}

func GetObservationPeriod(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	observationperiod := &model.ObservationPeriod{}
	if DB.First(observationperiod, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, observationperiod)
}

func AddObservationPeriod(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	observationperiod := &model.ObservationPeriod{}

	if err := readJSON(r, observationperiod); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(observationperiod).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, observationperiod)
}

func UpdateObservationPeriod(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	observationperiod := &model.ObservationPeriod{}
	if DB.First(observationperiod, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.ObservationPeriod{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(observationperiod, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(observationperiod).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, observationperiod)
}

func DeleteObservationPeriod(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	observationperiod := &model.ObservationPeriod{}

	if DB.First(observationperiod, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(observationperiod).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
