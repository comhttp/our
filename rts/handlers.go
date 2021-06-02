package rts

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/oknors/comhttpus/utl"
	"net/http"
	"text/template"
)

var (
	funcMap = template.FuncMap{
		"truncate": utl.Truncate,
	}
)

func parseFiles(tpl string) (*template.Template, error) {
	return template.ParseFiles("tpl/head.gohtml", "tpl/js.gohtml", "tpl/style.gohtml", "tpl/error.gohtml", "tpl/footer.gohtml", "tpl/index.gohtml")
}

func Handlers() http.Handler {
	r := mux.NewRouter()
	tld := r.Host("com-http.{tld}").Subrouter()
	tld.HandleFunc("/", indexHandler())

	tld.HandleFunc("/{section}", sectionHandler())
	tld.HandleFunc("/{section}/{page}", pageHandler())
	tld.Headers("Access-Control-Allow-Origin", "*")

	sub := r.Host("{sub}.com-http.{tld}").Subrouter()
	sub.HandleFunc("/", subHandler())
	sub.HandleFunc("/{section}", sectionHandler())
	sub.HandleFunc("/{section}/{page}", pageHandler())
	sub.Headers("Access-Control-Allow-Origin", "*")

	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))

}

func indexHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		data := map[string]interface{}{
			"TLD":   tld,
			"Title": "Beyond blockchain - " + tld,
		}
		template.Must(parseFiles("index")).Funcs(funcMap).ExecuteTemplate(w, "index", data)
	}
}

func subHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		sub := mux.Vars(r)["sub"]
		data := map[string]interface{}{
			"TLD":   tld,
			"Sub":   sub,
			"Title": "Beyond blockchain - " + tld + sub,
		}
		funcMap := template.FuncMap{
			"truncate": utl.Truncate,
		}

		template.Must(parseFiles(sub)).Funcs(funcMap).ExecuteTemplate(w, "index", data)
	}
}
func sectionHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		sub := mux.Vars(r)["sub"]
		section := mux.Vars(r)["section"]
		data := map[string]interface{}{
			"TLD":   tld,
			"Slug":  sub,
			"Title": "Beyond blockchain - " + tld + sub,
		}
		funcMap := template.FuncMap{
			"truncate": utl.Truncate,
		}
		//if tld == "us" {
		//	landing = "coin"
		//}
		template.Must(parseFiles("section/"+section+"/index")).Funcs(funcMap).ExecuteTemplate(w, "base", data)
	}
}

func pageHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		sub := mux.Vars(r)["sub"]
		section := mux.Vars(r)["section"]
		page := mux.Vars(r)["page"]
		data := map[string]interface{}{
			"TLD":     tld,
			"Section": section,
			"Page":    page,
			"Slug":    sub,
			"Title":   "Beyond blockchain - " + tld + sub,
		}
		funcMap := template.FuncMap{
			"truncate": utl.Truncate,
		}
		//if tld == "us" {
		//	landing = "coin"
		//}
		template.Must(parseFiles("section/"+section+"/"+page)).Funcs(funcMap).ExecuteTemplate(w, "base", data)
	}
}
