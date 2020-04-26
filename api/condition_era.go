package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configConditionErasRouter(router *httprouter.Router) {
	router.GET("/conditioneras", GetAllConditionEras)
	router.POST("/conditioneras", AddConditionEra)
	router.GET("/conditioneras/:id", GetConditionEra)
	router.PUT("/conditioneras/:id", UpdateConditionEra)
	router.DELETE("/conditioneras/:id", DeleteConditionEra)
}

func configGinConditionErasRouter(router gin.IRoutes) {
	router.GET("/conditioneras", ConverHttprouterToGin(GetAllConditionEras))
	router.POST("/conditioneras", ConverHttprouterToGin(AddConditionEra))
	router.GET("/conditioneras/:id", ConverHttprouterToGin(GetConditionEra))
	router.PUT("/conditioneras/:id", ConverHttprouterToGin(UpdateConditionEra))
	router.DELETE("/conditioneras/:id", ConverHttprouterToGin(DeleteConditionEra))
}

func GetAllConditionEras(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	conditioneras := []*model.ConditionEra{}

	conditioneras_orm := DB.Model(&model.ConditionEra{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		conditioneras_orm = conditioneras_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		conditioneras_orm = conditioneras_orm.Order(order)
	}

	if err = conditioneras_orm.Find(&conditioneras).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &conditioneras)
}

func GetConditionEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conditionera := &model.ConditionEra{}
	if DB.First(conditionera, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, conditionera)
}

func AddConditionEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	conditionera := &model.ConditionEra{}

	if err := readJSON(r, conditionera); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(conditionera).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conditionera)
}

func UpdateConditionEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	conditionera := &model.ConditionEra{}
	if DB.First(conditionera, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.ConditionEra{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(conditionera, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(conditionera).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conditionera)
}

func DeleteConditionEra(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conditionera := &model.ConditionEra{}

	if DB.First(conditionera, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(conditionera).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
