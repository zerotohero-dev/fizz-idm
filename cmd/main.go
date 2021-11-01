/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spiffe/go-spiffe/v2/spiffetls"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-idm/internal/api"
	"github.com/zerotohero-dev/fizz-idm/internal/data"
	"github.com/zerotohero-dev/fizz-idm/internal/downstream"
)

func main() {
	e := *env.New()

	appEnv := e.Idm

	svcName := appEnv.ServiceName

	go func() {
		// #region mTLS client
		ctx := context.Background()
		fmt.Println("before opening connection")
		conn, err := spiffetls.Dial(ctx, "tcp", "127.0.0.1:8443", tlsconfig.AuthorizeAny())
		fmt.Println("after opening connection")
		fmt.Println(conn)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Everything is awesome!")
		}
		// #engregion mTLS client
	}()


	// Configure the environment:
	app.Configure(e, svcName, appEnv.HoneybadgerApiKey, appEnv.Sanitize)
	// Connect to the database:
	data.Init(e)
	// Initialize downstream services:
	downstream.Init(e)

	r := mux.NewRouter()
	api.InitializeEndpoints(e, r)
	app.RouteHealthEndpoints(e.Idm.PathPrefix, r)

	app.ListenAndServe(e, svcName, appEnv.Port, r)
}
