package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configPayerPlanPeriodsRouter(router *httprouter.Router) {
	router.GET("/payerplanperiods", GetAllPayerPlanPeriods)
	router.POST("/payerplanperiods", AddPayerPlanPeriod)
	router.GET("/payerplanperiods/:id", GetPayerPlanPeriod)
	router.PUT("/payerplanperiods/:id", UpdatePayerPlanPeriod)
	router.DELETE("/payerplanperiods/:id", DeletePayerPlanPeriod)
}

func configGinPayerPlanPeriodsRouter(router gin.IRoutes) {
	router.GET("/payerplanperiods", ConverHttprouterToGin(GetAllPayerPlanPeriods))
	router.POST("/payerplanperiods", ConverHttprouterToGin(AddPayerPlanPeriod))
	router.GET("/payerplanperiods/:id", ConverHttprouterToGin(GetPayerPlanPeriod))
	router.PUT("/payerplanperiods/:id", ConverHttprouterToGin(UpdatePayerPlanPeriod))
	router.DELETE("/payerplanperiods/:id", ConverHttprouterToGin(DeletePayerPlanPeriod))
}

func GetAllPayerPlanPeriods(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	payerplanperiods := []*model.PayerPlanPeriod{}

	payerplanperiods_orm := DB.Model(&model.PayerPlanPeriod{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		payerplanperiods_orm = payerplanperiods_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		payerplanperiods_orm = payerplanperiods_orm.Order(order)
	}

	if err = payerplanperiods_orm.Find(&payerplanperiods).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &payerplanperiods)
}

func GetPayerPlanPeriod(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	payerplanperiod := &model.PayerPlanPeriod{}
	if DB.First(payerplanperiod, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, payerplanperiod)
}

func AddPayerPlanPeriod(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payerplanperiod := &model.PayerPlanPeriod{}

	if err := readJSON(r, payerplanperiod); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(payerplanperiod).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, payerplanperiod)
}

func UpdatePayerPlanPeriod(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	payerplanperiod := &model.PayerPlanPeriod{}
	if DB.First(payerplanperiod, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.PayerPlanPeriod{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(payerplanperiod, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(payerplanperiod).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, payerplanperiod)
}

func DeletePayerPlanPeriod(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	payerplanperiod := &model.PayerPlanPeriod{}

	if DB.First(payerplanperiod, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(payerplanperiod).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
