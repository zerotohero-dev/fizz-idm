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

package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-idm/internal/service"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-validation/pkg/sanitization"
)

func MakeLoginEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		gr, hasContentTypeProblem := request.(reqres.ContentTypeProblemRequest)
		if hasContentTypeProblem {
			return reqres.LogInResponse{
				Err: gr.Err,
			}, nil
		}

		req := request.(reqres.LogInRequest)
		if req.Err != "" {
			return reqres.LogInResponse{
				Err: req.Err,
			}, nil
		}

		email := sanitization.SanitizeEmail(req.Email)
		password := sanitization.SanitizePassword(req.Password)

		if len(email) == 0 {
			return reqres.LogInResponse{
				Err: "loginEndpoint: missing email",
			}, nil
		}

		if len(password) == 0 {
			return reqres.LogInResponse{
				Err: "loginEndpoint: blank password",
			}, nil
		}

		result, err := svc.LogIn(email, password)
		if err != nil {
			log.Err("loginEndpoint: error logging in user", err.Error())

			return reqres.LogInResponse{
				Err: "loginEndpoint: cannot authenticate user",
			}, nil
		}

		return reqres.LogInResponse{
			AuthToken: result.Token,
		}, nil
	}
}
