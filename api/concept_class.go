package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configConceptClassesRouter(router *httprouter.Router) {
	router.GET("/conceptclasses", GetAllConceptClasses)
	router.POST("/conceptclasses", AddConceptClass)
	router.GET("/conceptclasses/:id", GetConceptClass)
	router.PUT("/conceptclasses/:id", UpdateConceptClass)
	router.DELETE("/conceptclasses/:id", DeleteConceptClass)
}

func configGinConceptClassesRouter(router gin.IRoutes) {
	router.GET("/conceptclasses", ConverHttprouterToGin(GetAllConceptClasses))
	router.POST("/conceptclasses", ConverHttprouterToGin(AddConceptClass))
	router.GET("/conceptclasses/:id", ConverHttprouterToGin(GetConceptClass))
	router.PUT("/conceptclasses/:id", ConverHttprouterToGin(UpdateConceptClass))
	router.DELETE("/conceptclasses/:id", ConverHttprouterToGin(DeleteConceptClass))
}

func GetAllConceptClasses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	conceptclasses := []*model.ConceptClass{}

	conceptclasses_orm := DB.Model(&model.ConceptClass{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		conceptclasses_orm = conceptclasses_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		conceptclasses_orm = conceptclasses_orm.Order(order)
	}

	if err = conceptclasses_orm.Find(&conceptclasses).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &conceptclasses)
}

func GetConceptClass(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conceptclass := &model.ConceptClass{}
	if DB.First(conceptclass, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, conceptclass)
}

func AddConceptClass(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	conceptclass := &model.ConceptClass{}

	if err := readJSON(r, conceptclass); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(conceptclass).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conceptclass)
}

func UpdateConceptClass(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	conceptclass := &model.ConceptClass{}
	if DB.First(conceptclass, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.ConceptClass{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(conceptclass, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(conceptclass).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conceptclass)
}

func DeleteConceptClass(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conceptclass := &model.ConceptClass{}

	if DB.First(conceptclass, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(conceptclass).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
