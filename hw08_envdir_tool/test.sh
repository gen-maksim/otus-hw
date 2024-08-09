#!/usr/bin/env bash
set -xeuo pipefail

go build -o go-envdir

export HELLO="SHOULD_REPLACE"
export FOO="SHOULD_REPLACE"
export UNSET="SHOULD_REMOVE"
export ADDED="from original env"
export EMPTY="SHOULD_BE_EMPTY"

result=$(./go-envdir "C:/Users/loyal/GoProjects/otus-hw/hw08_envdir_tool/testdata/env" "C:/Program Files/Git/bin/bash" "C:/Users/loyal/GoProjects/otus-hw/hw08_envdir_tool/testdata/echo.sh" arg1=1 arg2=2)
expected='HELLO is ("hello")
BAR is (bar)
FOO is (   foo
with new line)
UNSET is ()
ADDED is (from original env)
EMPTY is ()
arguments are arg1=1 arg2=2'

[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}" && exit 1)

rm -f go-envdir
echo "PASS"
