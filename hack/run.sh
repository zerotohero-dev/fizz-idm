#!/usr/bin/env zsh

#  \
#  \\,
#   \\\,^,.,,.                     Zero to Hero
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,

TAG="0.0.10"

# shellcheck disable=SC1090
source ~/.zprofile

docker run -e FIZZ_DEPLOYMENT_TYPE="$FIZZ_DEPLOYMENT_TYPE" \
-e FIZZ_IDM_SVC_PORT="$FIZZ_IDM_SVC_PORT" \
-e FIZZ_IDM_HONEYBADGER_API_KEY="$FIZZ_IDM_HONEYBADGER_API_KEY" \
-e FIZZ_IDM_USERS_TABLE_NAME="$FIZZ_IDM_USERS_TABLE_NAME" \
-e FIZZ_IDM_VERIFIED_URL="$FIZZ_IDM_VERIFIED_URL" \
-e FIZZ_IDM_CRYPTO_ENDPOINT_URL="$FIZZ_IDM_CRYPTO_ENDPOINT_URL" \
-e FIZZ_LOG_DESTINATION="$FIZZ_LOG_DESTINATION" \
zerotohero-dev/fizz-idm:$TAG
