package dblayer

import (
	"github.com/zinirun/go-music/backend/src/models"
)

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
}
