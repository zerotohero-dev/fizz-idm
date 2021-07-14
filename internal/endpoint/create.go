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
	entity "github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-idm/internal/service"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-validation/pkg/sanitization"
)

func MakeCreateAccountEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		gr, ok := request.(reqres.ContentTypeProblemRequest)

		if ok {
			return reqres.CreateAccountResponse{
				Err: gr.Err,
			}, nil
		}

		req := request.(reqres.CreateAccountRequest)

		if req.Err != "" {
			return reqres.CreateAccountResponse{
				Err: req.Err,
			}, nil
		}

		req.Name = sanitization.SanitizeName(req.Name)
		req.Email = sanitization.SanitizeEmail(req.Email)
		req.Password = sanitization.SanitizePassword(req.Password)

		if req.Name == "" {
			req.Name = "FizzBuzz Pro"
		}

		if len(req.Password) < 6 {
			return reqres.CreateAccountResponse{
				Err: "MakeCreateAccountEndpoint: password should be at least six characters",
			}, nil
		}

		user, err := svc.CreateAccount(entity.User{
			Name:                    req.Name,
			Email:                   req.Email,
			Password:                req.Password,
			SubscribedToMailingList: req.SubscribeToMailingList,
		})

		if user == nil {
			if err != nil {
				log.Err(
					"MakeCreateAccountEndpoint: Error creating account: %s",
					err.Error(),
				)
			} else {
				log.Err("MakeCreateAccountEndpoint: Error creating account")
			}

			return reqres.CreateAccountResponse{
				Err: "MakeCreateAccountEndpoint: cannot sign up user",
			}, nil
		}

		if err != nil {
			log.Err("MakeCreateAccountEndpoint: %s", err.Error())

			return reqres.CreateAccountResponse{
				Err: "MakeCreateAccountEndpoint: cannot sign up user",
			}, nil
		}

		return reqres.CreateAccountResponse{}, nil
	}
}
