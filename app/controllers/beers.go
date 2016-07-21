package controllers

import (
	"github.com/revel/revel"
	"github.com/sohjsolwin/mercury/app/models"
)

type Beers struct {
	*revel.Controller
}

var beers = []models.Beer{
	models.Beer{1, "La Trappe Quadrupel Oak Aged", "Ale", "Bierbrouwerij De Konigshoeven"},
	models.Beer{2, "Cocoa Psycho", "Porter", "BrewDog"},
	models.Beer{3, "American Dream", "Lager", "Mikkeller"},
}

func (c Beers) List() revel.Result {
	return c.RenderJson(beers)
}

func (c Beers) Show(beerID int) revel.Result {
	var res models.Beer

	for _, beer := range beers {
		if beer.ID == beerID {
			res = beer
		}
	}

	if res.ID == 0 {
		return c.NotFound("Could not find beer")
	}

	return c.RenderJson(res)
}
