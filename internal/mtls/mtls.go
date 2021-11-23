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
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"time"
)

var serverAddressCrypto string
var serverAddressMailer string
var spireSocketPath string
var spireAppNameCrypto string
var spireAppNameMailer string
var requestTimeout time.Duration
var initialized = false

func Init(e env.FizzEnv) {
	if initialized {
		return
	}

	isDevelopment := e.Deployment.Type == env.Development

	serverAddressCrypto = e.Idm.MtlsServerAddressCrypto
	serverAddressMailer = e.Idm.MtlsServerAddressMailer
	spireSocketPath = e.Spire.SocketPath
	requestTimeout = e.Spire.MtlsTimeout

	if isDevelopment {
		spireAppNameCrypto = e.Spire.AppNameFizzDefault
		spireAppNameMailer = e.Spire.AppNameFizzDefault
	} else {
		spireAppNameCrypto = e.Crypto.ServiceName
		spireAppNameMailer = e.Mailer.ServiceName
	}

	initialized = true
}
