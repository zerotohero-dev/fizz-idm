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

func CryptoTokenCreate(request reqres.TokenCreateRequest) (*reqres.TokenCreateResponse, error) {
	payload := serialize(reqres.MtlsApiRequest{
		Service:  reqres.CryptoService,
		Endpoint: endpoint.Crypto.SecureToken,
		Method:   method.Get,
		Body:     serialize(request),
	})

	conn, cancel := connectCrypto()
	defer disconnect(conn, cancel)()

	err := send(conn, payload)
	if err != nil {
		return nil, errors.Wrap(err, "CryptoTokenCreate: Failed to send request")
	}

	var tr reqres.TokenCreateResponse
	err = deserialize(conn, &tr)
	if err != nil {
		return nil, errors.Wrap(err, "CryptoTokenCreate: Problem receiving response")
	}

	return &tr, nil
}

func CryptoHashCreate(request reqres.HashCreateRequest) (*reqres.HashCreateResponse, error) {
	payload := serialize(reqres.MtlsApiRequest{
		Service:  reqres.CryptoService,
		Endpoint: endpoint.Crypto.SecureHash,
		Method:   method.Post,
		Body:     serialize(request),
	})

	conn, cancel := connectCrypto()
	defer disconnect(conn, cancel)()

	err := send(conn, payload)
	if err != nil {
		return nil, errors.Wrap(err, "CryptoHashCreate: Failed to send request")
	}

	var hr reqres.HashCreateResponse
	err = deserialize(conn, &hr)
	if err != nil {
		return nil, errors.Wrap(err, "CryptoHashCreate: Problem receiving response")
	}

	return &hr, nil
}

func CryptoHashVerify(request reqres.HashVerifyRequest) (*reqres.HashVerifyResponse, error) {
	payload := serialize(reqres.MtlsApiRequest{
		Service:  reqres.CryptoService,
		Endpoint: endpoint.Crypto.SecureHashVerify,
		Method:   method.Post,
		Body:     serialize(request),
	})

	conn, cancel := connectCrypto()
	defer disconnect(conn, cancel)()

	err := send(conn, payload)
	if err != nil {
		return nil, errors.Wrap(err, "CryptoHashVerify: Failed to send request")
	}

	var hr reqres.HashVerifyResponse
	err = deserialize(conn, &hr)
	if err != nil {
		return nil, errors.Wrap(err, "CryptoHashVerify: Problem receiving response")
	}

	return &hr, nil
}

func CryptoJwtCreate(request reqres.JwtCreateRequest) (*reqres.JwtCreateResponse, error) {
	payload := serialize(reqres.MtlsApiRequest{
		Service:  reqres.CryptoService,
		Endpoint: endpoint.Crypto.Jwt,
		Method:   method.Post,
		Body:     serialize(request),
	})

	conn, cancel := connectCrypto()
	defer disconnect(conn, cancel)()

	err := send(conn, payload)
	if err != nil {
		return nil, errors.Wrap(err, "CryptoJwtCreate: Failed to send request")
	}

	var jr reqres.JwtCreateResponse
	err = deserialize(conn, &jr)
	if err != nil {
		return nil, errors.Wrap(err, "CryptoJwtCreate: Problem receiving response")
	}

	return &jr, nil
}

func CryptoJwtVerify(request reqres.JwtVerifyRequest) (*reqres.JwtVerifyResponse, error) {
	payload := serialize(reqres.MtlsApiRequest{
		Service:  reqres.CryptoService,
		Endpoint: endpoint.Crypto.JwtVerify,
		Method:   method.Post,
		Body:     serialize(request),
	})

	conn, cancel := connectCrypto()
	defer disconnect(conn, cancel)()

	err := send(conn, payload)
	if err != nil {
		return nil, errors.Wrap(err, "CryptoJwtVerify: Failed to send request")
	}

	var jr reqres.JwtVerifyResponse
	err = deserialize(conn, &jr)
	if err != nil {
		return nil, errors.Wrap(err, "CryptoJwtVerify: Problem receiving response")
	}

	return &jr, nil
}
