package revmgo

import (
	"fmt"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

var (
	// Global config
	Config    *MongoConfig
	// Global mgo Session
	Session   *mgo.Session 
	// Global mgo Database
	Database  *mgo.Database
	// Optimization: Stores the method to call in mgoSessionDupl so that it only 
	// has to be looked up once (or whenever Session changes)
	Duplicate func() *mgo.Session 
)

type MongoConfig struct {
	Host   string
	Method string
	Db     string
}

func Init() {
	// Read configuration.
	h := revel.Config.StringDefault("revmgo.host", "localhost")
	m := revel.Config.StringDefault("revmgo.method", "clone")
	d := revel.Config.StringDefault("revmgo.database", "test")

	Config = &MongoConfig{h, m, d}

	// Let's try to connect to Mongo DB right upon starting revel but don't
-	// raise an error. Errors will be handled if there is actually a request
	if err := Dial(); err != nil {
		revel.WARN.Printf("Could not connect to Mongo DB. Error: %s", err)
	}

	// register the custom bson.ObjectId binder
	revel.TypeBinders[reflect.TypeOf(bson.NewObjectId())] = revel.Binder{
		// Make a ObjectId from a request containing it in string format.
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			if len(val) == 0 {
				return reflect.Zero(typ)
			}
			if bson.IsObjectIdHex(val) {
				objId := bson.ObjectIdHex(val)
				return reflect.ValueOf(objId)
			} else {
				revel.ERROR.Print("Invalid ObjectId")
				return reflect.Zero(typ)
			}
		}),
		// Turns ObjectId back to hexString for reverse routing
		Unbind: func(output map[string]string, name string, val interface{}) {

			hexStr := fmt.Sprintf("%s", val.(bson.ObjectId).Hex())

			if bson.IsObjectIdHex(hexStr) {
				output[name] = hexStr
			} else {
				revel.ERROR.Print("Invalid ObjectId")
				output[name] = ""
			}
		},
	}
}

// Main Dial func 
func Dial() error {

	var (
		m   func() *mgo.Session
		err error
	)

	Session, err = mgo.Dial(Config.Host)

	if err != nil {
		return err
	}

	revel.INFO.Println("Mongo session started")

	switch Config.Method {
	case "clone":
		m = Session.Clone
	case "copy":
		m = Session.Copy
	case "new":
		m = Session.New
	}

	if m == nil {
		revel.WARN.Printf("Method %s is not allowed.", Config.Method)
		Config.Method = "clone"
		m = Session.Clone
	}
	
	Duplicate = m

	Database = Session.DB(Config.Db)

	return nil
}

type MongoController struct {
	*revel.Controller
	MongoSession *mgo.Session
	Database     *mgo.Database
}

// Connect to mgo if we haven't already and return a copy/new/clone of the session
func (m *MongoController) Begin() revel.Result {
	// We may not be connected yet if revel was started before Mongo DB or
	// Mongo DB was restarted
	if Session == nil {
		if err := Dial(); err != nil {
			// Extend the error description to include that this is a Mongo Error
			err = fmt.Errorf("Could not connect to Mongo DB. Error: %s", err)
			return m.RenderError(err)
		}
	}

	// Calls Clone(), Copy() or New() depending on the configuration
	m.MongoSession = Duplicate()
	m.Database = m.MongoSession.DB(Config.Db)

	return nil

}

// Close the controller session if we have an active one.
func (m *MongoController) End() revel.Result {
	// This is necessary since End() will be called no matter what
	// (revel.FINALLY) so it may not be connected in which case MongoSession
	// were a nil pointer and panic
	if m.MongoSession != nil {
		m.MongoSession.Close()
		m.Database = nil
	}

	return nil
}

func init() {
	revel.OnAppStart(Init)
	revel.InterceptMethod((*MongoController).Begin, revel.BEFORE)
	revel.InterceptMethod((*MongoController).End, revel.FINALLY)
}
