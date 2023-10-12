#!/bin/sh

if test -n "$1"; then
    # need -R not -r to copy hidden files
    cp -R "$1/.verifier" /root
fi

mkdir -p /root/log
verifierd start --api.address tcp://0.0.0.0:1317 --rpc.laddr tcp://0.0.0.0:26657 --trace
