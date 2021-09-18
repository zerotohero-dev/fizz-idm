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
	"context"
	"fmt"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-idm/internal/downstream"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
)


func sendEmailVerificationToken(name, email string, emailVerificationToken string) {
	go func() {

		res, err := downstream.Endpoints().MailerVerification(

			// We are using context.Background() because we do not want this to
			// cancel prematurely. go-kit cancels the owner context as soon as
			// the function exits.
			context.Background(), reqres.RelayEmailVerificationMessageRequest{
				Email: email,
				Name:  name,
				Token: emailVerificationToken,
			})

		if err != nil {
			log.Err("Problem sending activation email (%s) (%s)",
				log.RedactEmail(email), err.Error())
		}

		er := res.(reqres.RelayEmailVerificationMessageResponse)
		if er.Err != "" {
			log.Err("Problem sending activation email (%s) (%s)",
				log.RedactEmail(email), er.Err)
		}
	}()
}

func sendWaitlistEmail(name, email string) {
	go func() {
		res, err := downstream.Endpoints().MailerVerification(

			// We are using context.Background() because we do not want this to
			// cancel prematurely. go-kit cancels the owner context as soon as
			// the function exits.
			context.Background(), reqres.RelayWelcomeMessageRequest{
				Email: email,
				Name:  name,
			})

		fmt.Println(res)

		if err != nil {
			log.Err("Problem sending Waitlist email (%s) (%s)",
				log.RedactEmail(email), err.Error())
		}

		er := res.(reqres.RelayWelcomeMessageResponse)
		if er.Err != "" {
			log.Err("Problem sending Waitlist email (%s) (%s)",
				log.RedactEmail(email), er.Err)
		}
	}()
}