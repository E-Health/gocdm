package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configDoseErasRouter(router *httprouter.Router) {
	router.GET("/doseeras", GetAllDoseEras)
	router.POST("/doseeras", AddDoseEra)
	router.GET("/doseeras/:id", GetDoseEra)
	router.PUT("/doseeras/:id", UpdateDoseEra)
	router.DELETE("/doseeras/:id", DeleteDoseEra)
}

func configGinDoseErasRouter(router gin.IRoutes) {
	router.GET("/doseeras", ConverHttprouterToGin(GetAllDoseEras))
	router.POST("/doseeras", ConverHttprouterToGin(AddDoseEra))
	router.GET("/doseeras/:id", ConverHttprouterToGin(GetDoseEra))
	router.PUT("/doseeras/:id", ConverHttprouterToGin(UpdateDoseEra))
	router.DELETE("/doseeras/:id", ConverHttprouterToGin(DeleteDoseEra))
}

func GetAllDoseEras(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	doseeras := []*model.DoseEra{}

	doseeras_orm := DB.Model(&model.DoseEra{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		doseeras_orm = doseeras_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		doseeras_orm = doseeras_orm.Order(order)
	}

	if err = doseeras_orm.Find(&doseeras).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &doseeras)
}

func GetDoseEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	doseera := &model.DoseEra{}
	if DB.First(doseera, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, doseera)
}

func AddDoseEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	doseera := &model.DoseEra{}

	if err := readJSON(r, doseera); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(doseera).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, doseera)
}

func UpdateDoseEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	doseera := &model.DoseEra{}
	if DB.First(doseera, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.DoseEra{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(doseera, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(doseera).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, doseera)
}

func DeleteDoseEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	doseera := &model.DoseEra{}

	if DB.First(doseera, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(doseera).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
