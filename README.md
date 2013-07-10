revmgo
======

mgo module for revel framework

## Module Use

### app.conf

Add the following line to your revel application's init.go inside the init() function.

    revel.OnAppStart(func() { revmgo.AppInit() })

Additional settings can be configured via the following directives

#### revmgo.dial

Please review the documentation at [mgo.Session.Dial()](http://godoc.org/labix.org/v2/mgo#Dial) for information on the syntax and valid settings.

#### revmgo.method

This can be one of 'clone', 'copy', 'new'. See [mgo.Session.New()](http://godoc.org/labix.org/v2/mgo#Session.New) for more information.


### init.go

Then in your init.go include the following

    revel.OnAppStart(revmgo.Init)

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
