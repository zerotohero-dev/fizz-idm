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

import "github.com/zerotohero-dev/fizz-entity/pkg/data"

func (s service) CreateAccount(user data.User) (*data.User, error) {
	// Verify that the user exists in the database.
	// Verify that the user is still unverified.
	// Verify that the token this user has matches the token in the database.
	// Update the user’s status.
	/// panic("implement me")

	return nil, nil
}
