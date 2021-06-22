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

package api

import (
	"context"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-app/pkg/app"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-idm/internal/endpoint"
	"github.com/zerotohero-dev/fizz-idm/internal/service"
	"github.com/zerotohero-dev/fizz-idm/internal/transport"
)

func InitializeEndpoints(e env.FizzEnv, router *mux.Router) {
	svc := service.New(e, context.Background())

	// Gets user info for the logged-in user.
	app.Route(
		router, http.NewServer(
			endpoint.MakeInfoEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeInfoRequest),
			app.EncodeResponse,
		),
		"GET", "/v1/info",
	)

	// Authenticates the user.
	app.Route(
		router, http.NewServer(
			endpoint.MakeLoginEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeLoginRequest),
			app.EncodeResponse,
		),
		"POST", "/v1/login",
	)

	// Registers the user.
	app.Route(
		router, http.NewServer(
			endpoint.MakeSignupEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeSignupRequest),
			app.EncodeResponse,
		),
		"POST", "/v1/signup",
	)

	// Sends an email verification email to the user.
	app.Route(
		router, http.NewServer(
			endpoint.MakeVerifyEmailVerificationTokenEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeVerifyEmailVerificationTokenRequest),
			app.EncodeResponse,
		),
		"POST", "/v1/verify",
	)

	// Sends a “reset password” email to the user.
	app.Route(
		router, http.NewServer(
			endpoint.MakeSendPasswordResetTokenEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeSendPasswordResetTokenRequest),
			app.EncodeResponse,
		),
		"POST", "/v1/remind",
	)

	// Resets the user’s password if the token matches.
	app.Route(
		router, http.NewServer(
			endpoint.MakeResetPasswordEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeResetPasswordRequest),
			app.EncodeResponse,
		),
		"POST", "/v1/reset",
	)
}