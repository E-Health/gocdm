package api

import (
	"net/http"

	"model/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configNoteNlpsRouter(router *httprouter.Router) {
	router.GET("/notenlps", GetAllNoteNlps)
	router.POST("/notenlps", AddNoteNlp)
	router.GET("/notenlps/:id", GetNoteNlp)
	router.PUT("/notenlps/:id", UpdateNoteNlp)
	router.DELETE("/notenlps/:id", DeleteNoteNlp)
}

func configGinNoteNlpsRouter(router gin.IRoutes) {
	router.GET("/notenlps", ConverHttprouterToGin(GetAllNoteNlps))
	router.POST("/notenlps", ConverHttprouterToGin(AddNoteNlp))
	router.GET("/notenlps/:id", ConverHttprouterToGin(GetNoteNlp))
	router.PUT("/notenlps/:id", ConverHttprouterToGin(UpdateNoteNlp))
	router.DELETE("/notenlps/:id", ConverHttprouterToGin(DeleteNoteNlp))
}

func GetAllNoteNlps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	notenlps := []*model.NoteNlp{}

	notenlps_orm := DB.Model(&model.NoteNlp{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		notenlps_orm = notenlps_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		notenlps_orm = notenlps_orm.Order(order)
	}

	if err = notenlps_orm.Find(&notenlps).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &notenlps)
}

func GetNoteNlp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	notenlp := &model.NoteNlp{}
	if DB.First(notenlp, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, notenlp)
}

func AddNoteNlp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	notenlp := &model.NoteNlp{}

	if err := readJSON(r, notenlp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(notenlp).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, notenlp)
}

func UpdateNoteNlp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	notenlp := &model.NoteNlp{}
	if DB.First(notenlp, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.NoteNlp{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(notenlp, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(notenlp).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, notenlp)
}

func DeleteNoteNlp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	notenlp := &model.NoteNlp{}

	if DB.First(notenlp, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(notenlp).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
