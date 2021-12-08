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
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-idm/internal/api"
	"github.com/zerotohero-dev/fizz-idm/internal/data"
	"github.com/zerotohero-dev/fizz-idm/internal/mtls"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
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

	// This is for testing purposes only!
	go func() {
		log.Info("Will try to create a token for a demo…")

		// Establish an mTLS connection with the crypto service to get a token.
		res, err := mtls.CryptoTokenCreate(reqres.TokenCreateRequest{})

		if err != nil {
			log.Err("Error creating token: %s", err.Error())
			return
		}

		if res == nil {
			log.Err("nil response")
			return
		}

		// The token that we got:
		log.Info("Generated token: %s", res.Token)
	}()

	go func() {
		value := "potato"
		res, err := mtls.CryptoHashCreate(reqres.HashCreateRequest{
			Value: value,
		})

		if err != nil {
			log.Err("Error creating token: %s", err.Error())
			return
		}

		if res == nil {
			log.Err("nil response")
			return
		}

		log.Info("Generated hash: %s", res.Hash)

		vr, err := mtls.CryptoHashVerify(reqres.HashVerifyRequest{
			Value: value,
			Hash:  res.Hash,
		})

		if err != nil {
			log.Err("Error creating token: %s", err.Error())
			return
		}

		if vr == nil {
			log.Err("nil response")
			return
		}

		log.Info("verified: %s", vr.Verified)
	}()

	func() {
		res, err := mtls.CryptoJwtCreate(reqres.JwtCreateRequest{
			Email: "me@volkan.io",
		})

		if err != nil {
			log.Err("Error creating jwt: %s", err.Error())
			return
		}

		if res == nil {
			log.Err("nil response")
			return
		}

		log.Info("verified: %s", res.Token)

		vr, err := mtls.CryptoJwtVerify(reqres.JwtVerifyRequest{
			Token: res.Token,
		})

		if err != nil {
			log.Err("Error creating jwt: %s", err.Error())
			return
		}

		if vr == nil {
			log.Err("nil response")
			return
		}

		log.Info("verified: %s %s %s", vr.Email, vr.Expires, vr.Valid)
	}()

	app.ListenAndServe(e, svcName, appEnv.Port, r)
}
