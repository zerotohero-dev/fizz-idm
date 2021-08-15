#!/usr/bin/env zsh

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

