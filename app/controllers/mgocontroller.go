// This plugin provides integration with MongoDB via the mgo package.
package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

type MgoController struct {
	*revel.Controller
	MSession *mgo.Session
}

var MGOSession *mgo.Session // Global mgo Session

// Connect to mgo if we haven't already and return a copy/new/clone of the session
func (c *MgoController) Begin() revel.Result {
	var err error
	if MGOSession == nil {
		// Read configuration.
		//
		// TODO expand to include settings (currently just localhost)
		// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
		MGOSession, err = mgo.Dial("localhost")
		if err != nil {
			revel.ERROR.Panic(err)
		}
	}
	// TODO make the option here configurable. New(), Clone(), Copy()
	c.MSession = MGOSession.Clone()
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
