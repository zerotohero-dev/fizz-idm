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

package data

import (
	"github.com/pkg/errors"
	"github.com/zerotohero-dev/fizz-entity/pkg/connection"
	entity "github.com/zerotohero-dev/fizz-entity/pkg/data"
	"time"
)

// CreateUnverifiedUser creates an unverified user.
// It assumes that `u` is well-formed and sanitized.
// It does not sanitize `u`. The sanitization ought to have
// happened at the endpoint layer, as soon as we get the parameters
// from the request.
func CreateUnverifiedUser(u entity.User) error {
	u.Status = entity.Status.Unverified

	now := time.Now()
	u.RecordCreated = now
	u.RecordUpdated = now
	u.Password = ""
	u.StripeSubscription = nil

	ctx, _ := connection.CreateDbContext()
	_, err := users().InsertOne(ctx, u)

	if err != nil {
		return errors.Wrap(err, "CreateUnverifiedUser: unable to create user")
	}

	return nil
}
