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
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-idm/internal/api"
	"github.com/zerotohero-dev/fizz-idm/internal/data"
	"github.com/zerotohero-dev/fizz-idm/internal/mtls"
)

func main() {
	e := *env.New()

	appEnv := e.Idm
	svcName := appEnv.ServiceName

	// Initial setup:
	configureApp(svcName, e)
	mtls.Init(e)

	// Connect to the database:
	data.Init(e)

	r := mux.NewRouter()
	api.InitializeEndpoints(e, r)
	app.RouteHealthEndpoints(e.Idm.PathPrefix, r)

	// For demo only!
	go func() {
		value := "potato"
		res, err := mtls.CryptoHashCreate(reqres.HashCreateRequest{
			Value: value,
		})

		if err != nil || res == nil {
			return
		}

		fmt.Println("Generated hash:", res.Hash)

		vr, err := mtls.CryptoHashVerify(reqres.HashVerifyRequest{
			Value: value,
			Hash:  res.Hash,
		})

		if err != nil || vr == nil {
			return
		}

		fmt.Println("Verified:", vr.Verified)
		fmt.Println("Done.")
	}()

	app.ListenAndServe(e, svcName, appEnv.Port, r)
}
