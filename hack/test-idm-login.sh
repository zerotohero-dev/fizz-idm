#!/usr/bin/env zsh

#  \
#  \\,
#   \\\,^,.,,.                     Zero to Hero
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,

IDM_LOGIN_URL="http://localhost:9002/idm/v1/login"

curl --request POST \
  --url $IDM_LOGIN_URL \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "potato3@volkan.io",
	"password": "lorem-ipsum-dolar-sid-ahmed"
}'
