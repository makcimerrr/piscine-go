export HERO_ID=1
curl -s https://zone01normandie.org/assets/superhero/all.json | jq ' .[] | select( .id == 1)' | jq ' .connections' | jq ' .relatives'
