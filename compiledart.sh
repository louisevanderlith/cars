#!/bin/bash

ls -1 assets/dart/*.entry.dart | while read -r line; do 
    dest=$(echo $line | sed -e "s/dart/js/g");
    dart2js -o $dest $line
done