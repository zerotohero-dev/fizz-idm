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
	"github.com/pkg/errors"
	"github.com/zerotohero-dev/fizz-entity/pkg/endpoint"
	"github.com/zerotohero-dev/fizz-entity/pkg/method"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
)

func MailerWelcome(request reqres.RelayWelcomeMessageRequest) (
	*reqres.RelayWelcomeMessageResponse, error) {

	payload := serialize(reqres.MtlsApiRequest{
		Service:  reqres.MailerService,
		Endpoint: endpoint.Mailer.Welcome,
		Method:   method.Post,
		Body:     serialize(request),
	})

	conn, cancel := connectMailer()
	defer disconnect(conn, cancel)()

	err := send(conn, payload)
	if err != nil {
		return nil, errors.Wrap(err, "MailerWelcome: Failed to send request")
	}

	var mr reqres.RelayWelcomeMessageResponse
	err = deserialize(conn, &mr)
	if err != nil {
		return nil, errors.Wrap(err, "MailerWelcome: Problem receiving response")
	}

	return &mr, nil
}

func MailerVerify(request reqres.RelayEmailVerificationMessageRequest) (
	*reqres.RelayEmailVerificationMessageResponse, error) {

	payload := serialize(reqres.MtlsApiRequest{
		Service:  reqres.MailerService,
		Endpoint: endpoint.Mailer.Verification,
		Method:   method.Post,
		Body:     serialize(request),
	})

	conn, cancel := connectMailer()
	defer disconnect(conn, cancel)()

	err := send(conn, payload)
	if err != nil {
		return nil, errors.Wrap(err, "MailerVerify: Failed to send request")
	}

	var mr reqres.RelayEmailVerificationMessageResponse
	err = deserialize(conn, &mr)
	if err != nil {
		return nil, errors.Wrap(err, "MailerVerify: Problem receiving response")
	}

	return &mr, nil
}
