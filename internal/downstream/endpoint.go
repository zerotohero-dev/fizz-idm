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

var urls = struct {
	CryptoTokenCreate  string
	CryptoHashCreate   string
	CryptoHashVerify   string
	CryptoJwtCreate    string
	CryptoJwtVerify    string
	MailerWelcome      string
	MailerVerification string
	MailerReset        string
	MailerConfirm      string
	MailerSubscribed   string
}{
	CryptoTokenCreate:  "/v1/token",
	CryptoHashCreate:   "/v1/hash",
	CryptoHashVerify:   "/v1/hash/verify",
	CryptoJwtCreate:    "/v1/jwt",
	CryptoJwtVerify:    "/v1/jwt/verify",
	MailerWelcome:      "/v1/relay/welcome",
	MailerVerification: "/v1/relay/verification",
	MailerReset:        "/v1/relay/reset",
	MailerConfirm:      "/v1/relay/confirm",
	MailerSubscribed:   "/v1/relay/subscribed",
}

type downstream struct {
	CryptoTokenCreate  endpoint.Endpoint
	CryptoHashCreate   endpoint.Endpoint
	CryptoHashVerify   endpoint.Endpoint
	CryptoJwtCreate    endpoint.Endpoint
	CryptoJwtVerify    endpoint.Endpoint
	MailerWelcome      endpoint.Endpoint
	MailerVerification endpoint.Endpoint
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
	case urls.CryptoTokenCreate:
		return http.NewClient(
			"GET", u, encodeRequest, decodeTokenCreateResponse).Endpoint()
	case urls.CryptoHashCreate:
		return http.NewClient(
			"POST", u, encodeRequest, decodeHashCreateResponse).Endpoint()
	case urls.CryptoHashVerify:
		return http.NewClient(
			"POST", u, encodeRequest, decodeHashVerifyResponse).Endpoint()
	case urls.CryptoJwtCreate:
		return http.NewClient(
			"POST", u, encodeRequest, decodeJwtCreateResponse).Endpoint()
	case urls.CryptoJwtVerify:
		return http.NewClient(
			"POST", u, encodeRequest, decodeJwtVerifyResponse).Endpoint()
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
	case urls.MailerWelcome:
		return http.NewClient("POST", u, encodeRequest,
			decodeRelayWelcomeMessageResponse).Endpoint()
	case urls.MailerVerification:
		return http.NewClient("POST", u, encodeRequest,
			decodeRelaySendEmailVerificationMessageResponse).Endpoint()
	case urls.MailerReset:
		panic("Implement me")
	case urls.MailerConfirm:
		panic("Implement me")
	case urls.MailerSubscribed:
		panic("Implement me")
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
		CryptoTokenCreate:  makeCryptoEndpoint(en, urls.CryptoTokenCreate),
		CryptoHashCreate:   makeCryptoEndpoint(en, urls.CryptoHashCreate),
		CryptoHashVerify:   makeCryptoEndpoint(en, urls.CryptoHashVerify),
		CryptoJwtCreate:    makeCryptoEndpoint(en, urls.CryptoJwtCreate),
		CryptoJwtVerify:    makeCryptoEndpoint(en, urls.CryptoJwtVerify),
		MailerWelcome:      makeMailerEndpoint(en, urls.MailerWelcome),
		MailerVerification: makeMailerEndpoint(en, urls.MailerVerification),
		// TODO: there are endpoints to implement, you’ll probably need them too.
	}
}
