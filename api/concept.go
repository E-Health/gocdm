package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configConceptsRouter(router *httprouter.Router) {
	router.GET("/concepts", GetAllConcepts)
	router.POST("/concepts", AddConcept)
	router.GET("/concepts/:id", GetConcept)
	router.PUT("/concepts/:id", UpdateConcept)
	router.DELETE("/concepts/:id", DeleteConcept)
}

func configGinConceptsRouter(router gin.IRoutes) {
	router.GET("/concepts", ConverHttprouterToGin(GetAllConcepts))
	router.POST("/concepts", ConverHttprouterToGin(AddConcept))
	router.GET("/concepts/:id", ConverHttprouterToGin(GetConcept))
	router.PUT("/concepts/:id", ConverHttprouterToGin(UpdateConcept))
	router.DELETE("/concepts/:id", ConverHttprouterToGin(DeleteConcept))
}

func GetAllConcepts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	concepts := []*model.Concept{}

	concepts_orm := DB.Model(&model.Concept{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		concepts_orm = concepts_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		concepts_orm = concepts_orm.Order(order)
	}

	if err = concepts_orm.Find(&concepts).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &concepts)
}

func GetConcept(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	concept := &model.Concept{}
	if DB.First(concept, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, concept)
}

func AddConcept(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	concept := &model.Concept{}

	if err := readJSON(r, concept); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(concept).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, concept)
}

func UpdateConcept(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	concept := &model.Concept{}
	if DB.First(concept, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Concept{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(concept, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(concept).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, concept)
}

func DeleteConcept(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	concept := &model.Concept{}

	if DB.First(concept, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(concept).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
