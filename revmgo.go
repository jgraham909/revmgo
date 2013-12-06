package revmgo

import (
	"fmt"
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"reflect"
)

var (
	Session *mgo.Session // Global mgo Session
	Dial    string       // http://godoc.org/labix.org/v2/mgo#Dial
	Method  string       // clone, copy, new http://godoc.org/labix.org/v2/mgo#Session.New
	Mode    string       // strong, monotonic, eventual http://godoc.org/labix.org/v2/mgo#Session.SetMode
	// Holds the function to call for a given Method
	mgoSessionDupl func() *mgo.Session
)

// Optimization: Stores the method to call in mgoSessionDupl so that it only
// has to be looked up once (or whenever Session changes)
func setDuplMethod() {
	// Save which function to call for each request:
	switch Method {
	case "clone":
		mgoSessionDupl = Session.Clone
	case "copy":
		mgoSessionDupl = Session.Copy
	case "new":
		mgoSessionDupl = Session.New
	default:
		mgoSessionDupl = Session.Clone
	}
}

func setMode() {
	switch Mode {
	case "strong":
		Session.SetMode(mgo.Strong, true)
	case "monotonic":
		Session.SetMode(mgo.Monotonic, true)
	case "eventual":
		Session.SetMode(mgo.Eventual, true)
	default:
		Session.SetMode(mgo.Monotonic, true)
	}
}

func AppInit() {
	var err error
	// Read configuration.
	Dial = revel.Config.StringDefault("revmgo.dial", "localhost")
	Method = revel.Config.StringDefault("revmgo.method", "clone")
	Mode = revel.Config.StringDefault("revmgo.mode", "monotonic")
	if err = MethodError(Method); err != nil {
		revel.ERROR.Panic(err)
	}

	if err = ModeError(Mode); err != nil {
		revel.ERROR.Panic(err)
	}

	// Let's try to connect to Mongo DB right upon starting revel but don't
	// raise an error. Errors will be handled if there is actually a request
	if Session == nil {
		Session, err = mgo.Dial(Dial)
		if err != nil {
			// Only warn since we'll retry later for each request
			revel.WARN.Printf("Could not connect to Mongo DB. Error: %s", err)
		} else {
			setDuplMethod()
			setMode()
			Session.SetSafe(&mgo.Safe{})
		}
	}

	// register the custom bson.ObjectId binder
	objId := bson.NewObjectId()
	revel.TypeBinders[reflect.TypeOf(objId)] = ObjectIdBinder
}

func ControllerInit() {
	revel.InterceptMethod((*MongoController).Begin, revel.BEFORE)
	revel.InterceptMethod((*MongoController).End, revel.FINALLY)
}

type MongoController struct {
	*revel.Controller
	MongoSession *mgo.Session // named MongoSession to avoid collision with revel.Session
}

// Connect to mgo if we haven't already and return a copy/new/clone of the session
func (c *MongoController) Begin() revel.Result {
	// We may not be connected yet if revel was started before Mongo DB or
	// Mongo DB was restarted
	if Session == nil {
		var err error
		Session, err = mgo.Dial(Dial)
		if err != nil {
			// Extend the error description to include that this is a Mongo Error
			err = fmt.Errorf("Could not connect to Mongo DB. Error: %s", err)
			return c.RenderError(err)
		} else {
			setDuplMethod()
			setMode()
			Session.SetSafe(&mgo.Safe{})
		}
	}
	// Calls Clone(), Copy() or New() depending on the configuration
	c.MongoSession = mgoSessionDupl()
	return nil
}

// Close the controller session if we have an active one.
func (c *MongoController) End() revel.Result {
	// This is necessary since End() will be called no matter what
	// (revel.FINALLY) so it may not be connected in which case MongoSession
	// were a nil pointer and panic
	if c.MongoSession != nil {
		c.MongoSession.Close()
	}
	return nil
}

func MethodError(m string) error {
	switch m {
	case "clone", "copy", "new":
		return nil
	}
	return fmt.Errorf("revmgo: Invalid session instantiation method '%s'", m)
}

func ModeError(m string) error {
	switch m {
	case "strong", "monotonic", "eventual":
		return nil
	}
	return fmt.Errorf("revmgo: Invalid session mode '%s'", m)
}

// Custom TypeBinder for bson.ObjectId
// Makes additional Id parameters in actions obsolete
var ObjectIdBinder = revel.Binder{
	// Make a ObjectId from a request containing it in string format.
	Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
		if len(val) == 0 {
			return reflect.Zero(typ)
		}
		if bson.IsObjectIdHex(val) {
			objId := bson.ObjectIdHex(val)
			return reflect.ValueOf(objId)
		} else {
			revel.ERROR.Print("ObjectIdBinder.Bind - invalid ObjectId!")
			return reflect.Zero(typ)
		}
	}),
	// Turns ObjectId back to hexString for reverse routing
	Unbind: func(output map[string]string, name string, val interface{}) {
		var hexStr string
		hexStr = fmt.Sprintf("%s", val.(bson.ObjectId).Hex())
		// not sure if this is too carefull but i wouldn't want invalid ObjectIds in my App
		if bson.IsObjectIdHex(hexStr) {
			output[name] = hexStr
		} else {
			revel.ERROR.Print("ObjectIdBinder.Unbind - invalid ObjectId!")
			output[name] = ""
		}
	},
}
