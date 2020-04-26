package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configMeasurementsRouter(router *httprouter.Router) {
	router.GET("/measurements", GetAllMeasurements)
	router.POST("/measurements", AddMeasurement)
	router.GET("/measurements/:id", GetMeasurement)
	router.PUT("/measurements/:id", UpdateMeasurement)
	router.DELETE("/measurements/:id", DeleteMeasurement)
}

func configGinMeasurementsRouter(router gin.IRoutes) {
	router.GET("/measurements", ConverHttprouterToGin(GetAllMeasurements))
	router.POST("/measurements", ConverHttprouterToGin(AddMeasurement))
	router.GET("/measurements/:id", ConverHttprouterToGin(GetMeasurement))
	router.PUT("/measurements/:id", ConverHttprouterToGin(UpdateMeasurement))
	router.DELETE("/measurements/:id", ConverHttprouterToGin(DeleteMeasurement))
}

func GetAllMeasurements(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	measurements := []*model.Measurement{}

	measurements_orm := DB.Model(&model.Measurement{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		measurements_orm = measurements_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		measurements_orm = measurements_orm.Order(order)
	}

	if err = measurements_orm.Find(&measurements).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &measurements)
}

func GetMeasurement(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	measurement := &model.Measurement{}
	if DB.First(measurement, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, measurement)
}

func AddMeasurement(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	measurement := &model.Measurement{}

	if err := readJSON(r, measurement); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(measurement).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, measurement)
}

func UpdateMeasurement(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	measurement := &model.Measurement{}
	if DB.First(measurement, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Measurement{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(measurement, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(measurement).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, measurement)
}

func DeleteMeasurement(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	measurement := &model.Measurement{}

	if DB.First(measurement, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(measurement).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
