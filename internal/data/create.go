package data

import (
	"github.com/pkg/errors"
	"github.com/zerotohero-dev/fizz-entity/pkg/connection"
	entity "github.com/zerotohero-dev/fizz-entity/pkg/data"
	"time"
)

// CreateUnverifiedUser creates an unverified user.
// It assumes that `u` is well-formed and sanitized.
// It does not alter the input user `u`, and it passes it to the DB as is.
func CreateUnverifiedUser(u entity.User) error {
	u.Status = entity.Status.Unverified
	u.Token.EmailVerificationToken = ""

	now := time.Now()
	u.RecordCreated = now
	u.RecordUpdated = now
	u.Password = ""

	ctx, _ := connection.CreateDbContext()
	_, err := users().InsertOne(ctx, u)

	if err != nil {
		return errors.Wrap(err, "CreateUnverifiedUser: unable to create user")
	}

	return nil
}
