package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configCdmSourcesRouter(router *httprouter.Router) {
	router.GET("/cdmsources", GetAllCdmSources)
	router.POST("/cdmsources", AddCdmSource)
	router.GET("/cdmsources/:id", GetCdmSource)
	router.PUT("/cdmsources/:id", UpdateCdmSource)
	router.DELETE("/cdmsources/:id", DeleteCdmSource)
}

func configGinCdmSourcesRouter(router gin.IRoutes) {
	router.GET("/cdmsources", ConverHttprouterToGin(GetAllCdmSources))
	router.POST("/cdmsources", ConverHttprouterToGin(AddCdmSource))
	router.GET("/cdmsources/:id", ConverHttprouterToGin(GetCdmSource))
	router.PUT("/cdmsources/:id", ConverHttprouterToGin(UpdateCdmSource))
	router.DELETE("/cdmsources/:id", ConverHttprouterToGin(DeleteCdmSource))
}

func GetAllCdmSources(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	cdmsources := []*model.CdmSource{}

	cdmsources_orm := DB.Model(&model.CdmSource{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		cdmsources_orm = cdmsources_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		cdmsources_orm = cdmsources_orm.Order(order)
	}

	if err = cdmsources_orm.Find(&cdmsources).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &cdmsources)
}

func GetCdmSource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	cdmsource := &model.CdmSource{}
	if DB.First(cdmsource, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, cdmsource)
}

func AddCdmSource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cdmsource := &model.CdmSource{}

	if err := readJSON(r, cdmsource); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(cdmsource).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, cdmsource)
}

func UpdateCdmSource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	cdmsource := &model.CdmSource{}
	if DB.First(cdmsource, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.CdmSource{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(cdmsource, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(cdmsource).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, cdmsource)
}

func DeleteCdmSource(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	cdmsource := &model.CdmSource{}

	if DB.First(cdmsource, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(cdmsource).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
