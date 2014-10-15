package controllers

import (
	"github.com/creativelikeadog/revmgo/app"
	"github.com/creativelikeadog/revmgo/sample/app/models"
	"github.com/revel/revel"
)

type Book struct {
	*revel.Controller
	revmgo.MongoController
}

func (c Book) Index() revel.Result {
	return c.Render()
}

func (c Book) View(id string) revel.Result {
	b := models.FindById(c.Database, id)

	if b.Id.Hex() != id {
		return c.NotFound("Could not find a book with '%s' as id.", id)
	}

	return c.Render(b)
}
