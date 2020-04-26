package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configRelationshipsRouter(router *httprouter.Router) {
	router.GET("/relationships", GetAllRelationships)
	router.POST("/relationships", AddRelationship)
	router.GET("/relationships/:id", GetRelationship)
	router.PUT("/relationships/:id", UpdateRelationship)
	router.DELETE("/relationships/:id", DeleteRelationship)
}

func configGinRelationshipsRouter(router gin.IRoutes) {
	router.GET("/relationships", ConverHttprouterToGin(GetAllRelationships))
	router.POST("/relationships", ConverHttprouterToGin(AddRelationship))
	router.GET("/relationships/:id", ConverHttprouterToGin(GetRelationship))
	router.PUT("/relationships/:id", ConverHttprouterToGin(UpdateRelationship))
	router.DELETE("/relationships/:id", ConverHttprouterToGin(DeleteRelationship))
}

func GetAllRelationships(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	relationships := []*model.Relationship{}

	relationships_orm := DB.Model(&model.Relationship{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		relationships_orm = relationships_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		relationships_orm = relationships_orm.Order(order)
	}

	if err = relationships_orm.Find(&relationships).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &relationships)
}

func GetRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	relationship := &model.Relationship{}
	if DB.First(relationship, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, relationship)
}

func AddRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	relationship := &model.Relationship{}

	if err := readJSON(r, relationship); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(relationship).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, relationship)
}

func UpdateRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	relationship := &model.Relationship{}
	if DB.First(relationship, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Relationship{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(relationship, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(relationship).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, relationship)
}

func DeleteRelationship(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	relationship := &model.Relationship{}

	if DB.First(relationship, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(relationship).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
