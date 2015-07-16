revmgo [![Build Status](https://travis-ci.org/janekolszak/revmgo.svg?branch=master)](https://travis-ci.org/janekolszak/revmgo)
======
mgo module for revel framework

This is a **mantained** fork of https://github.com/jgraham909/revmgo.

### Installation
``` bash
    go get github.com/janekolszak/revmgo
```
### Test
``` bash
    revel test github.com/janekolszak/revmgo/testapp dev
```
### Configuration
In app.conf:
- **revmgo.dial** = [mgo.Session.Dial()](http://godoc.org/labix.org/v2/mgo#Dial)
- **revmgo.method** = One of 'clone', 'copy', 'new'. See [mgo.Session.New()](http://godoc.org/labix.org/v2/mgo#Session.New)

### Initialization
- In app.init() in `app/init.go` add:
``` go
    revel.OnAppStart(revmgo.AppInit)
```

- In controllers.init() in `app/controllers/init.go` add
``` go
    revmgo.ControllerInit()
```
So a minimal controller's init() would be:

``` go
    package controllers

    import "github.com/janekolszak/revmgo"

    func init() {
        revmgo.ControllerInit()
    }
```

### Usage
Embed the MongoController on your custom controller;
``` go
    package controllers

    import (
        "github.com/janekolszak/revmgo"
        "github.com/revel/revel"
    )

    type App struct {
        *revel.Controller
        revmgo.MongoController
  		// ...
  	}
```
The controller will now have a MongoSession variable of type `*mgo.Session`. Use this to query your mongo datastore.

### See Also

*  http://labix.org/v2/mgo for documentation on the mgo driver
*  https://github.com/jgraham909/bloggo for a reference implementation (Still a work in progress)


