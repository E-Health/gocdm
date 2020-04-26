package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configConceptSynonymsRouter(router *httprouter.Router) {
	router.GET("/conceptsynonyms", GetAllConceptSynonyms)
	router.POST("/conceptsynonyms", AddConceptSynonym)
	router.GET("/conceptsynonyms/:id", GetConceptSynonym)
	router.PUT("/conceptsynonyms/:id", UpdateConceptSynonym)
	router.DELETE("/conceptsynonyms/:id", DeleteConceptSynonym)
}

func configGinConceptSynonymsRouter(router gin.IRoutes) {
	router.GET("/conceptsynonyms", ConverHttprouterToGin(GetAllConceptSynonyms))
	router.POST("/conceptsynonyms", ConverHttprouterToGin(AddConceptSynonym))
	router.GET("/conceptsynonyms/:id", ConverHttprouterToGin(GetConceptSynonym))
	router.PUT("/conceptsynonyms/:id", ConverHttprouterToGin(UpdateConceptSynonym))
	router.DELETE("/conceptsynonyms/:id", ConverHttprouterToGin(DeleteConceptSynonym))
}

func GetAllConceptSynonyms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	conceptsynonyms := []*model.ConceptSynonym{}

	conceptsynonyms_orm := DB.Model(&model.ConceptSynonym{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		conceptsynonyms_orm = conceptsynonyms_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		conceptsynonyms_orm = conceptsynonyms_orm.Order(order)
	}

	if err = conceptsynonyms_orm.Find(&conceptsynonyms).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &conceptsynonyms)
}

func GetConceptSynonym(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conceptsynonym := &model.ConceptSynonym{}
	if DB.First(conceptsynonym, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, conceptsynonym)
}

func AddConceptSynonym(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	conceptsynonym := &model.ConceptSynonym{}

	if err := readJSON(r, conceptsynonym); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(conceptsynonym).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conceptsynonym)
}

func UpdateConceptSynonym(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	conceptsynonym := &model.ConceptSynonym{}
	if DB.First(conceptsynonym, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.ConceptSynonym{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(conceptsynonym, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(conceptsynonym).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conceptsynonym)
}

func DeleteConceptSynonym(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conceptsynonym := &model.ConceptSynonym{}

	if DB.First(conceptsynonym, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(conceptsynonym).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
