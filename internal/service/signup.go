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
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-idm/internal/data"
	"github.com/zerotohero-dev/fizz-idm/internal/downstream"
)


type Status int

const (
	Waitlist Status = iota
	Signup
)

func (state Status) String() string {
	return [...]string{"waitlist", "signup"}[state]
}

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

	launchState := env.New().Idm.LaunchState

	if launchState == Waitlist.String() {
		sendWaitlistEmail(user.Name, user.Email)
	} else {
		sendEmailVerificationToken(user.Name, user.Email, tr.Token)
	}

	return nil
}
