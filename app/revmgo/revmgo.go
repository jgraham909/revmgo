package revmgo

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

var MGOSession *mgo.Session // Global mgo Session

type MongoPlugin struct {
	revel.EmptyPlugin
}

func (p MongoPlugin) OnAppStart() {
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
}

func init() {
	revel.RegisterPlugin(MongoPlugin{})
}
