revmgo
======

mgo module for revel framework

## Module Use

### app.conf

Settings can be configured via the following directives in app.conf.

#### revmgo.dial

Please review the documentation at [mgo.Session.Dial()](http://godoc.org/gopkg.in/mgo.v2#Dial) for information on the syntax and valid settings.

#### revmgo.method

This can be one of 'clone', 'copy', 'new'. See [mgo.Session.New()](http://godoc.org/gopkg.in/mgo.v2#Session.New) for more information.


### app.init()

Add the following inside the app.init() function in `app/init.go`.

    revel.OnAppStart(revmgo.AppInit)

### controllers.init()

Similarly for your controllers init() function you must add the `revmgo.ControllerInit()` method. A minimum `app/controllers/init.go` file is represented below.

    package controllers

    import "github.com/jgraham909/revmgo"

    func init() {
        revmgo.ControllerInit()
    }


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

*  http://gopkg.in/mgo.v2 for documentation on the mgo driver
*  https://github.com/jgraham909/bloggo for a reference implementation (Still a work in progress)

[![Build Status](https://travis-ci.org/jgraham909/revmgo.png)](https://travis-ci.org/jgraham909/revmgo)
