#!/usr/bin/env zsh

#  \
#  \\,
#   \\\,^,.,,.                     Zero to Hero
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,


IMAGE=$ECR_IMAGE_FIZZ_IDM
TAG=$ECR_TAG_FIZZ_IDM
REPO=$ECR_REPO 

echo "»»» building"
docker build -t "$IMAGE":"$TAG" .
retVal=$?
if [ $retVal -ne 0 ]; then
  echo "Error building the image."
  exit 1
fi

echo "»»» tagging"
docker tag "$IMAGE":"$TAG" "$REPO"/"$IMAGE":"$TAG"
retVal=$?
if [ $retVal -ne 0 ]; then
  echo "Error tagging the image."
  exit 1
fi

echo "»»» pushing"
docker push "$REPO"/"$IMAGE":"$TAG"
retVal=$?
if [ $retVal -ne 0 ]; then
  echo "Error pushing the image."
  exit 1
fi

echo "»»» Everything is awesome! «««"
