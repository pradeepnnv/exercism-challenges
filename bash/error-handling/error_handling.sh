#!/usr/bin/env bash

# The following comments should help you get started:
# - Bash is flexible. You may use functions or write a "raw" script.
#
# - Complex code can be made easier to read by breaking it up
#   into functions, however this is sometimes overkill in bash.
#
# - You can find links about good style and other resources
#   for Bash in './README.md'. It came with this exercise.
#
#   Example:
#   # other functions here
#   # ...
#   # ...
#
#   main () {
#     # your main function code here
#   }
#
#   # call main with all of the positional arguments
#   main "$@"
#
# *** PLEASE REMOVE THESE COMMENTS BEFORE SUBMITTING YOUR SOLUTION ***
#!/bin/bash
IFS='' # This is need to ensure bash does not split it's arguments if they have spaces. Refer to https://www.gnu.org/software/bash/manual/bash.html#Word-Splitting.

if [ $#  -eq 1 ]
then
    name="$1"
    printf "Hello, %s" $name
else
    echo "Usage: error_handling.sh <person>"
    exit 1
fi  