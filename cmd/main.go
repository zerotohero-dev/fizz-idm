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
	"bufio"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-idm/internal/api"
	"github.com/zerotohero-dev/fizz-idm/internal/data"
	"github.com/zerotohero-dev/fizz-idm/internal/downstream"
	"io"
	"log"
	"net"
	"time"
)

//// #region mTLS client
//ctx := context.Background()
//fmt.Println("before opening connection")
//conn, err := spiffetls.Dial(ctx, "tcp", "127.0.0.1:8443", tlsconfig.AuthorizeAny())
//fmt.Println("after opening connection")
//fmt.Println(conn)
//if err != nil {
//panic(err)
//} else {
//fmt.Println("Everything is awesome!")
//}
//// #engregion mTLS client

const (
	socketPath    = "unix:///tmp/spire-agent/public/api.sock"
	serverAddress = "localhost:55553"
)

func startSpireMtlsClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 3000000*time.Second)
	defer cancel()

	spiffeId := spiffeid.Must("fizzbuzz.pro", "app", "idm")

	fmt.Println(spiffeId)

	// SPIFFE_ENDPOINT_SOCKET=unix:///tmp/spire-agent/public/api.sock

	conn, err := spiffetls.DialWithMode(ctx, "tcp", serverAddress,
		spiffetls.MTLSClientWithSourceOptions(
			tlsconfig.AuthorizeID(spiffeId),
			// tlsconfig.AuthorizeAny(),
			workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)),
		))
	if err != nil {
		log.Fatalf("Unable to create TLS connection: %v", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)

	_, _ = fmt.Fprintf(conn, "Hello server\n")

	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatalf("Unable to read server response: %v", err)
	}
	log.Printf("Server says: %q", status)
}

func main() {
	e := *env.New()

	appEnv := e.Idm

	svcName := appEnv.ServiceName

	// Configure the environment:
	app.Configure(e, svcName, appEnv.HoneybadgerApiKey, appEnv.Sanitize)
	// Connect to the database:
	data.Init(e)
	// Initialize downstream services:
	downstream.Init(e)

	r := mux.NewRouter()
	api.InitializeEndpoints(e, r)
	app.RouteHealthEndpoints(e.Idm.PathPrefix, r)

	go startSpireMtlsClient()

	app.ListenAndServe(e, svcName, appEnv.Port, r)
}
