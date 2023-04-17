#! /bin/bash

curl -s https://api.github.com/users/mdubois | jq ' .id'
