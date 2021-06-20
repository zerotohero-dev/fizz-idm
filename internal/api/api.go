/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-idm/internal/service"
)

func InitializeEndpoints(e env.FizzEnv, router *mux.Router) {
	svc := service.New(e, context.Background())

	fmt.Println(svc, router)

	panic("Implement me!")
}