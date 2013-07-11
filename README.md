revmgo
======

mgo module for revel framework

## Module Use

### app.conf

Settings can be configured via the following directives in app.conf.

#### revmgo.dial

Please review the documentation at [mgo.Session.Dial()](http://godoc.org/labix.org/v2/mgo#Dial) for information on the syntax and valid settings.

#### revmgo.method

This can be one of 'clone', 'copy', 'new'. See [mgo.Session.New()](http://godoc.org/labix.org/v2/mgo#Session.New) for more information.


### init.go

Add the following inside the init() function in your application's init.go.

    revel.OnAppStart(revmgo.AppInit)

### Embedding the controller

In any controller you want to have mongo connectivity you must include the
MongoController.

Add the following import line in source files that will embed MongoController.

     "github.com/jgraham909/revmgo"

Embed the MongoController on your custom controller;

    type Application struct {
  		*revel.Controller
      revmgo.MongoController
  		// Other fields
  	}


Your controller will now have a MongoSession variable of type *mgo.Session. Use this
to query your mongo datastore.

### See Also

*  http://labix.org/v2/mgo for documentation on the mgo driver
*  https://github.com/jgraham909/bloggo for a reference implementation (Still a work in progress)
