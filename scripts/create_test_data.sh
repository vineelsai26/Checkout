#!/bin/bash

rand_str() {
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

i=0
iend=100

while [ $i -le $iend ]; do
    echo "Creating dir$i"
    mkdir -p dir$i
    cd dir$i

    j=0
    jend=100
    while [ $j -le $jend ]; do
        echo "Creating file$j in dir$i"
        touch file$j
        echo $(rand_str 1000) >file$j
        j=$(($j + 1))
    done
    cd ..
    i=$(($i + 1))
done

ls -laR
cd ..

touch testfile
echo $(rand_str 10000) >testfile
