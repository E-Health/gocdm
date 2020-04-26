package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configCareSitesRouter(router *httprouter.Router) {
	router.GET("/caresites", GetAllCareSites)
	router.POST("/caresites", AddCareSite)
	router.GET("/caresites/:id", GetCareSite)
	router.PUT("/caresites/:id", UpdateCareSite)
	router.DELETE("/caresites/:id", DeleteCareSite)
}

func configGinCareSitesRouter(router gin.IRoutes) {
	router.GET("/caresites", ConverHttprouterToGin(GetAllCareSites))
	router.POST("/caresites", ConverHttprouterToGin(AddCareSite))
	router.GET("/caresites/:id", ConverHttprouterToGin(GetCareSite))
	router.PUT("/caresites/:id", ConverHttprouterToGin(UpdateCareSite))
	router.DELETE("/caresites/:id", ConverHttprouterToGin(DeleteCareSite))
}

func GetAllCareSites(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	caresites := []*model.CareSite{}

	caresites_orm := DB.Model(&model.CareSite{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		caresites_orm = caresites_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		caresites_orm = caresites_orm.Order(order)
	}

	if err = caresites_orm.Find(&caresites).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &caresites)
}

func GetCareSite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	caresite := &model.CareSite{}
	if DB.First(caresite, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, caresite)
}

func AddCareSite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	caresite := &model.CareSite{}

	if err := readJSON(r, caresite); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(caresite).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, caresite)
}

func UpdateCareSite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	caresite := &model.CareSite{}
	if DB.First(caresite, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.CareSite{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(caresite, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(caresite).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, caresite)
}

func DeleteCareSite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	caresite := &model.CareSite{}

	if DB.First(caresite, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(caresite).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
