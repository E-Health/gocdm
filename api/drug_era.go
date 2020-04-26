package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configDrugErasRouter(router *httprouter.Router) {
	router.GET("/drugeras", GetAllDrugEras)
	router.POST("/drugeras", AddDrugEra)
	router.GET("/drugeras/:id", GetDrugEra)
	router.PUT("/drugeras/:id", UpdateDrugEra)
	router.DELETE("/drugeras/:id", DeleteDrugEra)
}

func configGinDrugErasRouter(router gin.IRoutes) {
	router.GET("/drugeras", ConverHttprouterToGin(GetAllDrugEras))
	router.POST("/drugeras", ConverHttprouterToGin(AddDrugEra))
	router.GET("/drugeras/:id", ConverHttprouterToGin(GetDrugEra))
	router.PUT("/drugeras/:id", ConverHttprouterToGin(UpdateDrugEra))
	router.DELETE("/drugeras/:id", ConverHttprouterToGin(DeleteDrugEra))
}

func GetAllDrugEras(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	drugeras := []*model.DrugEra{}

	drugeras_orm := DB.Model(&model.DrugEra{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		drugeras_orm = drugeras_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		drugeras_orm = drugeras_orm.Order(order)
	}

	if err = drugeras_orm.Find(&drugeras).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &drugeras)
}

func GetDrugEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	drugera := &model.DrugEra{}
	if DB.First(drugera, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, drugera)
}

func AddDrugEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	drugera := &model.DrugEra{}

	if err := readJSON(r, drugera); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(drugera).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, drugera)
}

func UpdateDrugEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	drugera := &model.DrugEra{}
	if DB.First(drugera, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.DrugEra{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(drugera, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(drugera).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, drugera)
}

func DeleteDrugEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	drugera := &model.DrugEra{}

	if DB.First(drugera, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(drugera).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
