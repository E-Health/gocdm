package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configConceptRelationshipsRouter(router *httprouter.Router) {
	router.GET("/conceptrelationships", GetAllConceptRelationships)
	router.POST("/conceptrelationships", AddConceptRelationship)
	router.GET("/conceptrelationships/:id", GetConceptRelationship)
	router.PUT("/conceptrelationships/:id", UpdateConceptRelationship)
	router.DELETE("/conceptrelationships/:id", DeleteConceptRelationship)
}

func configGinConceptRelationshipsRouter(router gin.IRoutes) {
	router.GET("/conceptrelationships", ConverHttprouterToGin(GetAllConceptRelationships))
	router.POST("/conceptrelationships", ConverHttprouterToGin(AddConceptRelationship))
	router.GET("/conceptrelationships/:id", ConverHttprouterToGin(GetConceptRelationship))
	router.PUT("/conceptrelationships/:id", ConverHttprouterToGin(UpdateConceptRelationship))
	router.DELETE("/conceptrelationships/:id", ConverHttprouterToGin(DeleteConceptRelationship))
}

func GetAllConceptRelationships(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	conceptrelationships := []*model.ConceptRelationship{}

	conceptrelationships_orm := DB.Model(&model.ConceptRelationship{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		conceptrelationships_orm = conceptrelationships_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		conceptrelationships_orm = conceptrelationships_orm.Order(order)
	}

	if err = conceptrelationships_orm.Find(&conceptrelationships).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &conceptrelationships)
}

func GetConceptRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conceptrelationship := &model.ConceptRelationship{}
	if DB.First(conceptrelationship, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, conceptrelationship)
}

func AddConceptRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	conceptrelationship := &model.ConceptRelationship{}

	if err := readJSON(r, conceptrelationship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(conceptrelationship).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conceptrelationship)
}

func UpdateConceptRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	conceptrelationship := &model.ConceptRelationship{}
	if DB.First(conceptrelationship, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.ConceptRelationship{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(conceptrelationship, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(conceptrelationship).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conceptrelationship)
}

func DeleteConceptRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conceptrelationship := &model.ConceptRelationship{}

	if DB.First(conceptrelationship, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(conceptrelationship).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
