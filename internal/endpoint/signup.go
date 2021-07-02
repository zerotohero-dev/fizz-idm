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
	"github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-idm/internal/service"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-validation/pkg/sanitization"
	"time"
)
const defaultName = "FizzBuzz Pro"

func MakeSignupEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		gr, ok := request.(reqres.ContentTypeProblemRequest)
		if ok {
			return reqres.SignUpResponse{
				Err: gr.Err,
			}, nil
		}

		req := request.(reqres.SignUpRequest)
		if req.Err != "" {
			return reqres.SignUpResponse{
				Err: req.Err,
			}, nil
		}

		req.Name =  sanitization.SanitizeName(req.Name)
		req.Email = sanitization.SanitizeEmail(req.Email)

		if req.Name == "" {
			req.Name = defaultName
		}

		if req.Email == "" {
			return reqres.SignUpResponse{
				Err: "signUpEndpoint: email required",
			}, nil
		}

		now := (time.Now().UnixNano()) / 1000000
		err := svc.SignUp(data.User{
			Info:                    data.Info{
				Email: req.Email,
				Name:  req.Name,
			},
			Password:                "",
			Status:                  "",
			SubscribedToMailingList: false,
			RecordCreated:           now,
			RecordUpdated:           now,
		})


		if err != nil {
			log.Err("signUpEndpoint: %s", err.Error())

			return reqres.SignUpResponse{
				Err: "signUpEndpoint: cannot sign up user",
			}, nil
		}

		return reqres.SignUpResponse{}, nil
	}
}