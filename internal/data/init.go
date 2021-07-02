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
	"github.com/zerotohero-dev/fizz-entity/pkg/connection"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
)

var dbName = ""
var usersTableName = ""

func Init(e env.FizzEnv) {
	connection.Connect(e.Idm.DbConnectionString)
	dbName = e.Idm.DbName
	usersTableName = e.Idm.UsersTableName
}
