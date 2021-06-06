package rts

import (
	"fmt"
	"github.com/comhttp/our/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"text/template"
)

var (
	funcMap = template.FuncMap{
		"truncate": utl.Truncate,
	}
	our = "com-http"
)

func parseFiles(tpl string) (*template.Template, error) {
	return template.ParseFiles(tpl, "tpl/amp/lib/boilerplate.gohtml", "tpl/amp/lib/css.gohtml", "tpl/amp/lib/font.gohtml", "tpl/amp/lib/colors.gohtml", "tpl/amp/base.gohtml", "tpl/amp/head.gohtml", "tpl/amp/main.gohtml", "tpl/amp/el/header.gohtml", "tpl/amp/el/footer.gohtml", "tpl/amp/el/nodecoins.gohtml", "tpl/amp/el/algocoins.gohtml", "tpl/amp/el/restcoins.gohtml", "tpl/amp/js.gohtml", "tpl/amp/style.gohtml", "tpl/el/loader.gohtml")
}

func Handlers() http.Handler {
	r := mux.NewRouter()
	tld := r.Host("com-http.{tld}").Subrouter()
	tld.HandleFunc("/", indexHandler())
	tld.Headers("Access-Control-Allow-Origin", "*")

	sub := r.Host("{slug}.com-http.{tld}").Subrouter()
	sub.HandleFunc("/", indexHandler())
	sub.HandleFunc("/{app}", appHandler())
	sub.HandleFunc("/{app}/{section}", sectionHandler())
	sub.HandleFunc("/{app}/{section}/{item}", itemHandler())
	sub.Headers("Access-Control-Allow-Origin", "*")
	sub.Headers("Content-Type", "application/json")

	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))

}

func indexHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		slug := mux.Vars(r)["slug"]
		app := "index"
		if slug != "" {
			app = "coin"
		}
		data := map[string]interface{}{
			"TLD":  tld,
			"Slug": slug,
			//"Type": t,
			"App":     app,
			"Section": "index",
			"Title":   "Beyond blockchain - " + tld + " - ",
		}
		fmt.Println("Top level domain keyword 1:  ", tld)
		fmt.Println("App 1: ", app)
		fmt.Println("Slug 1: ", slug)
		template.Must(parseFiles("tpl/amp/rts/"+tld+"/index.gohtml")).Funcs(funcMap).ExecuteTemplate(w, "base", data)
	}
}

//func subIndexHandler() func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		tld := mux.Vars(r)["tld"]
//		sub := mux.Vars(r)["sub"]
//		title := strings.Title(sub  + " - Beyond blockchain" + " - " + tld)
//		data := map[string]interface{}{
//			"TLD":   tld,
//			"Sub":   "coin",
//			"App":  "index",
//			"Title": title,
//		}
//		fmt.Println("Top level domain keyword: ", tld)
//		fmt.Println("Subdomain keyword: ", sub)
//		funcMap := template.FuncMap{
//			"truncate": utl.Truncate,
//		}
//		template.Must(parseFiles("tpl/amp/sub.gohtml")).Funcs(funcMap).ExecuteTemplate(w, "base", data)
//	}
//}
func appHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		slug := mux.Vars(r)["slug"]
		app := mux.Vars(r)["app"]
		//if app == ""{
		//	app = "home"
		//}
		//if app == "" && slug != ""{
		//	app = "coin"
		//}

		title := strings.Title(slug + " - Beyond blockchain - " + tld + " - " + app)
		data := map[string]interface{}{
			"TLD":     tld,
			"Slug":    slug,
			"App":     app,
			"Section": "home",
			"URL":     "home",
			"Title":   title,
		}
		funcMap := template.FuncMap{
			"truncate": utl.Truncate,
		}
		//if tld == "us" {
		//	landing = "coin"
		//}
		template.Must(parseFiles("tpl/amp/index.gohtml")).Funcs(funcMap).ExecuteTemplate(w, "base", data)
	}
}
func sectionHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		slug := mux.Vars(r)["slug"]
		app := mux.Vars(r)["app"]
		section := mux.Vars(r)["section"]
		data := map[string]interface{}{
			"TLD":     tld,
			"Slug":    slug,
			"App":     app,
			"Section": section,
			"Item":    "section",
			"Title":   slug + " - " + tld + " - " + section + " - Beyond blockchain",
		}
		funcMap := template.FuncMap{
			"truncate": utl.Truncate,
		}
		//if tld == "us" {
		//	landing = "coin"
		//}
		//template.Must(parseFiles("tpl/amp/" + section+"/index")).Funcs(funcMap).ExecuteTemplate(w, "base", data)
		template.Must(parseFiles("tpl/amp/index.gohtml")).Funcs(funcMap).ExecuteTemplate(w, "base", data)
	}
}

func itemHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		slug := mux.Vars(r)["slug"]
		app := mux.Vars(r)["app"]
		section := mux.Vars(r)["section"]
		item := mux.Vars(r)["item"]
		data := map[string]interface{}{
			"TLD":     tld,
			"Slug":    slug,
			"App":     app,
			"Section": section,
			"Item":    item,
			"Title":   slug + " - " + tld + " - " + section + " - " + item + " - Beyond blockchain",
		}
		funcMap := template.FuncMap{
			"truncate": utl.Truncate,
		}
		//if tld == "us" {
		//	landing = "coin"
		//}
		//template.Must(parseFiles("tpl/amp/"+app+"/"+section+"/"+item)).Funcs(funcMap).ExecuteTemplate(w, "base", data)
		template.Must(parseFiles("tpl/amp/index.gohtml")).Funcs(funcMap).ExecuteTemplate(w, "base", data)

	}
}
