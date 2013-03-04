// This plugin provides integration with MongoDB via the mgo package.
package revmgo

import (
	"fmt"
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

type Controller struct {
	*revel.Controller
	MSession *mgo.Session
}

var (
	// Global mgo Session
	MSession *mgo.Session
)

type MongoPlugin struct {
	revel.EmptyPlugin
}

func (p MongoPlugin) OnAppStart() {
	// Read configuration.
	//
	// TODO expand to include settings (currently just localhost)
	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]

	// Initialize our session
	var err error
	MSession, err = mgo.Dial("localhost")
	if err != nil {
		revel.ERROR.Fatal(err)
	}
}

// Attach a session to our controller.
func (p MongoPlugin) BeforeRequest(c *Controller) {
	// TODO make the option here configurable. New(), Clone(), Copy()
	c.MSession = MSession.Copy()
}

// Close the controller session if we have an active one.
func (p MongoPlugin) AfterRequest(c *Controller) {
	if c.MSession != nil {
		c.MSession.Close()
	}
	c.MSession = nil
}

// Close the controller session if we have an active one.
func (p MongoPlugin) OnException(c *Controller, err interface{}) {
	if c.MSession != nil {
		c.MSession.Close()
	}
	c.MSession = nil
}
