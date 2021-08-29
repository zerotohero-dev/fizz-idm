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
	"github.com/zerotohero-dev/fizz-idm/internal/data"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
)

func (s service) CreateAccount(user entity.User) error {
	u, err := data.UnverifiedUserByEmailAndEmailVerificationToken(
		user.Email, user.EmailVerificationToken,
	)

	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"error getting unverified user by email '%s'",
				log.RedactEmail(user.Email),
			),
		)
	}

	if u == nil {
		return errors.New(fmt.Sprintf(
			"error getting unverified user by email '%s'",
			log.RedactEmail(user.Email),
		))
	}

	// Update status, and any other attributes that might be passed.
	err = data.SetUserVerified(user)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"error verifying the user: '%s'",
				log.RedactEmail(user.Email),
			),
		)
	}

	return nil
}
