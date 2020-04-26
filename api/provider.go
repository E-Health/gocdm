package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configProvidersRouter(router *httprouter.Router) {
	router.GET("/providers", GetAllProviders)
	router.POST("/providers", AddProvider)
	router.GET("/providers/:id", GetProvider)
	router.PUT("/providers/:id", UpdateProvider)
	router.DELETE("/providers/:id", DeleteProvider)
}

func configGinProvidersRouter(router gin.IRoutes) {
	router.GET("/providers", ConverHttprouterToGin(GetAllProviders))
	router.POST("/providers", ConverHttprouterToGin(AddProvider))
	router.GET("/providers/:id", ConverHttprouterToGin(GetProvider))
	router.PUT("/providers/:id", ConverHttprouterToGin(UpdateProvider))
	router.DELETE("/providers/:id", ConverHttprouterToGin(DeleteProvider))
}

func GetAllProviders(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	providers := []*model.Provider{}

	providers_orm := DB.Model(&model.Provider{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		providers_orm = providers_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		providers_orm = providers_orm.Order(order)
	}

	if err = providers_orm.Find(&providers).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &providers)
}

func GetProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	provider := &model.Provider{}
	if DB.First(provider, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, provider)
}

func AddProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	provider := &model.Provider{}

	if err := readJSON(r, provider); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(provider).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, provider)
}

func UpdateProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	provider := &model.Provider{}
	if DB.First(provider, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Provider{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(provider, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(provider).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, provider)
}

func DeleteProvider(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	provider := &model.Provider{}

	if DB.First(provider, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(provider).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
