package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configLocationHistoriesRouter(router *httprouter.Router) {
	router.GET("/locationhistories", GetAllLocationHistories)
	router.POST("/locationhistories", AddLocationHistory)
	router.GET("/locationhistories/:id", GetLocationHistory)
	router.PUT("/locationhistories/:id", UpdateLocationHistory)
	router.DELETE("/locationhistories/:id", DeleteLocationHistory)
}

func configGinLocationHistoriesRouter(router gin.IRoutes) {
	router.GET("/locationhistories", ConverHttprouterToGin(GetAllLocationHistories))
	router.POST("/locationhistories", ConverHttprouterToGin(AddLocationHistory))
	router.GET("/locationhistories/:id", ConverHttprouterToGin(GetLocationHistory))
	router.PUT("/locationhistories/:id", ConverHttprouterToGin(UpdateLocationHistory))
	router.DELETE("/locationhistories/:id", ConverHttprouterToGin(DeleteLocationHistory))
}

func GetAllLocationHistories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	locationhistories := []*model.LocationHistory{}

	locationhistories_orm := DB.Model(&model.LocationHistory{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		locationhistories_orm = locationhistories_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		locationhistories_orm = locationhistories_orm.Order(order)
	}

	if err = locationhistories_orm.Find(&locationhistories).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &locationhistories)
}

func GetLocationHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	locationhistory := &model.LocationHistory{}
	if DB.First(locationhistory, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, locationhistory)
}

func AddLocationHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	locationhistory := &model.LocationHistory{}

	if err := readJSON(r, locationhistory); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(locationhistory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, locationhistory)
}

func UpdateLocationHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	locationhistory := &model.LocationHistory{}
	if DB.First(locationhistory, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.LocationHistory{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(locationhistory, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(locationhistory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, locationhistory)
}

func DeleteLocationHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	locationhistory := &model.LocationHistory{}

	if DB.First(locationhistory, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(locationhistory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
