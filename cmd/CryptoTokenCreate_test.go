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
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-idm/internal/mtls"
	"testing"
)

func TestCryptoTokenCreate(t *testing.T) {
	e := *env.New()
	mtls.Init(e)

	res, err := mtls.CryptoTokenCreate(reqres.TokenCreateRequest{})
	if err != nil {
		t.Fatal("Error creating token:", err.Error())
		return
	}

	if res == nil {
		t.Fatal("nil response")
		return
	}

	// The token that we got:
	t.Log("Generated token:", res.Token)
	t.Log("Done.")
}
