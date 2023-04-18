#! /bin/bash

curl -s "https://zone01normandie.org/assets/superhero/all.json" | jq -r --arg HERO_ID "$HERO_ID" '.[] | select(.id == ($HERO_ID|tonumber)) | .connections .relatives'
