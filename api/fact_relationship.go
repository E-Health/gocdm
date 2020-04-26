package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configFactRelationshipsRouter(router *httprouter.Router) {
	router.GET("/factrelationships", GetAllFactRelationships)
	router.POST("/factrelationships", AddFactRelationship)
	router.GET("/factrelationships/:id", GetFactRelationship)
	router.PUT("/factrelationships/:id", UpdateFactRelationship)
	router.DELETE("/factrelationships/:id", DeleteFactRelationship)
}

func configGinFactRelationshipsRouter(router gin.IRoutes) {
	router.GET("/factrelationships", ConverHttprouterToGin(GetAllFactRelationships))
	router.POST("/factrelationships", ConverHttprouterToGin(AddFactRelationship))
	router.GET("/factrelationships/:id", ConverHttprouterToGin(GetFactRelationship))
	router.PUT("/factrelationships/:id", ConverHttprouterToGin(UpdateFactRelationship))
	router.DELETE("/factrelationships/:id", ConverHttprouterToGin(DeleteFactRelationship))
}

func GetAllFactRelationships(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	factrelationships := []*model.FactRelationship{}

	factrelationships_orm := DB.Model(&model.FactRelationship{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		factrelationships_orm = factrelationships_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		factrelationships_orm = factrelationships_orm.Order(order)
	}

	if err = factrelationships_orm.Find(&factrelationships).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &factrelationships)
}

func GetFactRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	factrelationship := &model.FactRelationship{}
	if DB.First(factrelationship, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, factrelationship)
}

func AddFactRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	factrelationship := &model.FactRelationship{}

	if err := readJSON(r, factrelationship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(factrelationship).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, factrelationship)
}

func UpdateFactRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	factrelationship := &model.FactRelationship{}
	if DB.First(factrelationship, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.FactRelationship{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(factrelationship, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(factrelationship).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, factrelationship)
}

func DeleteFactRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	factrelationship := &model.FactRelationship{}

	if DB.First(factrelationship, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(factrelationship).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
