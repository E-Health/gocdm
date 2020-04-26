package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configVocabulariesRouter(router *httprouter.Router) {
	router.GET("/vocabularies", GetAllVocabularies)
	router.POST("/vocabularies", AddVocabulary)
	router.GET("/vocabularies/:id", GetVocabulary)
	router.PUT("/vocabularies/:id", UpdateVocabulary)
	router.DELETE("/vocabularies/:id", DeleteVocabulary)
}

func configGinVocabulariesRouter(router gin.IRoutes) {
	router.GET("/vocabularies", ConverHttprouterToGin(GetAllVocabularies))
	router.POST("/vocabularies", ConverHttprouterToGin(AddVocabulary))
	router.GET("/vocabularies/:id", ConverHttprouterToGin(GetVocabulary))
	router.PUT("/vocabularies/:id", ConverHttprouterToGin(UpdateVocabulary))
	router.DELETE("/vocabularies/:id", ConverHttprouterToGin(DeleteVocabulary))
}

func GetAllVocabularies(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	vocabularies := []*model.Vocabulary{}

	vocabularies_orm := DB.Model(&model.Vocabulary{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		vocabularies_orm = vocabularies_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		vocabularies_orm = vocabularies_orm.Order(order)
	}

	if err = vocabularies_orm.Find(&vocabularies).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &vocabularies)
}

func GetVocabulary(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	vocabulary := &model.Vocabulary{}
	if DB.First(vocabulary, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, vocabulary)
}

func AddVocabulary(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vocabulary := &model.Vocabulary{}

	if err := readJSON(r, vocabulary); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(vocabulary).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, vocabulary)
}

func UpdateVocabulary(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	vocabulary := &model.Vocabulary{}
	if DB.First(vocabulary, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Vocabulary{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(vocabulary, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(vocabulary).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, vocabulary)
}

func DeleteVocabulary(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	vocabulary := &model.Vocabulary{}

	if DB.First(vocabulary, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(vocabulary).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
