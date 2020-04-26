package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configObservationsRouter(router *httprouter.Router) {
	router.GET("/observations", GetAllObservations)
	router.POST("/observations", AddObservation)
	router.GET("/observations/:id", GetObservation)
	router.PUT("/observations/:id", UpdateObservation)
	router.DELETE("/observations/:id", DeleteObservation)
}

func configGinObservationsRouter(router gin.IRoutes) {
	router.GET("/observations", ConverHttprouterToGin(GetAllObservations))
	router.POST("/observations", ConverHttprouterToGin(AddObservation))
	router.GET("/observations/:id", ConverHttprouterToGin(GetObservation))
	router.PUT("/observations/:id", ConverHttprouterToGin(UpdateObservation))
	router.DELETE("/observations/:id", ConverHttprouterToGin(DeleteObservation))
}

func GetAllObservations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	observations := []*model.Observation{}

	observations_orm := DB.Model(&model.Observation{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		observations_orm = observations_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		observations_orm = observations_orm.Order(order)
	}

	if err = observations_orm.Find(&observations).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &observations)
}

func GetObservation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	observation := &model.Observation{}
	if DB.First(observation, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, observation)
}

func AddObservation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	observation := &model.Observation{}

	if err := readJSON(r, observation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(observation).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, observation)
}

func UpdateObservation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	observation := &model.Observation{}
	if DB.First(observation, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Observation{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(observation, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(observation).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, observation)
}

func DeleteObservation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	observation := &model.Observation{}

	if DB.First(observation, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(observation).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
