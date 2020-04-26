package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configSpecimenRouter(router *httprouter.Router) {
	router.GET("/specimen", GetAllSpecimen)
	router.POST("/specimen", AddSpeciman)
	router.GET("/specimen/:id", GetSpeciman)
	router.PUT("/specimen/:id", UpdateSpeciman)
	router.DELETE("/specimen/:id", DeleteSpeciman)
}

func configGinSpecimenRouter(router gin.IRoutes) {
	router.GET("/specimen", ConverHttprouterToGin(GetAllSpecimen))
	router.POST("/specimen", ConverHttprouterToGin(AddSpeciman))
	router.GET("/specimen/:id", ConverHttprouterToGin(GetSpeciman))
	router.PUT("/specimen/:id", ConverHttprouterToGin(UpdateSpeciman))
	router.DELETE("/specimen/:id", ConverHttprouterToGin(DeleteSpeciman))
}

func GetAllSpecimen(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	specimen := []*model.Speciman{}

	specimen_orm := DB.Model(&model.Speciman{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		specimen_orm = specimen_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		specimen_orm = specimen_orm.Order(order)
	}

	if err = specimen_orm.Find(&specimen).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &specimen)
}

func GetSpeciman(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	speciman := &model.Speciman{}
	if DB.First(speciman, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, speciman)
}

func AddSpeciman(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	speciman := &model.Speciman{}

	if err := readJSON(r, speciman); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(speciman).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, speciman)
}

func UpdateSpeciman(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	speciman := &model.Speciman{}
	if DB.First(speciman, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Speciman{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(speciman, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(speciman).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, speciman)
}

func DeleteSpeciman(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	speciman := &model.Speciman{}

	if DB.First(speciman, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(speciman).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
