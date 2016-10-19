package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"

	"github.com/smartystreets/go-aws-auth"
	"github.com/zenazn/goji"
        "github.com/zenazn/goji/web"
)

func proxy(destination string, signVersion int) func(c web.C, w http.ResponseWriter, r *http.Request) {
	uri, err := url.Parse(destination)
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(uri)
	return func(c web.C, w http.ResponseWriter, r *http.Request) {
		r.Host = uri.Host
		switch signVersion {
			case 4:
				awsauth.Sign4(r)
		}
		proxy.ServeHTTP(w, r)
	}
}

func ping(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

var (
	dest = flag.String("dest", "", "Destination Host (eg. API Gateway, ESS Host)")
	sign = flag.Int("sign", 4, "Sign Version")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	destination := *dest
	signVersion := *sign
	goji.Handle("/ping", ping)
	goji.Handle("/*", proxy(destination, signVersion))
	goji.Serve()
}
