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

package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-idm/internal/service"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"github.com/zerotohero-dev/fizz-validation/pkg/sanitization"
)

func MakeInfoEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		gr, hasContentTypeProblem := request.(reqres.ContentTypeProblemRequest)
		if hasContentTypeProblem {
			return reqres.UserInfoResponse{
				Err: gr.Err,
			}, nil
		}

		req := request.(reqres.UserInfoRequest)
		if req.Err != "" {
			return reqres.UserInfoResponse{
				Err: req.Err,
			}, nil
		}

		authToken := sanitization.SanitizeToken(req.AuthToken)
		if authToken == "" {
			return reqres.UserInfoResponse{
				Err: "MakeInfoEndpoint: missing auth token",
			}, nil
		}

		infoResult, err := svc.Info(authToken)
		if err != nil {
			log.Info("MakeInfoEndpoint:err: (%s)", err.Error())

			return reqres.UserInfoResponse{
				Err: "MakeInfoEndpoint: error checking user activation",
			}, nil
		}

		subscriptionId := ""
		if infoResult.StripeSubscription != nil {
			subscriptionId = infoResult.StripeSubscription.StripeProductId
		}

		return reqres.UserInfoResponse{
			Email:         infoResult.Email,
			Subscription:  subscriptionId,
			Name:          infoResult.Name,
		}, nil
	}
}
