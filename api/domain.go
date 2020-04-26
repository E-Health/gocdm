package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configDomainsRouter(router *httprouter.Router) {
	router.GET("/domains", GetAllDomains)
	router.POST("/domains", AddDomain)
	router.GET("/domains/:id", GetDomain)
	router.PUT("/domains/:id", UpdateDomain)
	router.DELETE("/domains/:id", DeleteDomain)
}

func configGinDomainsRouter(router gin.IRoutes) {
	router.GET("/domains", ConverHttprouterToGin(GetAllDomains))
	router.POST("/domains", ConverHttprouterToGin(AddDomain))
	router.GET("/domains/:id", ConverHttprouterToGin(GetDomain))
	router.PUT("/domains/:id", ConverHttprouterToGin(UpdateDomain))
	router.DELETE("/domains/:id", ConverHttprouterToGin(DeleteDomain))
}

func GetAllDomains(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	domains := []*model.Domain{}

	domains_orm := DB.Model(&model.Domain{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		domains_orm = domains_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		domains_orm = domains_orm.Order(order)
	}

	if err = domains_orm.Find(&domains).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &domains)
}

func GetDomain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	domain := &model.Domain{}
	if DB.First(domain, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, domain)
}

func AddDomain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	domain := &model.Domain{}

	if err := readJSON(r, domain); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(domain).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, domain)
}

func UpdateDomain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	domain := &model.Domain{}
	if DB.First(domain, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Domain{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(domain, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(domain).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, domain)
}

func DeleteDomain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	domain := &model.Domain{}

	if DB.First(domain, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(domain).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
