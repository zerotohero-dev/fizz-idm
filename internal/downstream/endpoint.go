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

package downstream

import (
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net/url"
)

type downstream struct {
	CryptoTokenCreate  endpoint.Endpoint
	CryptoHashCreate   endpoint.Endpoint
	CryptoHashVerify   endpoint.Endpoint
	CryptoJwtCreate    endpoint.Endpoint
	CryptoJwtVerify    endpoint.Endpoint
	MailerWelcome      endpoint.Endpoint
	MailerVerification endpoint.Endpoint
	MailerVerified     endpoint.Endpoint
}

func join(baseUrl, path string) *url.URL {
	if len(baseUrl) == 0 {
		return nil
	}

	if len(path) == 0 {
		return nil
	}

	if baseUrl[len(baseUrl)-1] != '/' {
		baseUrl = baseUrl + "/"
	}

	if path[0] == '/' {
		path = path[1:]
	}

	instance := fmt.Sprintf("%s%s", baseUrl, path)

	u, err := url.Parse(instance)
	if err != nil {
		log.Err("makeCryptoEndpoint: unable to parse URL for path '%s'.", path)
		return nil
	}

	return u
}

func makeCryptoEndpoint(en env.FizzEnv, path string) endpoint.Endpoint {
	baseUrl := en.Idm.CryptoEndpointUrl
	u := join(baseUrl, path)

	if u == nil {
		log.Err("makeCryptoEndpoint: unable to parse URL for path '%s'.", path)
		return nil
	}

	switch path {
	case "/v1/token/create":
		return http.NewClient(
			"POST", u, encodeRequest, decodeTokenCreateResponse).Endpoint()
	case "/v1/hash/create":
		return http.NewClient(
			"POST", u, encodeRequest, decodeHashCreateResponse).Endpoint()
	case "/v1/jwt/sign":
		return http.NewClient(
			"POST", u, encodeRequest, decodeJwtCreateResponse).Endpoint()
	case "/v1/jwt/verify":
		return http.NewClient(
			"POST", u, encodeRequest, decodeJwtVerifyResponse).Endpoint()
	case "/v1/hash/verify":
		return http.NewClient(
			"POST", u, encodeRequest, decodeHashVerifyResponse).Endpoint()
	default:
		log.Warning("makeCryptoEndpoint: Unknown path '%s'", path)
		return nil
	}
}

func makeMailerEndpoint(en env.FizzEnv, path string) endpoint.Endpoint {
	baseUrl := en.Idm.MailerEndpointUrl
	u := join(baseUrl, path)

	if u == nil {
		log.Err("makeCryptoEndpoint: unable to parse URL for path '%s'.", path)
		return nil
	}

	switch path {
	case "/v1/send/welcome":
		return http.NewClient("POST", u, encodeRequest,
			decodeRelayWelcomeMessageResponse).Endpoint()
	case "/v1/send/verification":
		return http.NewClient("POST", u, encodeRequest,
			decodeRelaySendEmailVerificationMessageResponse).Endpoint()
	case "/v1/send/verified":
		return http.NewClient("POST", u, encodeRequest,
			decodeRelayEmailVerifiedEmailResponse).Endpoint()
	default:
		log.Warning("makeCryptoEndpoint: Unknown path '%s'", path)
		return nil
	}
}

var e *downstream

func Endpoints() *downstream { return e }

func Init(en env.FizzEnv) {
	if e != nil {
		return
	}

	e = &downstream{
		CryptoTokenCreate:  makeCryptoEndpoint(en, "/v1/token/create"),
		CryptoHashCreate:   makeCryptoEndpoint(en, "/v1/hash/create"),
		CryptoHashVerify:   makeCryptoEndpoint(en, "/v1/hash/verify"),
		CryptoJwtCreate:    makeCryptoEndpoint(en, "/v1/jwt/create"),
		CryptoJwtVerify:    makeCryptoEndpoint(en, "/v1/jwt/verify"),
		MailerWelcome:      makeMailerEndpoint(en, "/v1/send/welcome"),
		MailerVerification: makeMailerEndpoint(en, "/v1/send/verification"),
		MailerVerified:     makeMailerEndpoint(en, "/v1/send/verified"),
	}
}
