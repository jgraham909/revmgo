package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

type RevmgoController struct {
	*revel.Controller
	MSession *mgo.Session
}
