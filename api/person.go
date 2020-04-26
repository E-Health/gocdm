package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configPeopleRouter(router *httprouter.Router) {
	router.GET("/people", GetAllPeople)
	router.POST("/people", AddPerson)
	router.GET("/people/:id", GetPerson)
	router.PUT("/people/:id", UpdatePerson)
	router.DELETE("/people/:id", DeletePerson)
}

func configGinPeopleRouter(router gin.IRoutes) {
	router.GET("/people", ConverHttprouterToGin(GetAllPeople))
	router.POST("/people", ConverHttprouterToGin(AddPerson))
	router.GET("/people/:id", ConverHttprouterToGin(GetPerson))
	router.PUT("/people/:id", ConverHttprouterToGin(UpdatePerson))
	router.DELETE("/people/:id", ConverHttprouterToGin(DeletePerson))
}

func GetAllPeople(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	people := []*model.Person{}

	people_orm := DB.Model(&model.Person{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		people_orm = people_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		people_orm = people_orm.Order(order)
	}

	if err = people_orm.Find(&people).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &people)
}

func GetPerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	person := &model.Person{}
	if DB.First(person, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, person)
}

func AddPerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := &model.Person{}

	if err := readJSON(r, person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(person).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, person)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	person := &model.Person{}
	if DB.First(person, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Person{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(person, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(person).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	person := &model.Person{}

	if DB.First(person, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(person).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
