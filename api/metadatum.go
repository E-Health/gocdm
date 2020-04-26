package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configMetadataRouter(router *httprouter.Router) {
	router.GET("/metadata", GetAllMetadata)
	router.POST("/metadata", AddMetadatum)
	router.GET("/metadata/:id", GetMetadatum)
	router.PUT("/metadata/:id", UpdateMetadatum)
	router.DELETE("/metadata/:id", DeleteMetadatum)
}

func configGinMetadataRouter(router gin.IRoutes) {
	router.GET("/metadata", ConverHttprouterToGin(GetAllMetadata))
	router.POST("/metadata", ConverHttprouterToGin(AddMetadatum))
	router.GET("/metadata/:id", ConverHttprouterToGin(GetMetadatum))
	router.PUT("/metadata/:id", ConverHttprouterToGin(UpdateMetadatum))
	router.DELETE("/metadata/:id", ConverHttprouterToGin(DeleteMetadatum))
}

func GetAllMetadata(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	metadata := []*model.Metadatum{}

	metadata_orm := DB.Model(&model.Metadatum{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		metadata_orm = metadata_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		metadata_orm = metadata_orm.Order(order)
	}

	if err = metadata_orm.Find(&metadata).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &metadata)
}

func GetMetadatum(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	metadatum := &model.Metadatum{}
	if DB.First(metadatum, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, metadatum)
}

func AddMetadatum(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	metadatum := &model.Metadatum{}

	if err := readJSON(r, metadatum); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(metadatum).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, metadatum)
}

func UpdateMetadatum(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	metadatum := &model.Metadatum{}
	if DB.First(metadatum, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Metadatum{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(metadatum, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(metadatum).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, metadatum)
}

func DeleteMetadatum(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	metadatum := &model.Metadatum{}

	if DB.First(metadatum, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(metadatum).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
