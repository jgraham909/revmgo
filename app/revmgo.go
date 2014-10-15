package revmgo

import (
	"fmt"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

var (
	Config    *MongoConfig
	Session   *mgo.Session
	Database  *mgo.Database
	Duplicate func() *mgo.Session
)

type MongoConfig struct {
	Host   string
	Method string
	Db     string
}

func Init() {
	h := revel.Config.StringDefault("revmgo.host", "localhost")
	m := revel.Config.StringDefault("revmgo.method", "clone")
	d := revel.Config.StringDefault("revmgo.database", "test")

	Config = &MongoConfig{h, m, d}

	if err := Dial(); err != nil {
		revel.ERROR.Fatal(err)
	}

	revel.TypeBinders[reflect.TypeOf(bson.NewObjectId())] = revel.Binder{
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

func (m *MongoController) Begin() revel.Result {

	if Session == nil {
		if err := Dial(); err != nil {
			return m.RenderError(err)
		}
	}

	m.MongoSession = Duplicate()
	m.Database = m.MongoSession.DB(Config.Db)

	return nil

}

func (m *MongoController) End() revel.Result {

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
