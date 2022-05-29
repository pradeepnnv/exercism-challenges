#!/usr/bin/env bash
echo $@
name="$1"

if [ -z "$name" ]
then
  name='you'
fi
printf "One for %s, one for me." $name