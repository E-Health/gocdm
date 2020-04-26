package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configSourceToConceptMapsRouter(router *httprouter.Router) {
	router.GET("/sourcetoconceptmaps", GetAllSourceToConceptMaps)
	router.POST("/sourcetoconceptmaps", AddSourceToConceptMap)
	router.GET("/sourcetoconceptmaps/:id", GetSourceToConceptMap)
	router.PUT("/sourcetoconceptmaps/:id", UpdateSourceToConceptMap)
	router.DELETE("/sourcetoconceptmaps/:id", DeleteSourceToConceptMap)
}

func configGinSourceToConceptMapsRouter(router gin.IRoutes) {
	router.GET("/sourcetoconceptmaps", ConverHttprouterToGin(GetAllSourceToConceptMaps))
	router.POST("/sourcetoconceptmaps", ConverHttprouterToGin(AddSourceToConceptMap))
	router.GET("/sourcetoconceptmaps/:id", ConverHttprouterToGin(GetSourceToConceptMap))
	router.PUT("/sourcetoconceptmaps/:id", ConverHttprouterToGin(UpdateSourceToConceptMap))
	router.DELETE("/sourcetoconceptmaps/:id", ConverHttprouterToGin(DeleteSourceToConceptMap))
}

func GetAllSourceToConceptMaps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	sourcetoconceptmaps := []*model.SourceToConceptMap{}

	sourcetoconceptmaps_orm := DB.Model(&model.SourceToConceptMap{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		sourcetoconceptmaps_orm = sourcetoconceptmaps_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		sourcetoconceptmaps_orm = sourcetoconceptmaps_orm.Order(order)
	}

	if err = sourcetoconceptmaps_orm.Find(&sourcetoconceptmaps).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &sourcetoconceptmaps)
}

func GetSourceToConceptMap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	sourcetoconceptmap := &model.SourceToConceptMap{}
	if DB.First(sourcetoconceptmap, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, sourcetoconceptmap)
}

func AddSourceToConceptMap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sourcetoconceptmap := &model.SourceToConceptMap{}

	if err := readJSON(r, sourcetoconceptmap); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(sourcetoconceptmap).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, sourcetoconceptmap)
}

func UpdateSourceToConceptMap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	sourcetoconceptmap := &model.SourceToConceptMap{}
	if DB.First(sourcetoconceptmap, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.SourceToConceptMap{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(sourcetoconceptmap, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(sourcetoconceptmap).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, sourcetoconceptmap)
}

func DeleteSourceToConceptMap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	sourcetoconceptmap := &model.SourceToConceptMap{}

	if DB.First(sourcetoconceptmap, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(sourcetoconceptmap).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
