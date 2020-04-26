package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configCostsRouter(router *httprouter.Router) {
	router.GET("/costs", GetAllCosts)
	router.POST("/costs", AddCost)
	router.GET("/costs/:id", GetCost)
	router.PUT("/costs/:id", UpdateCost)
	router.DELETE("/costs/:id", DeleteCost)
}

func configGinCostsRouter(router gin.IRoutes) {
	router.GET("/costs", ConverHttprouterToGin(GetAllCosts))
	router.POST("/costs", ConverHttprouterToGin(AddCost))
	router.GET("/costs/:id", ConverHttprouterToGin(GetCost))
	router.PUT("/costs/:id", ConverHttprouterToGin(UpdateCost))
	router.DELETE("/costs/:id", ConverHttprouterToGin(DeleteCost))
}

func GetAllCosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	costs := []*model.Cost{}

	costs_orm := DB.Model(&model.Cost{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		costs_orm = costs_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		costs_orm = costs_orm.Order(order)
	}

	if err = costs_orm.Find(&costs).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &costs)
}

func GetCost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	cost := &model.Cost{}
	if DB.First(cost, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, cost)
}

func AddCost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cost := &model.Cost{}

	if err := readJSON(r, cost); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(cost).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, cost)
}

func UpdateCost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	cost := &model.Cost{}
	if DB.First(cost, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Cost{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(cost, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(cost).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, cost)
}

func DeleteCost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	cost := &model.Cost{}

	if DB.First(cost, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(cost).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
