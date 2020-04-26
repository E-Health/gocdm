package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configAttributeDefinitionsRouter(router *httprouter.Router) {
	router.GET("/attributedefinitions", GetAllAttributeDefinitions)
	router.POST("/attributedefinitions", AddAttributeDefinition)
	router.GET("/attributedefinitions/:id", GetAttributeDefinition)
	router.PUT("/attributedefinitions/:id", UpdateAttributeDefinition)
	router.DELETE("/attributedefinitions/:id", DeleteAttributeDefinition)
}

func configGinAttributeDefinitionsRouter(router gin.IRoutes) {
	router.GET("/attributedefinitions", ConverHttprouterToGin(GetAllAttributeDefinitions))
	router.POST("/attributedefinitions", ConverHttprouterToGin(AddAttributeDefinition))
	router.GET("/attributedefinitions/:id", ConverHttprouterToGin(GetAttributeDefinition))
	router.PUT("/attributedefinitions/:id", ConverHttprouterToGin(UpdateAttributeDefinition))
	router.DELETE("/attributedefinitions/:id", ConverHttprouterToGin(DeleteAttributeDefinition))
}

func GetAllAttributeDefinitions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	attributedefinitions := []*model.AttributeDefinition{}

	attributedefinitions_orm := DB.Model(&model.AttributeDefinition{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		attributedefinitions_orm = attributedefinitions_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		attributedefinitions_orm = attributedefinitions_orm.Order(order)
	}

	if err = attributedefinitions_orm.Find(&attributedefinitions).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &attributedefinitions)
}

func GetAttributeDefinition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	attributedefinition := &model.AttributeDefinition{}
	if DB.First(attributedefinition, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, attributedefinition)
}

func AddAttributeDefinition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	attributedefinition := &model.AttributeDefinition{}

	if err := readJSON(r, attributedefinition); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(attributedefinition).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, attributedefinition)
}

func UpdateAttributeDefinition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	attributedefinition := &model.AttributeDefinition{}
	if DB.First(attributedefinition, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.AttributeDefinition{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(attributedefinition, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(attributedefinition).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, attributedefinition)
}

func DeleteAttributeDefinition(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	attributedefinition := &model.AttributeDefinition{}

	if DB.First(attributedefinition, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(attributedefinition).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
