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

func TestCryptoHashCreate(t *testing.T) {
	e := *env.New()
	mtls.Init(e)

	value := "potato"
	res, err := mtls.CryptoHashCreate(reqres.HashCreateRequest{
		Value: value,
	})

	if err != nil {
		t.Fatal("Error creating token:", err.Error())
		return
	}

	if res == nil {
		t.Fatal("nil response")
		return
	}

	t.Log("Generated hash:", res.Hash)

	vr, err := mtls.CryptoHashVerify(reqres.HashVerifyRequest{
		Value: value,
		Hash:  res.Hash,
	})

	if err != nil {
		t.Fatal("Error verifying hash:", err.Error())
		return
	}

	if vr == nil {
		t.Fatal("nil response")
		return
	}

	t.Log("verified:", vr.Verified)
	t.Log("Done.")
}
