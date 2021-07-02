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
	"github.com/zerotohero-dev/fizz-idm/internal/data"
	"github.com/zerotohero-dev/fizz-idm/internal/downstream"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
)

func createUnverifiedUser(user entity.User) error {
	exists, err := data.UserExists(user.Email)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SignUp: error checking the existence of the user (%s)", user.Email),
		)
	}

	if exists {
		return errors.New(
			fmt.Sprintf("SignUp: user already exists in the db (%s)", user.Email),
		)
	}

	return data.CreateUnverifiedUser(user)
}

func setAccountActivationToken(email string, emailVerificationToken string) error {
	u, err := data.UnverifiedUserByEmail(email)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("setAccountActivationToken: error checking the existence of the user (%s)", email),
		)
	}

	if u == nil {
		return errors.New(
			fmt.Sprintf("setAccountActivationToken: cannot find user to set token (%s)", email),
		)
	}

	err = data.UpdateUnverifiedUserEmailVerificationToken(u.Email, emailVerificationToken)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SetTokenAndPassword: error updating the token (%s)", email),
		)
	}

	go func() {
		// context.Background() because we do not want this to cancel prematurely.
		// go-kit cancels the owner context as soon as the function exits.
		res, err := downstream.Endpoints().MailerVerification(

			context.Background(), reqres.RelaySendEmailVerificationMessageRequest{
				Email: email,
				Name:  u.Name,
				Token: emailVerificationToken,
			})

		if err != nil {
			log.Err("Problem sending activation email (%s) (%s)", log.RedactEmail(email), err.Error())
		}

		er := res.(reqres.RelaySendEmailVerificationMessageRequest)
		if er.Err != "" {
			log.Err("Problem sending activation email (%s) (%s)", log.RedactEmail(email), er.Err)
		}
	}()

	return nil
}

func (s service) SignUp(user entity.User) error {
	err := createUnverifiedUser(user)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SignUp: error creating user (%s)", user.Email),
		)
	}

	res, err := downstream.Endpoints().CryptoTokenCreate(s.ctx, reqres.TokenCreateRequest{})
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SignUp: error requesting account activation token (%s)", user.Email),
		)
	}

	tr, ok := res.(reqres.TokenCreateResponse)
	if tr.Err != "" {
		return errors.New(
			fmt.Sprintf("SignUp: Error in TokenResponse %s (%s)", tr.Err, user.Email),
		)
	}

	if !ok {
		return errors.New(
			fmt.Sprintf("SignUp: error creating account activation token (%s)", user.Email),
		)
	}

	err = setAccountActivationToken(user.Email, tr.Token)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SignUp: error setting account activation token (%s)", user.Email),
		)
	}

	return nil
}
