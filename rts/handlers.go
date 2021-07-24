package rts

import (
	"fmt"
	"github.com/comhttp/our/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
)

var (
	funcMap = template.FuncMap{
		"truncate": utl.Truncate,
	}
)

type Data struct {
	Our, Base, TLD, Slug, Coin, Bg, App, Section, URL, Page, ID, Title, Path, ProtoURL, Canonical string
}

func newData(our string) *Data {
	return &Data{
		Our: our,
	}
}

func (d *Data) base(base string) {
	d.Base = base
	return
}
func (d *Data) tld(tld string) {
	d.TLD = tld
	return
}
func (d *Data) slug(slug string) {
	d.Slug = slug
	return
}
func (d *Data) app(app string) {
	d.App = app
	return
}
func (d *Data) section(section string) {
	d.Section = section
	return
}

func Handlers() http.Handler {
	r := mux.NewRouter()
	d := newData("com-http")
	tld := r.Host("com-http.{tld}").Subrouter()
	tld.HandleFunc("/", d.appHandler())
	tld.HandleFunc("/{app}", d.appHandler())
	tld.HandleFunc("/{app}/{page}", d.appHandler())
	tld.Headers("Access-Control-Allow-Origin", "*")
	tld.StrictSlash(true)
	sub := r.Host("{slug}.com-http.{tld}").Subrouter()
	sub.HandleFunc("/", d.appHandler())
	sub.HandleFunc("/{app}", d.appHandler())
	sub.HandleFunc("/{app}/{page}", d.appHandler())
	sub.StrictSlash(true)
	sub.Headers("Access-Control-Allow-Origin", "*")
	sub.Headers("Content-Type", "application/json")
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}

func (d *Data) appHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d.TLD = mux.Vars(r)["tld"]
		d.Slug = mux.Vars(r)["slug"]
		d.App = mux.Vars(r)["app"]
		d.Page = mux.Vars(r)["page"]

		fmt.Println(":::::::::::::::::.TLD:::::::::::::::::. ", d.TLD)
		fmt.Println(":::::::::::::::::.Slug:::::::::::::::::. ", d.Slug)
		fmt.Println(":::::::::::::::::.App:::::::::::::::::. ", d.App)
		fmt.Println(":::::::::::::::::.Page:::::::::::::::::. ", d.Page)

		d.Base = "amp"
		d.Bg = "parallelcoin"
		d.Path = "rts/tld/" + d.TLD
		d.ProtoURL = "https://" + d.Our + "."
		d.Title = "Beyond blockchain - " + d.TLD
		fmt.Println("10000000000000000000000000000")

		if d.Page != "" {
			d.ID = d.Page
			d.Page = d.App
			d.Title = d.Slug + "-" + d.App + "-Beyond blockchain-" + d.TLD + "-" + d.Page
		} else {
			d.Page = "index"
		}

		if d.App != "" {
			d.Path = "rts/coin/" + d.TLD
			d.Title = d.Slug + " - Beyond blockchain - " + d.TLD
			d.ProtoURL = "https://" + d.Slug + "." + d.Our + "."
			d.Canonical = d.ProtoURL + d.TLD

		} else {
			d.App = d.TLD
		}

		if d.Slug != "" {
			d.Bg = d.Slug
			d.Path = "rts/coin/" + d.TLD
			d.Title = d.Slug + " - Beyond blockchain - " + d.TLD
			d.ProtoURL = "https://" + d.Slug + "." + d.Our + "."
			d.Canonical = d.ProtoURL + d.TLD
		} else {
			d.Section = "coin"
			d.Path = "rts/tld/" + d.TLD
		}

		if d.TLD == "net" {
			d.Base = "vue"
			d.Path = "vue"
			d.Page = "index"
			d.Title = "Beyond blockchain - " + d.TLD
			if d.Slug != "" {

			} else {
				d.Section = "coin"
				fmt.Println("2222222222222222222222222indexindex22222")
			}
			fmt.Println("111111111111111111111111111")
		}

		fmt.Println("ffffffffffffffffffffff")

		funcMap := template.FuncMap{
			"truncate": utl.Truncate,
		}
		fmt.Println("Top level domain keyword 1:  ", d.TLD)
		fmt.Println("App 1: ", d.App)
		fmt.Println("Slug 1: ", d.Slug)
		fmt.Println("Page 1: ", d.Page)
		fmt.Println("Path 1: ", d.Path)
		fmt.Println("Base 1: ", d.Base)
		template.Must(parseFiles(d.Base, d.Path+"/"+d.Page+".gohtml")).Funcs(funcMap).ExecuteTemplate(w, d.Base, d)
		fmt.Println("d.Base", d.Base)
		fmt.Println("dadadad", d.Path+"/"+d.Page+".gohtml")
	}
}
