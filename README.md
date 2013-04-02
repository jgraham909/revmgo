revmgo
======

mgo module for revel framework

## Module Use

### app.conf

Add the following line to your app.conf.

    module.revmgo = github.com/jgraham909/revmgo

### Embedding the controller

In any controller you want to have mongo connectivity you must include the
MongoController. 

Add the following import line in source files that will embed MongoController. Note that
we alias the import to 'm' since both the controller source file and the MongoController
have package 'controllers'.

    m "github.com/jgraham909/revmgo/app/controllers"

Embed the MongoController on your custom controller;

    type Application struct {
  		*revel.Controller
        m.MongoController
  		// Other fields
  	}


Your controller will now have a MongoSession variable of type *mgo.Session. Use this
to query your mongo datastore.

### See Also

*  http://labix.org/v2/mgo for documentation on the mgo driver
*  https://github.com/jgraham909/bloggo for a reference implementation (Still a work in progress)
