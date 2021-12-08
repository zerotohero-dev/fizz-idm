```text
 \
 \\,
  \\\,^,.,,.                     Zero to Hero
  ,;7~((\))`;;,,               <zerotohero.dev>
  ,(@') ;)`))\;;',    stay up to date, be curious: learn
   )  . ),((  ))\;,
  /;`,,/7),)) )) )\,,
 (& )`   (,((,((;( ))\,
```

# Runbook to Test IDM Functionality

1. `./test-idm-signup.sh` => Create an unverified account.
2. `./test-idm-create-account.sh` => Verify email and set the user’s password.
3. `./test-idm-login.sh` => log user in and get a cookie as response.
