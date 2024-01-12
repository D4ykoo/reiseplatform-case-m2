#!/bin/bash

# get command line arguments flags -g for generate and -d for delete
while getopts gd flag
do
    case "${flag}" in
        g) generate=true;;
        d) delete=true;;
    esac
done

NETWORKS=('dbnetcheckout' 'dbnetmonitor' 'dbnetusermanagement' 'dbnettravmgnt' 'privatnet' 'publicnet')

if [ "$generate" = true ] ; then
    echo "Generating networks..."
    for network in "${NETWORKS[@]}"
    do
        # true to ignore error if network already exists
        docker network create -d bridge $network || true
        echo "Docker network $network created!"
    done
elif [ "$delete" = true ] ; then
    echo "Deleting networks..."
    for network in "${NETWORKS[@]}"
    do
        docker network rm $network
        echo "Docker network $network removed!"
    done
else
    echo "No flag specified. Use -g to generate networks or -d to delete networks."
    exit 1
fi


