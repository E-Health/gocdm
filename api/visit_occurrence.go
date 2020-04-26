package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configVisitOccurrencesRouter(router *httprouter.Router) {
	router.GET("/visitoccurrences", GetAllVisitOccurrences)
	router.POST("/visitoccurrences", AddVisitOccurrence)
	router.GET("/visitoccurrences/:id", GetVisitOccurrence)
	router.PUT("/visitoccurrences/:id", UpdateVisitOccurrence)
	router.DELETE("/visitoccurrences/:id", DeleteVisitOccurrence)
}

func configGinVisitOccurrencesRouter(router gin.IRoutes) {
	router.GET("/visitoccurrences", ConverHttprouterToGin(GetAllVisitOccurrences))
	router.POST("/visitoccurrences", ConverHttprouterToGin(AddVisitOccurrence))
	router.GET("/visitoccurrences/:id", ConverHttprouterToGin(GetVisitOccurrence))
	router.PUT("/visitoccurrences/:id", ConverHttprouterToGin(UpdateVisitOccurrence))
	router.DELETE("/visitoccurrences/:id", ConverHttprouterToGin(DeleteVisitOccurrence))
}

func GetAllVisitOccurrences(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	visitoccurrences := []*model.VisitOccurrence{}

	visitoccurrences_orm := DB.Model(&model.VisitOccurrence{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		visitoccurrences_orm = visitoccurrences_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		visitoccurrences_orm = visitoccurrences_orm.Order(order)
	}

	if err = visitoccurrences_orm.Find(&visitoccurrences).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &visitoccurrences)
}

func GetVisitOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	visitoccurrence := &model.VisitOccurrence{}
	if DB.First(visitoccurrence, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, visitoccurrence)
}

func AddVisitOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	visitoccurrence := &model.VisitOccurrence{}

	if err := readJSON(r, visitoccurrence); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(visitoccurrence).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, visitoccurrence)
}

func UpdateVisitOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	visitoccurrence := &model.VisitOccurrence{}
	if DB.First(visitoccurrence, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.VisitOccurrence{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(visitoccurrence, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(visitoccurrence).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, visitoccurrence)
}

func DeleteVisitOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	visitoccurrence := &model.VisitOccurrence{}

	if DB.First(visitoccurrence, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(visitoccurrence).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
