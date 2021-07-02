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

package transport

import (
	"context"
	"encoding/json"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net/http"
)

func DecodeSignupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request reqres.SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Err("DecodeSignupRequest: %s", err.Error())

		request.Err = "DecodeSignupRequest: invalid payload."
	}

	return request, nil
}
