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

func (s service) Info(authToken string) (entity.UserInfo, error) {
	res, err := mtls.CryptoJwtVerify(reqres.JwtVerifyRequest{
		Token: authToken,
	})

	if err != nil {
		return entity.UserInfo{}, errors.New("Info: Invalid auth token.")
	}

	if res.Err != "" {
		return entity.UserInfo{}, errors.New(fmt.Sprintf("Info: %s", res.Err))
	}

	if !res.Valid {
		return entity.UserInfo{}, errors.New(
			fmt.Sprintf(
				"Info: User does not seem to be valid (%s)",
				// TODO: The logger library should internally do this.
				log.RedactEmail(res.Email),
			),
		)
	}

	email := res.Email
	user, err := data.VerifiedUserByEmail(email)
	if err != nil {
		return entity.UserInfo{}, errors.New(
			fmt.Sprintf(
				"Info: Error querying the database for user (%s)",
				log.RedactEmail(res.Email),
			),
		)
	}

	if user.Email == "" {
		return entity.UserInfo{}, errors.New(
			fmt.Sprintf("Info: Got blank user (%s)", log.RedactEmail(res.Email)),
		)
	}

	subscriptionId := ""
	if user.StripeSubscription != nil {
		subscriptionId = user.StripeSubscription.StripeProductId
	}

	return entity.UserInfo{
		Email:        user.Email,
		Name:         user.Name,
		Subscription: subscriptionId,
	}, nil
}
