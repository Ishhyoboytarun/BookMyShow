package services

import (
	"github.com/dunzoit/BookMyShow/dtos"
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/repos"
)

type Theatre struct {
	repo repos.TheatreRepository
}

func NewTheatreService(repo repos.TheatreRepository) *Theatre {
	return &Theatre{
		repo: repo,
	}
}

func (t *Theatre) GetTheatres(cityID int) ([]*models.Theatre, error) {
	return t.repo.GetTheatres(cityID)
}

func (t *Theatre) AddTheatre(theatre *models.Theatre) error {
	return t.repo.AddTheatre(theatre)
}

func (t *Theatre) GetTheatre(theatreID int) (*models.Theatre, error) {
	return t.repo.GetTheatre(theatreID)
}

func (t *Theatre) UpdateTheatre(theatreID int, updatedValues *dtos.TheatreRequest) error {
	return t.repo.UpdateTheatre(theatreID, updatedValues)
}

func (t *Theatre) DeleteTheatre(theatreID int) error {
	return t.repo.DeleteTheatre(theatreID)
}
