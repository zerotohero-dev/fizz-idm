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
	"github.com/zerotohero-dev/fizz-entity/pkg/collection"
	"github.com/zerotohero-dev/fizz-entity/pkg/connection"
	entity "github.com/zerotohero-dev/fizz-entity/pkg/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func users() *mongo.Collection {
	return collection.Collection(dbName, usersTableName)
}

func findUserByEmail(ctx context.Context, email string) (*mongo.Cursor, error) {
	return users().Find(ctx, bson.D{{entity.Keys.Email, email}})
}

func userCursorByStatus(
	ctx context.Context, email, status string,
) (*mongo.Cursor, error) {
	return users().Find(ctx, bson.D{
		{entity.Keys.Email, email},
		{entity.Keys.Status, status},
	})
}

func userCursorByStatusAndEmailVerificationToken(
	ctx context.Context, email, status, token string,
) (*mongo.Cursor, error) {
	return users().Find(ctx, bson.D{
		{entity.Keys.Email, email},
		{entity.Keys.Status, status},
		{entity.Keys.EmailVerificationToken, token},
	})
}

func findUnverifiedUserByEmail(
	ctx context.Context, email string,
) (*mongo.Cursor, error) {
	return userCursorByStatus(ctx, email, entity.Status.Unverified)
}

func findVerifiedUserByEmail(
	ctx context.Context, email string,
) (*mongo.Cursor, error) {
	return userCursorByStatus(ctx, email, entity.Status.Verified)
}

func findUnverifiedUserByEmailAndEmailVerificationToken(
	ctx context.Context, email, token string,
) (*mongo.Cursor, error) {
	return userCursorByStatusAndEmailVerificationToken(
		ctx, email, entity.Status.Unverified, token,
	)
}

func UserExists(email string) (bool, error) {
	ctx, _ := connection.CreateDbContext()
	cur, err := findUserByEmail(ctx, email)
	defer func() {
		_ = connection.CloseCursor(cur, ctx)
	}()

	if err != nil {
		return false, errors.Wrap(err, "UserExists: error creating cursor")
	}

	for cur.Next(ctx) {
		return true, nil
	}

	return false, nil
}

func UnverifiedUserByEmailAndEmailVerificationToken(email, token string) (*entity.User, error) {
	ctx, _ := connection.CreateDbContext()
	cur, err := findUnverifiedUserByEmailAndEmailVerificationToken(ctx, email, token)
	defer func() {
		_ = connection.CloseCursor(cur, ctx)
	}()

	if err != nil {
		return nil, errors.Wrap(err, "UnverifiedUserByEmailAndEmailVerificationToken: error creating cursor")
	}

	for cur.Next(ctx) {
		var u entity.User

		err = cur.Decode(&u)

		if err != nil {
			return nil, errors.Wrap(err, "UnverifiedUserByEmailAndEmailVerificationToken: error decoding user")
		}

		return &u, nil
	}

	return nil, nil
}

func UnverifiedUserByEmail(email string) (*entity.User, error) {
	ctx, _ := connection.CreateDbContext()
	cur, err := findUnverifiedUserByEmail(ctx, email)
	defer func() {
		_ = connection.CloseCursor(cur, ctx)
	}()

	if err != nil {
		return nil, errors.Wrap(err, "UnverifiedUserByEmail: error creating cursor")
	}

	for cur.Next(ctx) {
		var u entity.User

		err = cur.Decode(&u)

		if err != nil {
			return nil, errors.Wrap(err, "UnverifiedUserByEmail: error decoding user")
		}

		return &u, nil
	}

	return nil, nil
}

func VerifiedUserByEmail(email string) (*entity.User, error) {
	ctx, _ := connection.CreateDbContext()
	cur, err := findVerifiedUserByEmail(ctx, email)
	defer func() {
		_ = connection.CloseCursor(cur, ctx)
	}()

	if err != nil {
		return nil, errors.Wrap(err, "VerifiedUserByEmail: error creating cursor")
	}

	for cur.Next(ctx) {
		var u entity.User

		err = cur.Decode(&u)

		if err != nil {
			return nil, errors.Wrap(err, "VerifiedUserByEmail: error decoding user")
		}

		return &u, nil
	}

	return nil, nil
}
