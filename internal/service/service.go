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

/*
	Things to do:
	  1. IDM is not building for some reason: Fix it.
	  2. Containerize mailer µService.
	  3. Docker and k8s consume CPU… I need CPU for the streams.
	     and since I cannot buy a beefier machine for now, I’ll
		 use k8s as a service for now.
	  0. I need to test this new setup too.
*/

import (
	"context"

	"github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

type Service interface {
	Info(authToken string) (data.UserInfo, error)
	LogIn(email, password string) (data.LoginResult, error)
	SignUp(user data.User) error
	CreateAccount(user data.User) error
	SendPasswordResetToken(email string) error
	ResetPassword(email, password, passwordResetToken string) error
	Context() context.Context
}

type service struct {
	env env.FizzEnv
	ctx context.Context
}

func (s service) Context() context.Context {
	return s.ctx
}

func New(e env.FizzEnv, ctx context.Context) Service {
	return &service{
		env: e,
		ctx: ctx,
	}
}
