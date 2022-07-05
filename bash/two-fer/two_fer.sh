#!/bin/bash
IFS='' # This is need to ensure bash does not split it's arguments if they have spaces. Refer to https://www.gnu.org/software/bash/manual/bash.html#Word-Splitting.
name="$1"
if [ -z "$name" ]
then
  name='you'
fi
printf "One for %s, one for me." $name