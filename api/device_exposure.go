package api

import (
	"net/http"

	"github.com/E-Health/gocdm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configDeviceExposuresRouter(router *httprouter.Router) {
	router.GET("/deviceexposures", GetAllDeviceExposures)
	router.POST("/deviceexposures", AddDeviceExposure)
	router.GET("/deviceexposures/:id", GetDeviceExposure)
	router.PUT("/deviceexposures/:id", UpdateDeviceExposure)
	router.DELETE("/deviceexposures/:id", DeleteDeviceExposure)
}

func configGinDeviceExposuresRouter(router gin.IRoutes) {
	router.GET("/deviceexposures", ConverHttprouterToGin(GetAllDeviceExposures))
	router.POST("/deviceexposures", ConverHttprouterToGin(AddDeviceExposure))
	router.GET("/deviceexposures/:id", ConverHttprouterToGin(GetDeviceExposure))
	router.PUT("/deviceexposures/:id", ConverHttprouterToGin(UpdateDeviceExposure))
	router.DELETE("/deviceexposures/:id", ConverHttprouterToGin(DeleteDeviceExposure))
}

func GetAllDeviceExposures(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order := r.FormValue("order")

	deviceexposures := []*model.DeviceExposure{}

	deviceexposures_orm := DB.Model(&model.DeviceExposure{})

	if page > 0 {
		pagesize, err := readInt(r, "pagesize", 20)
		if err != nil || pagesize <= 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pagesize

		deviceexposures_orm = deviceexposures_orm.Offset(offset).Limit(pagesize)
	}

	if order != "" {
		deviceexposures_orm = deviceexposures_orm.Order(order)
	}

	if err = deviceexposures_orm.Find(&deviceexposures).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, &deviceexposures)
}

func GetDeviceExposure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	deviceexposure := &model.DeviceExposure{}
	if DB.First(deviceexposure, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, deviceexposure)
}

func AddDeviceExposure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deviceexposure := &model.DeviceExposure{}

	if err := readJSON(r, deviceexposure); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(deviceexposure).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, deviceexposure)
}

func UpdateDeviceExposure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	deviceexposure := &model.DeviceExposure{}
	if DB.First(deviceexposure, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.DeviceExposure{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(deviceexposure, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(deviceexposure).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, deviceexposure)
}

func DeleteDeviceExposure(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	deviceexposure := &model.DeviceExposure{}

	if DB.First(deviceexposure, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(deviceexposure).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
