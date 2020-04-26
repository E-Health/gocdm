package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configConditionOccurrencesRouter(router *httprouter.Router) {
	router.GET("/conditionoccurrences", GetAllConditionOccurrences)
	router.POST("/conditionoccurrences", AddConditionOccurrence)
	router.GET("/conditionoccurrences/:id", GetConditionOccurrence)
	router.PUT("/conditionoccurrences/:id", UpdateConditionOccurrence)
	router.DELETE("/conditionoccurrences/:id", DeleteConditionOccurrence)
}

func configGinConditionOccurrencesRouter(router gin.IRoutes) {
	router.GET("/conditionoccurrences", ConverHttprouterToGin(GetAllConditionOccurrences))
	router.POST("/conditionoccurrences", ConverHttprouterToGin(AddConditionOccurrence))
	router.GET("/conditionoccurrences/:id", ConverHttprouterToGin(GetConditionOccurrence))
	router.PUT("/conditionoccurrences/:id", ConverHttprouterToGin(UpdateConditionOccurrence))
	router.DELETE("/conditionoccurrences/:id", ConverHttprouterToGin(DeleteConditionOccurrence))
}

func GetAllConditionOccurrences(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	conditionoccurrences := []*model.ConditionOccurrence{}

	conditionoccurrences_orm := DB.Model(&model.ConditionOccurrence{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		conditionoccurrences_orm = conditionoccurrences_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		conditionoccurrences_orm = conditionoccurrences_orm.Order(order)
	}

	if err = conditionoccurrences_orm.Find(&conditionoccurrences).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &conditionoccurrences)
}

func GetConditionOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conditionoccurrence := &model.ConditionOccurrence{}
	if DB.First(conditionoccurrence, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, conditionoccurrence)
}

func AddConditionOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	conditionoccurrence := &model.ConditionOccurrence{}

	if err := readJSON(r, conditionoccurrence); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(conditionoccurrence).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conditionoccurrence)
}

func UpdateConditionOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	conditionoccurrence := &model.ConditionOccurrence{}
	if DB.First(conditionoccurrence, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.ConditionOccurrence{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(conditionoccurrence, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(conditionoccurrence).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conditionoccurrence)
}

func DeleteConditionOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conditionoccurrence := &model.ConditionOccurrence{}

	if DB.First(conditionoccurrence, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(conditionoccurrence).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
