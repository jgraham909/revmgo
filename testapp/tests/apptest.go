package tests

import (
	"github.com/jgraham909/revmgo"
	"github.com/jgraham909/revmgo/testapp/app/models"
	"github.com/robfig/revel"
)

type AppTest struct {
	revel.TestSuite
}

func (t *AppTest) Before() {
	// Make sure our collection is clean
	models.Collection(revmgo.Session).DropCollection()
}

func (t AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")

}

func (t AppTest) TestSave() {
	b := models.GetBook("MobyDick")
	t.AssertEqual("Moby Dick", b.Title)
	b.Save(revmgo.Session)
	d := models.GetBookByObjectId(revmgo.Session, b.Id)
	t.AssertEqual(b.Title, d.Title)
	t.AssertEqual(b.Id, d.Id)
	t.AssertEqual(b.Body, d.Body)
	t.AssertEqual(b.Tags, d.Tags)
}

func (t *AppTest) After() {
	// Cleanup any mess we made
	models.Collection(revmgo.Session).DropCollection()
}
