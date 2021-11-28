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
	"fmt"
	"github.com/pkg/errors"
	entity "github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-idm/internal/data"
	"github.com/zerotohero-dev/fizz-idm/internal/mtls"
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
	log.Info("implement me:", name, email)

	res, err := mtls.MailerVerify(reqres.RelayEmailVerificationMessageRequest{
		Email: email,
		Name:  name,
		Token: emailVerificationToken,
	})

	if err != nil {
		log.Err("Problem sending activation email (%s) (%s)",
			log.RedactEmail(email), err.Error())
	}

	if res.Err != "" {
		log.Err("Problem sending activation email (%s) (%s)",
			log.RedactEmail(email), res.Err)
	}
}

func (s service) SignUp(user entity.User) error {
	res, err := mtls.CryptoTokenCreate(reqres.TokenCreateRequest{})

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"SignUp: error requesting account activation token (%s)",
				user.Email,
			),
		)
	}

	if res.Err != "" {
		return errors.New(
			fmt.Sprintf(
				"SignUp: Error in TokenResponse %s (%s)",
				res.Err, user.Email,
			),
		)
	}

	if res.Token == "" {
		return errors.New(
			fmt.Sprintf(
				"SignUp: error creating account activation token (%s)",
				user.Email,
			),
		)
	}

	user.EmailVerificationToken = res.Token

	err = createUnverifiedUser(user)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf("SignUp: error creating user (%s)", user.Email),
		)
	}

	go sendEmailVerificationToken(user.Name, user.Email, user.EmailVerificationToken)

	return nil
}
