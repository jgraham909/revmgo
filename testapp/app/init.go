package app

import (
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(revmgo.AppInit)
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.OnAppStart(func() {

		keyValues := []map[string]string{}

		db := revmgo.Session.DB("test")
		col := db.C("config")
		query := col.Find(nil)

		err := query.All(&keyValues)

		if err != nil {
			revel.ERROR.Printf("%v", err)
		} else if keyValues != nil {
			for key, value := range keyValues {
				revel.INFO.Printf("%s: %s", key, value)
			}
		} else {
			revel.INFO.Printf("No keys found")
		}

	})
}
