package transport

import (
	"context"
	"encoding/json"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net/http"
	"time"
)

func EncodeLoginResponse(e env.FizzEnv) kitHttp.EncodeResponseFunc {
	return func(_ context.Context, w http.ResponseWriter, response interface{}) error {
		responseErr := app.ToErrorString(response)
		if responseErr != "" {
			log.Err("EncodeLoginResponse: error encoding response: %s", responseErr)

			res := reqres.GenericResponse{
				Err: "There is a problem in your request.",
			}

			w.WriteHeader(http.StatusBadRequest)

			return json.NewEncoder(w).Encode(res)
		}

		rr := response.(reqres.LogInResponse)

		// Make this a central function; it is repeated!
		expiration := time.Now().Add(e.Idm.JwtCookieExpiration)
		cookie := http.Cookie{
			Name: "auth",
			// Path / required for other services to consume this cookie.
			Path:     "/",
			Value:    rr.AuthToken,
			Expires:  expiration,
			SameSite: http.SameSiteNoneMode,
			HttpOnly: true,
			// Set it to Secure: false for local dev testing (if using http)
			Secure: true,
		}
		http.SetCookie(w, &cookie)

		// Override auth token: Client should not see it in the json response.
		// It should be visible in the cookie header only.
		rr.AuthToken = "<c>"

		return json.NewEncoder(w).Encode(rr)
	}
}
