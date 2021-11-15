package app

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/mod/exchange"
	"github.com/comhttp/jorm/mod/explorer"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type (
	OUR struct {
		Coins     []string
		Coin      string
		NodeCoins []string
		Explorers map[string]*explorer.Explorer
		WWW       *http.Server
		WS        *http.Server
		TLSconfig *tls.Config
		// comhttp    *COMHTTP
		goHTML     *template.Template
		config     cfg.Config
		jdbServers map[string]string
	}
	Index struct {
		Slug string      `json:"slug"`
		Data interface{} `json:"data"`
	}
)

func (o *OUR) JDBclient(jdbId string) (*jdb.JDB, error) {
	return jdb.NewJDB(o.jdbServers[jdbId])
}

func (o *OUR) ENSOhandlers() http.Handler {
	//coinsCollection := Queries(j.B["coins"],"coin")
	c, err := o.JDBclient("coins")
	utl.ErrorLog(err)
	cq := coin.Queries(c, "coin")

	e, err := o.JDBclient("exchanges")
	utl.ErrorLog(err)
	eq := exchange.Queries(e, "exchange")

	explorerJDBS := make(map[string]*jdb.JDB)

	for _, coin := range o.Explorers {
		jdbCl, err := o.JDBclient(coin.Coin)
		if err != nil {
			utl.ErrorLog(err)
		} else {
			explorerJDBS[coin.Coin] = jdbCl
		}

	}

	exq := explorer.Queries(explorerJDBS, "info")

	//exq := exchange.Queries(j.JDBclient("exchanges"), "exchange")
	//exq := exchange.Queries(j.JDBclient("explorers"),"explorer")
	r := mux.NewRouter()
	//s := r.Host("enso.okno.rs").Subrouter()
	r.StrictSlash(true)
	//n := r.PathPrefix("/n").Subrouter()
	coin.ENSOroutes(cq, r)
	exchange.ENSOroutes(eq, r)
	explorer.ENSOroutes(exq, r)
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}

func NewOUR(path string) (o *OUR) {
	o = new(OUR)
	// o.comhttp = &COMHTTP{}
	if path == "" {
		o.config.Path = "/var/db/jorm/"
	}
	c, _ := cfg.NewCFG(o.config.Path, nil)
	o.config = cfg.Config{}
	err := c.Read("conf", "conf", &o.config)
	utl.ErrorLog(err)

	o.jdbServers = make(map[string]string)
	err = c.Read("conf", "jdbs", &o.jdbServers)
	utl.ErrorLog(err)

	//ttt := j.JDBS.B["coins"].ReadAllPerPages("coin", 10, 1)

	o.WWW = &http.Server{
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("our")
	o.WWW.Handler = o.OURhandlers()
	o.WWW.Addr = ":" + o.config.Port["our"]
	return o
}

func (o *OUR) OURhandlers() http.Handler {
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OUR!"))
	})
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}

func (o *OUR) OurSRV() {
	fmt.Println("cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc: ")
	fmt.Println("Start OUR")
	fmt.Println("cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc: ")
	fmt.Println("cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc: ")
	fmt.Println("Start OUR process")

	s := strapi.New(o.config.Strapi)

	c, err := o.JDBclient("coins")
	utl.ErrorLog(err)
	fmt.Println("Start logos import")

	// logos := s.GetAll("logos")
	// lq := coin.Queries(c, "logo")

	// for _, logo := range logos {
	// 	// l := logo.([]map[string]interface{})[0].(map[string]interface{})
	// 	if logo != nil {
	// 		// l := logo.(map[string]interface{})
	// 		// l := ll[0].(map[string]interface{})
	// 		lq.WriteLogo(logo["slug"].(string), logo["data"])
	// 		time.Sleep(999 * time.Millisecond)
	// 	}
	// }

	// fmt.Println("End logos import")

	// }()

	// fmt.Println("logoslogoslogoslogoslogoslogoslogoslogoslogoslogos:", logos)

	cq := coin.Queries(c, "coin")

	cq.ProcessCoins(s)

	fmt.Println("End OUR process")

	log.Fatal().Err(o.WWW.ListenAndServe())
}
