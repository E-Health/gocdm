package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configDrugExposuresRouter(router *httprouter.Router) {
	router.GET("/drugexposures", GetAllDrugExposures)
	router.POST("/drugexposures", AddDrugExposure)
	router.GET("/drugexposures/:id", GetDrugExposure)
	router.PUT("/drugexposures/:id", UpdateDrugExposure)
	router.DELETE("/drugexposures/:id", DeleteDrugExposure)
}

func configGinDrugExposuresRouter(router gin.IRoutes) {
	router.GET("/drugexposures", ConverHttprouterToGin(GetAllDrugExposures))
	router.POST("/drugexposures", ConverHttprouterToGin(AddDrugExposure))
	router.GET("/drugexposures/:id", ConverHttprouterToGin(GetDrugExposure))
	router.PUT("/drugexposures/:id", ConverHttprouterToGin(UpdateDrugExposure))
	router.DELETE("/drugexposures/:id", ConverHttprouterToGin(DeleteDrugExposure))
}

func GetAllDrugExposures(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	drugexposures := []*model.DrugExposure{}

	drugexposures_orm := DB.Model(&model.DrugExposure{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		drugexposures_orm = drugexposures_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		drugexposures_orm = drugexposures_orm.Order(order)
	}

	if err = drugexposures_orm.Find(&drugexposures).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &drugexposures)
}

func GetDrugExposure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	drugexposure := &model.DrugExposure{}
	if DB.First(drugexposure, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, drugexposure)
}

func AddDrugExposure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	drugexposure := &model.DrugExposure{}

	if err := readJSON(r, drugexposure); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(drugexposure).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, drugexposure)
}

func UpdateDrugExposure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	drugexposure := &model.DrugExposure{}
	if DB.First(drugexposure, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.DrugExposure{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(drugexposure, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(drugexposure).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, drugexposure)
}

func DeleteDrugExposure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	drugexposure := &model.DrugExposure{}

	if DB.First(drugexposure, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(drugexposure).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
