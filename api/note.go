package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configNotesRouter(router *httprouter.Router) {
	router.GET("/notes", GetAllNotes)
	router.POST("/notes", AddNote)
	router.GET("/notes/:id", GetNote)
	router.PUT("/notes/:id", UpdateNote)
	router.DELETE("/notes/:id", DeleteNote)
}

func configGinNotesRouter(router gin.IRoutes) {
	router.GET("/notes", ConverHttprouterToGin(GetAllNotes))
	router.POST("/notes", ConverHttprouterToGin(AddNote))
	router.GET("/notes/:id", ConverHttprouterToGin(GetNote))
	router.PUT("/notes/:id", ConverHttprouterToGin(UpdateNote))
	router.DELETE("/notes/:id", ConverHttprouterToGin(DeleteNote))
}

func GetAllNotes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	notes := []*model.Note{}

	notes_orm := DB.Model(&model.Note{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		notes_orm = notes_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		notes_orm = notes_orm.Order(order)
	}

	if err = notes_orm.Find(&notes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &notes)
}

func GetNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	note := &model.Note{}
	if DB.First(note, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, note)
}

func AddNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	note := &model.Note{}

	if err := readJSON(r, note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(note).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, note)
}

func UpdateNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	note := &model.Note{}
	if DB.First(note, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Note{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(note, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(note).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, note)
}

func DeleteNote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	note := &model.Note{}

	if DB.First(note, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(note).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
