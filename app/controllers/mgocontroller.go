// This plugin provides integration with MongoDB via the mgo package.
package controllers

import (
	"github.com/jgraham909/revmgo/app/revmgo"
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

type MongoController struct {
	*revel.Controller
	MongoSession *mgo.Session
}

// Connect to mgo if we haven't already and return a copy/new/clone of the session
func (c *MongoController) Begin() revel.Result {
	// TODO make the option here configurable. New(), Clone(), Copy()
	c.MongoSession = revmgo.Session.Clone()
	return nil
}

// Close the controller session if we have an active one.
func (c *MongoController) End() revel.Result {
	c.MongoSession.Close()
	return nil
}

func init() {
	revel.InterceptMethod((*MongoController).Begin, revel.BEFORE)
	revel.InterceptMethod((*MongoController).End, revel.FINALLY)
}
