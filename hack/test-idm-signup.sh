#!/usr/bin/env zsh

#  \
#  \\,
#   \\\,^,.,,.                     Zero to Hero
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,

IDM_CREATE_URL="http://localhost:9002/idm/v1/signup"

curl --request POST \
  --url $IDM_CREATE_URL \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Mr Potato Head",
	"email": "potato4@volkan.io"
}'

