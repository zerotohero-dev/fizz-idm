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

	prefix := e.Idm.PathPrefix

	// Gets user info for the logged-in user.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeInfoEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeInfoRequest),
			app.EncodeResponse,
		),
		router, "GET", prefix, "/v1/info",
	)

	// Authenticates the user.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeLoginEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeLoginRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/login",
	)

	// Sends and email verification email to the user.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeSignupEndpoint(svc),
			app.ContentTypeValidatingMiddleware(transport.DecodeSignupRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/signup",
	)

	// Creates the user’s account (needs email verification token)
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeCreateAccountEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeCreateAccountRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/create",
	)

	// Sends a “reset password” email to the user.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeSendPasswordResetTokenEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeSendPasswordResetTokenRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/remind",
	)

	// Resets the user’s password if the token matches.
	app.RoutePrefixedPath(
		http.NewServer(
			endpoint.MakeResetPasswordEndpoint(svc),
			app.ContentTypeValidatingMiddleware(
				transport.DecodeResetPasswordRequest),
			app.EncodeResponse,
		),
		router, "POST", prefix, "/v1/reset",
	)
}
