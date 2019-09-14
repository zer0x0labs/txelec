#!/bin/bash
ARGS="${@}"
clear;
while(true); do
    OUTPUT=`$ARGS`
    clear
    echo -e "${OUTPUT[@]}"
    sleep 5m
done