#!/bin/bash

OCTO_HOME=$HOME/.octo/



install(){

    if [ $1!="" ]
    then
        echo "No version selected, octo latest will be installed."
    fi

    mkdir -p $OCTO_HOME
}

install
