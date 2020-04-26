package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configDrugStrengthsRouter(router *httprouter.Router) {
	router.GET("/drugstrengths", GetAllDrugStrengths)
	router.POST("/drugstrengths", AddDrugStrength)
	router.GET("/drugstrengths/:id", GetDrugStrength)
	router.PUT("/drugstrengths/:id", UpdateDrugStrength)
	router.DELETE("/drugstrengths/:id", DeleteDrugStrength)
}

func configGinDrugStrengthsRouter(router gin.IRoutes) {
	router.GET("/drugstrengths", ConverHttprouterToGin(GetAllDrugStrengths))
	router.POST("/drugstrengths", ConverHttprouterToGin(AddDrugStrength))
	router.GET("/drugstrengths/:id", ConverHttprouterToGin(GetDrugStrength))
	router.PUT("/drugstrengths/:id", ConverHttprouterToGin(UpdateDrugStrength))
	router.DELETE("/drugstrengths/:id", ConverHttprouterToGin(DeleteDrugStrength))
}

func GetAllDrugStrengths(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	drugstrengths := []*model.DrugStrength{}

	drugstrengths_orm := DB.Model(&model.DrugStrength{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		drugstrengths_orm = drugstrengths_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		drugstrengths_orm = drugstrengths_orm.Order(order)
	}

	if err = drugstrengths_orm.Find(&drugstrengths).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &drugstrengths)
}

func GetDrugStrength(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	drugstrength := &model.DrugStrength{}
	if DB.First(drugstrength, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, drugstrength)
}

func AddDrugStrength(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	drugstrength := &model.DrugStrength{}

	if err := readJSON(r, drugstrength); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(drugstrength).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, drugstrength)
}

func UpdateDrugStrength(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	drugstrength := &model.DrugStrength{}
	if DB.First(drugstrength, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.DrugStrength{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(drugstrength, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(drugstrength).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, drugstrength)
}

func DeleteDrugStrength(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	drugstrength := &model.DrugStrength{}

	if DB.First(drugstrength, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(drugstrength).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
