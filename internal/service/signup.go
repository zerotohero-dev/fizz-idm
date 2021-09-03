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

package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	entity "github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-idm/internal/data"
	"github.com/zerotohero-dev/fizz-idm/internal/downstream"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
)

func createUnverifiedUser(user entity.User) error {
	exists, err := data.UserExists(user.Email)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"createUnverifiedUser: error checking the existence of the user (%s)",
				user.Email,
			),
		)
	}

	if exists {
		return errors.New(
			fmt.Sprintf(
				"createUnverifiedUser: user already exists in the db (%s)",
				user.Email,
			),
		)
	}

	return data.CreateUnverifiedUser(user)
}

func sendEmailVerificationToken(name, email string, emailVerificationToken string) {
	go func() {

		res, err := downstream.Endpoints().MailerVerification(

			// We are using context.Background() because we do not want this to
			// cancel prematurely. go-kit cancels the owner context as soon as
			// the function exits.
			context.Background(), reqres.RelayEmailVerificationMessageRequest{
				Email: email,
				Name:  name,
				Token: emailVerificationToken,
			})

		if err != nil {
			log.Err("Problem sending activation email (%s) (%s)",
				log.RedactEmail(email), err.Error())
		}

		er := res.(reqres.RelayEmailVerificationMessageResponse)
		if er.Err != "" {
			log.Err("Problem sending activation email (%s) (%s)",
				log.RedactEmail(email), er.Err)
		}
	}()
}

func sendWaitlistEmail(name, email string) {
	go func() {
		res, err := downstream.Endpoints().MailerVerification(

			// We are using context.Background() because we do not want this to
			// cancel prematurely. go-kit cancels the owner context as soon as
			// the function exits.
			context.Background(), reqres.RelayWelcomeMessageRequest{
				Email: email,
				Name:  name,
			})

		if err != nil {
			log.Err("Problem sending Waitlist email (%s) (%s)",
				log.RedactEmail(email), err.Error())
		}

		er := res.(reqres.RelayWelcomeMessageResponse)
		if er.Err != "" {
			log.Err("Problem sending Waitlist email (%s) (%s)",
				log.RedactEmail(email), er.Err)
		}
	}()
}

func (s service) SignUp(user entity.User) error {
	res, err := downstream.Endpoints().CryptoTokenCreate(
		s.Context(), reqres.TokenCreateRequest{})
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"SignUp: error requesting account activation token (%s)",
				user.Email,
			),
		)
	}

	tr, ok := res.(reqres.TokenCreateResponse)
	if tr.Err != "" {
		return errors.New(
			fmt.Sprintf(
				"SignUp: Error in TokenResponse %s (%s)",
				tr.Err, user.Email,
			),
		)
	}

	if !ok {
		return errors.New(
			fmt.Sprintf(
				"SignUp: error creating account activation token (%s)",
				user.Email,
			),
		)
	}

	if tr.Token == "" {
		return errors.New(
			fmt.Sprintf(
				"SignUp: error creating account activation token (%s)",
				user.Email,
			),
		)
	}

	user.EmailVerificationToken = tr.Token

	err = createUnverifiedUser(user)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SignUp: error creating user (%s)", user.Email),
		)
	}

	launchState := env.FizzEnv().Idm.LaunchState

	if launchState == "waitlist" {
		sendWaitlistEmail(user.Name, user.Email)
	} else {
		sendEmailVerificationToken(user.Name, user.Email, tr.Token)
	}

	return nil
}
