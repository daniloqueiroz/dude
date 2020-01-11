#!/usr/bin/env bash

password=$1
pass show "$password" | { IFS= read -r pass; printf %s "$pass"; } | xdotool type --clearmodifiers --file -