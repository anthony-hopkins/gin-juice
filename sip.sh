#!/bin/bash

# This script handles setting up the project directory and providing the necessary pieces to get started building
# a web application in GO and the Gin framework.

# ARGS check for sanity.
if [[ "$#" -ne 1 ]]
then
  echo "This script requires one argument - GO/Gin project name."
  echo "Exiting....."
  exit 1
fi

# GLOBALS
APPNAME="${1}"

# Zip extraction and cleanup
mkdir ./${APPNAME}
cp -a ./gin-juice-main/* ./${APPNAME}
