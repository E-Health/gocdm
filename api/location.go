package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configLocationsRouter(router *httprouter.Router) {
	router.GET("/locations", GetAllLocations)
	router.POST("/locations", AddLocation)
	router.GET("/locations/:id", GetLocation)
	router.PUT("/locations/:id", UpdateLocation)
	router.DELETE("/locations/:id", DeleteLocation)
}

func configGinLocationsRouter(router gin.IRoutes) {
	router.GET("/locations", ConverHttprouterToGin(GetAllLocations))
	router.POST("/locations", ConverHttprouterToGin(AddLocation))
	router.GET("/locations/:id", ConverHttprouterToGin(GetLocation))
	router.PUT("/locations/:id", ConverHttprouterToGin(UpdateLocation))
	router.DELETE("/locations/:id", ConverHttprouterToGin(DeleteLocation))
}

func GetAllLocations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	locations := []*model.Location{}

	locations_orm := DB.Model(&model.Location{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		locations_orm = locations_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		locations_orm = locations_orm.Order(order)
	}

	if err = locations_orm.Find(&locations).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &locations)
}

func GetLocation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	location := &model.Location{}
	if DB.First(location, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, location)
}

func AddLocation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	location := &model.Location{}

	if err := readJSON(r, location); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(location).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, location)
}

func UpdateLocation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	location := &model.Location{}
	if DB.First(location, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Location{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(location, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(location).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, location)
}

func DeleteLocation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	location := &model.Location{}

	if DB.First(location, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(location).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
