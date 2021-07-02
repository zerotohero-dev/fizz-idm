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
	now := (time.Now().UnixNano()) / 1000000

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
