#!/bin/bash

rand_str () {
    # Return random alpha-numeric string of given LENGTH
    #
    # Usage: VALUE=$(rand_str $LENGTH)
    #    or: VALUE=$(rand_str)

    DEFAULT_LENGTH=64
    LENGTH=${1:-$DEFAULT_LENGTH}

    openssl rand -hex $LENGTH | cut -c1-$LENGTH
}

cd tests
mkdir -p testdir
cd testdir

for i in {1..10000} ; do
    date
    echo "Creating dir$i"
    date
    mkdir -p dir$i
    date
    cd dir$i
    date
    for j in {1..1000} ; do
        date
        touch file$j
        echo `rand_str 100000` > file$j
        date
    done
    cd ..
    date
done

ls -laR
cd ..

touch testfile
echo `rand_str 1000000` > testfile
