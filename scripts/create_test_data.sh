#!/bin/bash

rand_str () {
    # Return random alpha-numeric string of given LENGTH
    #
    # Usage: VALUE=$(rand_str $LENGTH)
    #    or: VALUE=$(rand_str)

    local DEFAULT_LENGTH=64
    local LENGTH=${1:-$DEFAULT_LENGTH}

    LC_ALL=C tr -dc A-Za-z0-9 < /dev/urandom | head -c $LENGTH
    # LC_ALL=C: required for Mac OS X - https://unix.stackexchange.com/a/363194/403075
    # -dc: delete complementary set == delete all except given set
}

cd tests
mkdir -p testdir
cd testdir

for i in {1..10000} ; do
    echo "Creating dir$i"
    mkdir -p dir$i
    cd dir$i
    for j in {1..1000} ; do
        touch file$j
        echo `rand_str 100000` > file$j
    done
    cd ..
done

ls -laR
cd ..

touch testfile
echo `rand_str 1000000` > testfile
