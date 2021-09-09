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
	"github.com/zerotohero-dev/fizz-idm/internal/downstream"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
)

func (s service) Info(authToken string) (entity.UserInfo, error) {
	res, err := downstream.Endpoints().CryptoJwtVerify(
		s.ctx, reqres.JwtVerifyRequest{Token: authToken},
	)

	if err != nil {
		return data.UserInfo{}, errors.New("Info: Invalid auth token.")
	}

	vr := res.(reqres.JwtVerifyResponse)
	if vr.Err != "" {
		return data.UserInfo{}, errors.New(fmt.Sprintf("Info: %s", vr.Err))
	}

	if !vr.Valid {
		return data.UserInfo{}, errors.New(
			fmt.Sprintf(
				"Info: User does not seem to be valid (%s)",
				log.RedactEmail(vr.Email),
			),
		)
	}

	email := vr.Email
	user, err := data.ActiveUserByEmail(email)

	if err != nil {
		return entity.UserInfo{}, errors.New(fmt.Sprintf("Info: Error querying the database for user (%s)", log.RedactEmail(vr.Email)))
	}

	if user.Email == "" {
		return entity.UserInfo{}, errors.New(fmt.Sprintf("Info: Got blank user (%s)", log.RedactEmail(vr.Email)))
	}

	return entity.UserInfo{
		Email:    user.Email,
		FullName: user.FullName,
		Alias:    user.Alias,
		Courses:  user.Courses,
	}, nil
}
