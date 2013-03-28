// This plugin provides integration with MongoDB via the mgo package.
package controllers

import (
	"github.com/jgraham909/revmgo/app/revmgo"
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

type MgoController struct {
	*revel.Controller
	MSession *mgo.Session
}

// Connect to mgo if we haven't already and return a copy/new/clone of the session
func (c *MgoController) Begin() revel.Result {
	// TODO make the option here configurable. New(), Clone(), Copy()
	c.MSession = revmgo.MGOSession.Clone()
	return nil
}

// Close the controller session if we have an active one.
func (c *MgoController) End() revel.Result {
	c.MSession.Close()
	return nil
}

func init() {
	revel.InterceptMethod((*MgoController).Begin, revel.BEFORE)
	revel.InterceptMethod((*MgoController).End, revel.FINALLY)
}
