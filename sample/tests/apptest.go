package tests

import (
	"github.com/creativelikeadog/revel-mgo/app"
	"github.com/creativelikeadog/revel-mgo/sample/app/models"
	"github.com/revel/revel"
)

type AppTest struct {
	revel.TestSuite
}

func (t *AppTest) Before() {
	// Make sure our collection is clean
	models.Collection(mgo.Database).DropCollection()
}

func (t AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")

}

func (t AppTest) TestSave() {
	b := models.GetBook("MobyDick")
	t.AssertEqual("Moby Dick", b.Title)
	b.Save(mgo.Database)
	d := models.FindByObjectId(mgo.Database, b.Id)
	t.AssertEqual(b.Title, d.Title)
	t.AssertEqual(b.Id, d.Id)
	t.AssertEqual(b.Body, d.Body)
	t.AssertEqual(b.Tags, d.Tags)
}

func (t *AppTest) After() {
	// Cleanup any mess we made
	models.Collection(mgo.Database).DropCollection()
}
