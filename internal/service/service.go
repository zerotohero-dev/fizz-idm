/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package service

import (
	"context"
	"github.com/zerotohero-dev/fizz-entity/pkg/data"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

type Service interface {
	SignUp(user data.User) error
	LogIn(email, password string) (data.LoginResult, error)
	VerifyEmailVerificationToken(email, emailVerificationToken string) (data.User, error)
	Info(authToken string) (data.Info, error)
	SendPasswordResetToken(email string) error
	ResetUserPassword(email, password, passwordResetToken string) error
}

type service struct {
	env env.FizzEnv
	ctx context.Context
}

func New(e env.FizzEnv, ctx context.Context) Service {
	return &service{
		env: e,
		ctx: ctx,
	}
}
