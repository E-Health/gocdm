package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configVisitDetailsRouter(router *httprouter.Router) {
	router.GET("/visitdetails", GetAllVisitDetails)
	router.POST("/visitdetails", AddVisitDetail)
	router.GET("/visitdetails/:id", GetVisitDetail)
	router.PUT("/visitdetails/:id", UpdateVisitDetail)
	router.DELETE("/visitdetails/:id", DeleteVisitDetail)
}

func configGinVisitDetailsRouter(router gin.IRoutes) {
	router.GET("/visitdetails", ConverHttprouterToGin(GetAllVisitDetails))
	router.POST("/visitdetails", ConverHttprouterToGin(AddVisitDetail))
	router.GET("/visitdetails/:id", ConverHttprouterToGin(GetVisitDetail))
	router.PUT("/visitdetails/:id", ConverHttprouterToGin(UpdateVisitDetail))
	router.DELETE("/visitdetails/:id", ConverHttprouterToGin(DeleteVisitDetail))
}

func GetAllVisitDetails(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	visitdetails := []*model.VisitDetail{}

	visitdetails_orm := DB.Model(&model.VisitDetail{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		visitdetails_orm = visitdetails_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		visitdetails_orm = visitdetails_orm.Order(order)
	}

	if err = visitdetails_orm.Find(&visitdetails).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &visitdetails)
}

func GetVisitDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	visitdetail := &model.VisitDetail{}
	if DB.First(visitdetail, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, visitdetail)
}

func AddVisitDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	visitdetail := &model.VisitDetail{}

	if err := readJSON(r, visitdetail); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(visitdetail).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, visitdetail)
}

func UpdateVisitDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	visitdetail := &model.VisitDetail{}
	if DB.First(visitdetail, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.VisitDetail{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(visitdetail, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(visitdetail).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, visitdetail)
}

func DeleteVisitDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	visitdetail := &model.VisitDetail{}

	if DB.First(visitdetail, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(visitdetail).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
