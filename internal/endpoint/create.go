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
	"github.com/zerotohero-dev/fizz-idm/internal/downstream"

	"github.com/go-kit/kit/endpoint"
	entity "github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-idm/internal/service"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-validation/pkg/sanitization"
)

func MakeCreateAccountEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		gr, contentTypeMismatch := request.(reqres.ContentTypeProblemRequest)

		if contentTypeMismatch {
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

		sanitizedName := sanitization.SanitizeName(req.Name)
		sanitizedEmail := sanitization.SanitizeEmail(req.Email)
		sanitizedPassword := sanitization.SanitizePassword(req.Password)
		sanitizedToken := sanitization.SanitizeToken(req.Token)

		if sanitizedName == "" {
			sanitizedName = sanitization.DefaultFullName
		}

		if len(sanitizedPassword) < sanitization.MinPasswordLength {
			return reqres.CreateAccountResponse{
				Err: "MakeCreateAccountEndpoint: password should be at least six characters",
			}, nil
		}

		// Hashing the password.
		res, err := downstream.Endpoints().CryptoHashCreate(
			svc.Context(),
			reqres.HashCreateRequest{
				Value: sanitizedPassword,
			},
		)

		if err != nil {
			log.Err("MakeCreateAccountEndpoint: %s", err.Error())

			return reqres.CreateAccountResponse{
				Err: "MakeCreateAccountEndpoint: cannot create an account",
			}, nil
		}

		cr := res.(reqres.HashCreateResponse)

		if cr.Err != "" {
			log.Err("MakeCreateAccountEndpoint: %s", cr.Err)

			return reqres.CreateAccountResponse{
				Err: "MakeCreateAccountEndpoint: cannot create an account",
			}, nil
		}

		hashedPassword := cr.Hash

		err = svc.CreateAccount(entity.User{
			Name:                    sanitizedName,
			Email:                   sanitizedEmail,
			Password:                hashedPassword,
			EmailVerificationToken:  sanitizedToken,
			SubscribedToMailingList: req.SubscribeToMailingList,
		})

		if err != nil {
			log.Err("MakeCreateAccountEndpoint: %s", err.Error())

			return reqres.CreateAccountResponse{
				Err: "MakeCreateAccountEndpoint: cannot create an account",
			}, nil
		}

		// Send the email in a separate process.
		go func() {
			res, err := downstream.Endpoints().MailerWelcome(

				// Using a separate context because this needs to stay alive
				// even after the create account API request finishes.
				context.Background(), reqres.RelayWelcomeMessageRequest{
					Email: req.Email,
					Name:  req.Name,
				})

			if err != nil {
				log.Err(
					"Problem sending welcome email (%s) (%s)",
					log.RedactEmail(req.Email), err.Error(),
				)

				return
			}

			er := res.(reqres.RelayWelcomeMessageResponse)
			if er.Err != "" {
				log.Err(
					"Problem sending welcome email (%s) (%s)",
					log.RedactEmail(req.Email), er.Err,
				)
			}
		}()

		return reqres.CreateAccountResponse{
			Verified: true,
		}, nil
	}
}
