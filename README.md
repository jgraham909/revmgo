revmgo
======

mgo module for revel framework

## Module Use

### app.conf

Add the following line to your app.conf.

    module.revmgo = github.com/jgraham909/revmgo

### Embedding the controller

In any controller you want to have mongo connectivity you must include the 
MgoController. If all routes need database then you should likely implement the
controller embedding as a base controller for your project and then embed that
your application specific controller.

Add the following import line in source files that will embed MgoController. Note that
we alias the import to 'm' since both the controller source file and the MgoController
have package 'controllers'.

    m "github.com/jgraham909/revmgo/app/controllers"

Embed the MgoController on your custom controller;

    type Application struct {
  		m.MgoController
  		// Other fields
  	}


Your controller will now have a MSession variable of type *mgo.Session. Use this
to query your mongo datastore. 

### See Also

*  http://labix.org/v2/mgo for documentation on the mgo driver
*  https://github.com/jgraham909/bloggo for a reference implementation (Still a work in progress)
