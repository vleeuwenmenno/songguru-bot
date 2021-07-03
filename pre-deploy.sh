#!/bin/bash

echo "Fetching binaries to build/"
cp deploy.sh ./build/deploy.sh

echo "Removing debug symbols ..."
rm *.pdb

echo "Build folder preview:"
ls -al