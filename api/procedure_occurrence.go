package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configProcedureOccurrencesRouter(router *httprouter.Router) {
	router.GET("/procedureoccurrences", GetAllProcedureOccurrences)
	router.POST("/procedureoccurrences", AddProcedureOccurrence)
	router.GET("/procedureoccurrences/:id", GetProcedureOccurrence)
	router.PUT("/procedureoccurrences/:id", UpdateProcedureOccurrence)
	router.DELETE("/procedureoccurrences/:id", DeleteProcedureOccurrence)
}

func configGinProcedureOccurrencesRouter(router gin.IRoutes) {
	router.GET("/procedureoccurrences", ConverHttprouterToGin(GetAllProcedureOccurrences))
	router.POST("/procedureoccurrences", ConverHttprouterToGin(AddProcedureOccurrence))
	router.GET("/procedureoccurrences/:id", ConverHttprouterToGin(GetProcedureOccurrence))
	router.PUT("/procedureoccurrences/:id", ConverHttprouterToGin(UpdateProcedureOccurrence))
	router.DELETE("/procedureoccurrences/:id", ConverHttprouterToGin(DeleteProcedureOccurrence))
}

func GetAllProcedureOccurrences(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	procedureoccurrences := []*model.ProcedureOccurrence{}

	procedureoccurrences_orm := DB.Model(&model.ProcedureOccurrence{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		procedureoccurrences_orm = procedureoccurrences_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		procedureoccurrences_orm = procedureoccurrences_orm.Order(order)
	}

	if err = procedureoccurrences_orm.Find(&procedureoccurrences).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &procedureoccurrences)
}

func GetProcedureOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	procedureoccurrence := &model.ProcedureOccurrence{}
	if DB.First(procedureoccurrence, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, procedureoccurrence)
}

func AddProcedureOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	procedureoccurrence := &model.ProcedureOccurrence{}

	if err := readJSON(r, procedureoccurrence); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(procedureoccurrence).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, procedureoccurrence)
}

func UpdateProcedureOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	procedureoccurrence := &model.ProcedureOccurrence{}
	if DB.First(procedureoccurrence, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.ProcedureOccurrence{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(procedureoccurrence, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(procedureoccurrence).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, procedureoccurrence)
}

func DeleteProcedureOccurrence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	procedureoccurrence := &model.ProcedureOccurrence{}

	if DB.First(procedureoccurrence, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(procedureoccurrence).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
