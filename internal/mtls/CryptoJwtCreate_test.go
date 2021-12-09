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

func TestCryptoJwtCreate(t *testing.T) {
	e := *env.New()
	Init(e)

	res, err := CryptoJwtCreate(reqres.JwtCreateRequest{
		Email: "me@volkan.io",
	})

	if err != nil {
		t.Fatal("Error creating jwt:", err.Error())
		return
	}

	if res == nil {
		t.Fatal("nil response")
		return
	}

	t.Log("Generated jwt:", res.Token)

	vr, err := CryptoJwtVerify(reqres.JwtVerifyRequest{
		Token: res.Token,
	})

	if err != nil {
		t.Fatal("Error verifying jwt:", err.Error())
		return
	}

	if vr == nil {
		t.Fatal("nil response")
		return
	}

	t.Log("verified:", vr.Email, vr.Valid, vr.Expires)
	t.Log("Done.")
}
