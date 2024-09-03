#!/bin/bash

output_file="oracleX.proto"
rm -rf $output_file

cd service
header_file="header.proto"
cat $header_file >> ../$output_file

for file in *.proto; do
    if [ "$file" != "header.proto" ]; then
        echo "// ------------------ $file -------------------" >> ../$output_file
        grep -vE '^(import|package|syntax|option)' $file >> ../$output_file
        echo '' >> ../$output_file
        echo '' >> ../$output_file
    fi
done
