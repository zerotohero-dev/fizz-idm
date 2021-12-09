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
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"testing"
)

func TestCryptoTokenCreate(t *testing.T) {
	e := *env.New()
	Init(e)

	res, err := CryptoTokenCreate(reqres.TokenCreateRequest{})
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
