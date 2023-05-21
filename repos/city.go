package repos

import (
	"database/sql"

	"github.com/dunzoit/BookMyShow/dtos"
	"github.com/dunzoit/BookMyShow/models"
)

const (
	getCities  = `SELECT id, name, state FROM city`
	addCities  = `INSERT INTO city VALUES($1, $2, $3)`
	getCity    = `SELECT id, name, state FROM city where id = $1`
	updateCity = `UPDATE cities SET id = $1, name = $2, state = $3 WHERE id = $4;`
	deleteCity = `DELETE FROM cities WHERE id = $1;`
)

type City struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) *City {
	return &City{
		db: db,
	}
}

func (c *City) GetCities() (*dtos.GetCitiesResponse, error) {

	rows, err := c.db.Query(getCities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var rowsResponse *dtos.GetCitiesResponse
	for rows.Next() {
		var id int
		var name string
		var state string
		if err := rows.Scan(&id, &name, &state); err != nil {
			return nil, err
		}
		row := &models.City{
			Id:    id,
			Name:  name,
			State: state,
		}
		rowsResponse.Cities = append(rowsResponse.Cities, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rowsResponse, nil
}

func (c *City) AddCities(cities *dtos.AddCitiesRequest) error {

	stmt, err := c.db.Prepare(addCities)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with each set of values
	for _, city := range cities.Cities {
		_, err := stmt.Exec(city.Id, city.Name, city.State)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *City) GetCity(cityId int) (*dtos.GetCityResponse, error) {

	rows, err := c.db.Query(getCity, cityId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	rows.Next()
	var id int
	var name string
	var state string
	if err := rows.Scan(&id, &name, &state); err != nil {
		return nil, err
	}

	var rowResponse *dtos.GetCityResponse
	rowResponse.City = &models.City{
		Id:    id,
		Name:  name,
		State: state,
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rowResponse, nil
}

func (c *City) UpdateCity(cityId int, updatedValues *dtos.UpdateCityRequest) error {
	_, err := c.db.Exec(updateCity, updatedValues.Id, updatedValues.Name, updatedValues.State, cityId)
	if err != nil {
		return err
	}
	return nil
}

func (c *City) DeleteCity(cityId int) error {
	_, err := c.db.Exec(deleteCity, cityId)
	if err != nil {
		return err
	}
	return nil
}