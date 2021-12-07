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

package mtls

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"io"
	"net"
)

func serialize(request interface{}) string {
	if !initialized {
		panic("serialize: mTLS service has not been initialized")
	}

	body, err := json.Marshal(request)
	if err != nil {
		log.Err("serialize: Problem serializing request: %s", err.Error())
	}
	return string(body)
}

func deserialize(conn net.Conn, response interface{}) error {
	res, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil && err != io.EOF {
		return errors.Wrap(err, "deserialize: Problem reading response")
	}

	err = json.Unmarshal([]byte(res), response)
	if err != nil {
		return errors.Wrap(
			err, "deserialize: Problem parsing downstream response",
		)
	}

	return nil
}

func send(conn net.Conn, payload string) error {
	if conn == nil {
		return errors.New("send: Failed to connect downstream")
	}

	_, err := fmt.Fprintf(conn, payload+"\n")
	if err != nil {
		return errors.Wrap(err, "send: Problem sending payload")
	}

	return nil
}

func connect(appName, serverAddress string) (net.Conn, context.CancelFunc) {
	log.Info("connect: (%s) (%s)", appName, serverAddress)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)

	spiffeId := spiffeid.Must("fizzbuzz.pro", "app", appName)

	conn, err := spiffetls.DialWithMode(ctx, "tcp", serverAddress,
		spiffetls.MTLSClientWithSourceOptions(
			tlsconfig.AuthorizeID(spiffeId),
			workloadapi.WithClientOptions(workloadapi.WithAddr(spireSocketPath)),
		))

	if err != nil {
		log.Err("connect: Error connecting: %s", err.Error())
		return nil, cancel
	}

	return conn, cancel
}

func connectMailer() (net.Conn, context.CancelFunc) {
	return connect(spireAppNameMailer, serverAddressMailer)
}

func connectCrypto() (net.Conn, context.CancelFunc) {
	return connect(spireAppNameCrypto, serverAddressCrypto)
}

func disconnect(conn net.Conn, cancel context.CancelFunc) func() {
	return func() {
		cancel()
		if conn == nil {
			log.Info("mTLS: Unable to close nil connection")
			return
		}
		err := conn.Close()
		if err != nil {
			log.Info("mTLS: Unable to cleanly close connection: %s", err.Error())
			return
		}
	}
}
