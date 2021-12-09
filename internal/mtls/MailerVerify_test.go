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

func TestMailerVerify(t *testing.T) {
	e := *env.New()
	Init(e)

	request := reqres.RelayEmailVerificationMessageRequest{
		Email: "me@volkan.io",
		Name:  "Volkan",
		Token: "potato",
	}

	res, err := MailerVerify(request)
	if err != nil {
		t.Fatal("Error relaying message:", err.Error())
		return
	}

	if res == nil {
		t.Fatal("nil response")
		return
	}

	// The token that we got:
	t.Log("Success?", res.Success)
	t.Log("Done.")

}
