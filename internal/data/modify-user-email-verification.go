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
	"context"
	"github.com/pkg/errors"
	"github.com/zerotohero-dev/fizz-entity/pkg/connection"
	entity "github.com/zerotohero-dev/fizz-entity/pkg/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func updateUnverifiedUserEmailVerificationToken(
	ctx context.Context, email, emailVerificationToken string,
) (*mongo.UpdateResult, error) {
	now := time.Now()

	return users().UpdateOne(ctx, bson.D{
		{entity.Keys.Email, email},
		{entity.Keys.Status, entity.Status.Unverified},
	}, bson.D{{
		"$set", bson.D{
			{entity.Keys.EmailVerificationToken, emailVerificationToken},
			{entity.Keys.RecordUpdated, now},
		},
	}})
}

// verifyUnverifiedUserIfEmailVerificationTokenMatches sets the status of the
// user to `verified` if the user has a valid email verification token and
// also the user’s current status is `unverified` right now.
// If the user is already verified this function has no effect.
func verifyUnverifiedUserIfEmailVerificationTokenMatches(
	ctx context.Context, email, name, token string, optIn bool, password string,
) (*mongo.UpdateResult, error) {
	now := time.Now()

	return users().UpdateOne(ctx, bson.D{
		{entity.Keys.Email, email},
		{entity.Keys.Status, entity.Status.Unverified},
		{entity.Keys.EmailVerificationToken, token},
	}, bson.D{{
		"$set", bson.D{
			{entity.Keys.Name, name},
			{entity.Keys.Status, entity.Status.Verified},
			{entity.Keys.EmailVerificationToken, ""},
			{entity.Keys.SubscribedToMailingList, optIn},
			{entity.Keys.Password, password},
			{entity.Keys.RecordUpdated, now},
		},
	}})
}

// TODO: delete this if we are not using it.
func UpdateUnverifiedUserEmailVerificationToken(email, accountActivationToken string) error {
	ctx, _ := connection.CreateDbContext()
	_, err := updateUnverifiedUserEmailVerificationToken(
		ctx, email, accountActivationToken,
	)

	if err != nil {
		return errors.Wrap(err, "UpdateUnverifiedUserEmailVerificationToken: error updating data")
	}

	return nil
}

func SetUserVerified(user entity.User) error {
	ctx, _ := connection.CreateDbContext()
	_, err := verifyUnverifiedUserIfEmailVerificationTokenMatches(
		ctx, user.Email, user.Name, user.EmailVerificationToken,
		user.SubscribedToMailingList, user.Password,
	)

	if err != nil {
		return errors.Wrap(err, "UpdateUnverifiedUserEmailVerificationToken: error updating data")
	}

	return nil
}
