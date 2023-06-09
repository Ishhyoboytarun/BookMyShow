package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"BookMyShow/dtos"
	"BookMyShow/exceptions"
	"BookMyShow/models"
	"BookMyShow/services"
	"github.com/gorilla/mux"
)

type City struct {
	service services.CityServices
}

func NewCityController(service services.CityServices) *City {
	return &City{
		service: service,
	}
}

func (c *City) GetCities(w http.ResponseWriter, req *http.Request) {
	cities, err := c.service.GetCities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(cities)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *City) AddCity(w http.ResponseWriter, req *http.Request) {
	cities := &dtos.AddCityRequest{}
	err := json.NewDecoder(req.Body).Decode(cities)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = c.service.AddCities(cities)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *City) GetCity(w http.ResponseWriter, req *http.Request) {
	cityId, err := strconv.Atoi(mux.Vars(req)["city_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidCityId))
		return
	}

	city, err := c.service.GetCity(cityId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(city)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *City) UpdateCity(w http.ResponseWriter, req *http.Request) {
	cityId, err := strconv.Atoi(mux.Vars(req)["city_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidCityId))
		return
	}

	updatedValues := &models.City{}
	err = json.NewDecoder(req.Body).Decode(updatedValues)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = c.service.UpdateCity(cityId, updatedValues)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *City) DeleteCity(w http.ResponseWriter, req *http.Request) {
	cityId, err := strconv.Atoi(mux.Vars(req)["city_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidCityId))
		return
	}

	err = c.service.DeleteCity(cityId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}
