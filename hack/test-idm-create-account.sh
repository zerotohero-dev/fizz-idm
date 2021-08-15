#!/usr/bin/env zsh

#  \
#  \\,
#   \\\,^,.,,.                     Zero to Hero
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,

IDM_CREATE_URL="http://localhost:9002/idm/v1/create"

curl --request POST \
  --url $IDM_CREATE_URL \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Mr Potato Head",
	"email": "potato3@volkan.io",
	"token": "potato",
	"password": "lorem-ipsum-dolar-sid-ahmed",
	"optIn": true
}'

