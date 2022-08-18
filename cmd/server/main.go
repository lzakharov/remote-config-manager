package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/lzakharov/remote-config-manager/pkg/provider/etcd"
	"github.com/lzakharov/remote-config-manager/pkg/transport"
)

func main() {
	addr := flag.String("addr", ":8080", `Server address.`)
	endpoints := flag.String("endpoints", "localhost:2379", `Comma separated list of Etcd endpoints`)
	flag.Parse()

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(*endpoints, ","),
		DialTimeout: time.Minute,
	})
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	srv := transport.NewServer(etcd.NewProvider(client))

	router := chi.NewRouter()
	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedHeaders: []string{"*"},
		}),
		middleware.Logger,
		middleware.Recoverer,
	)

	router.Route("/api", func(r chi.Router) {
		r.Get("/list", srv.ListKeys)
		r.Get("/get", srv.Get)
		r.Post("/put", srv.Put)
	})

	err = http.ListenAndServe(*addr, router)
	if err != nil {
		log.Panic(err)
	}
}
