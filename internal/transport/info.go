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
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"net/http"
)

func DecodeInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request reqres.UserInfoRequest

	authCookie, _ := r.Cookie("auth")
	if authCookie != nil {
		request.AuthToken = authCookie.Value
	}

	return request, nil
}
