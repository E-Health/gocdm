package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configConceptAncestorsRouter(router *httprouter.Router) {
	router.GET("/conceptancestors", GetAllConceptAncestors)
	router.POST("/conceptancestors", AddConceptAncestor)
	router.GET("/conceptancestors/:id", GetConceptAncestor)
	router.PUT("/conceptancestors/:id", UpdateConceptAncestor)
	router.DELETE("/conceptancestors/:id", DeleteConceptAncestor)
}

func configGinConceptAncestorsRouter(router gin.IRoutes) {
	router.GET("/conceptancestors", ConverHttprouterToGin(GetAllConceptAncestors))
	router.POST("/conceptancestors", ConverHttprouterToGin(AddConceptAncestor))
	router.GET("/conceptancestors/:id", ConverHttprouterToGin(GetConceptAncestor))
	router.PUT("/conceptancestors/:id", ConverHttprouterToGin(UpdateConceptAncestor))
	router.DELETE("/conceptancestors/:id", ConverHttprouterToGin(DeleteConceptAncestor))
}

func GetAllConceptAncestors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	conceptancestors := []*model.ConceptAncestor{}

	conceptancestors_orm := DB.Model(&model.ConceptAncestor{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		conceptancestors_orm = conceptancestors_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		conceptancestors_orm = conceptancestors_orm.Order(order)
	}

	if err = conceptancestors_orm.Find(&conceptancestors).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &conceptancestors)
}

func GetConceptAncestor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conceptancestor := &model.ConceptAncestor{}
	if DB.First(conceptancestor, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, conceptancestor)
}

func AddConceptAncestor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	conceptancestor := &model.ConceptAncestor{}

	if err := readJSON(r, conceptancestor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(conceptancestor).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conceptancestor)
}

func UpdateConceptAncestor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	conceptancestor := &model.ConceptAncestor{}
	if DB.First(conceptancestor, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.ConceptAncestor{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(conceptancestor, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(conceptancestor).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, conceptancestor)
}

func DeleteConceptAncestor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	conceptancestor := &model.ConceptAncestor{}

	if DB.First(conceptancestor, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(conceptancestor).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
