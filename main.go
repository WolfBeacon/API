package main

import (
	"net/http"

	"github.com/bugsnag/bugsnag-go"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"

	"github.com/cyrusroshan/API/db"
	"github.com/cyrusroshan/API/routes"
	"github.com/cyrusroshan/API/utils"
)

func bugsnagHandler() martini.Handler {
	return func(c martini.Context) {
		defer bugsnag.AutoNotify()

		c.Next()
	}
}

func main() {
	keyHolder := utils.KeyStore()

	m := martini.New()
	router := martini.NewRouter()

	router.NotFound(func() (int, []byte) {
		return 404, []byte("Requested page not found.")
	})

	m.Use(martini.Logger())
	m.Use(martini.Recovery())
	// m.Use(bugsnagHandler())
	m.Use(gzip.All())

	hackathonDB := db.InitHackathons(keyHolder("DATABASE_URL"))
	defer hackathonDB.Db.Close()
	// hackathonDB.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	m.Map(hackathonDB)

	m.Use(func(c martini.Context, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	router.Group("/cms", func(r martini.Router) {
		r.Get("/get/:id", cms.GetHackathon)
		r.Get("/delete/:id", cms.DeleteHackathon)
		r.Post("/new", cms.NewHackathon)
		r.Post("/edit/:id", cms.EditHackathon)
	})

	m.MapTo(router, (*martini.Routes)(nil))
	m.Action(router.Handle)

	m.Run()
}
