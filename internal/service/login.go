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
)

func (s service) LogIn(email, password string) (entity.LoginResult, error) {
	user, err := data.VerifiedUserByEmail(email)
	if err != nil {
		return entity.LoginResult{
				Token: "",
			}, errors.Wrap(
				err,
				fmt.Sprintf("LogIn: error getting verified user from email (%s)", email),
			)
	}

	if user == nil {
		return entity.LoginResult{
			Token: "",
		}, errors.New(fmt.Sprintf("LogIn: User not found (%s)", email))
	}

	res, err := mtls.CryptoHashVerify(reqres.HashVerifyRequest{
		Value: password,
		Hash:  user.Password,
	})

	if err != nil {
		return entity.LoginResult{
				Token: "",
			}, errors.Wrap(
				err,
				fmt.Sprintf("LogIn: Error while matching password (%s)", email),
			)
	}

	if res.Err != "" {
		return entity.LoginResult{
				Token: "",
			}, errors.New(
				fmt.Sprintf("LogIn: Error in password verification: %s (%s)", res.Err, email),
			)
	}

	if !res.Verified {
		return entity.LoginResult{
			Token: "",
		}, errors.New(fmt.Sprintf("LogIn: Password mismatch (%s)", email))
	}

	jwtRes, err := mtls.CryptoJwtCreate(reqres.JwtCreateRequest{
		Email: user.Email,
	})

	if err != nil {
		return entity.LoginResult{
			Token: "",
		}, errors.Wrap(err, fmt.Sprintf("LogIn: cannot sign token (%s)", email))
	}

	if jwtRes.Err != "" {
		return entity.LoginResult{
				Token: "",
			}, errors.New(
				fmt.Sprintf("LogIn: Error in JWT sign Response %s (%s)", jwtRes.Err, email),
			)
	}

	token := jwtRes.Token
	if token == "" {
		return entity.LoginResult{
			Token: "",
		}, errors.New(fmt.Sprintf("LogIn: Empty token computed (%s)", email))
	}

	return entity.LoginResult{
		Token: token,
	}, nil
}
