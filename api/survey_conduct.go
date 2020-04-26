package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configSurveyConductsRouter(router *httprouter.Router) {
	router.GET("/surveyconducts", GetAllSurveyConducts)
	router.POST("/surveyconducts", AddSurveyConduct)
	router.GET("/surveyconducts/:id", GetSurveyConduct)
	router.PUT("/surveyconducts/:id", UpdateSurveyConduct)
	router.DELETE("/surveyconducts/:id", DeleteSurveyConduct)
}

func configGinSurveyConductsRouter(router gin.IRoutes) {
	router.GET("/surveyconducts", ConverHttprouterToGin(GetAllSurveyConducts))
	router.POST("/surveyconducts", ConverHttprouterToGin(AddSurveyConduct))
	router.GET("/surveyconducts/:id", ConverHttprouterToGin(GetSurveyConduct))
	router.PUT("/surveyconducts/:id", ConverHttprouterToGin(UpdateSurveyConduct))
	router.DELETE("/surveyconducts/:id", ConverHttprouterToGin(DeleteSurveyConduct))
}

func GetAllSurveyConducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	surveyconducts := []*model.SurveyConduct{}

	surveyconducts_orm := DB.Model(&model.SurveyConduct{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		surveyconducts_orm = surveyconducts_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		surveyconducts_orm = surveyconducts_orm.Order(order)
	}

	if err = surveyconducts_orm.Find(&surveyconducts).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &surveyconducts)
}

func GetSurveyConduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	surveyconduct := &model.SurveyConduct{}
	if DB.First(surveyconduct, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, surveyconduct)
}

func AddSurveyConduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	surveyconduct := &model.SurveyConduct{}

	if err := readJSON(r, surveyconduct); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(surveyconduct).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, surveyconduct)
}

func UpdateSurveyConduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	surveyconduct := &model.SurveyConduct{}
	if DB.First(surveyconduct, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.SurveyConduct{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(surveyconduct, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(surveyconduct).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, surveyconduct)
}

func DeleteSurveyConduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	surveyconduct := &model.SurveyConduct{}

	if DB.First(surveyconduct, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(surveyconduct).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
