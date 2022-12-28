#!/bin/bash

ROOT=$(pwd)
mkdir build

install_binary() {
   title=$1
   download=$2
   branch=$3

   cd build
   if [ ! -d $title ]; then
      git clone $download
   fi
   
   cd $title
   git checkout $branch
   GOBIN=$ROOT/build/binary go install -mod=readonly -trimpath ./...
   local_build_succeeded=${PIPESTATUS[0]}

   if [[ "$local_build_succeeded" == "0" ]]; then
      echo "Done" 
   else
      echo "Failed"
   fi

   cd ../..

   return $local_build_succeeded
}

build_docker() {
   title=$1
   branch=$2

   echo "Building $title Docker...  "
   DOCKER_BUILDKIT=1 docker build --build-arg COMMIT_HASH=$branch --tag fee-abstraction:$title -f scripts/network/Dockerfile.$title . | true
   # PIPESTATUS tracks result code of current command
   docker_build_succeeded=${PIPESTATUS[0]}

   if [[ "$docker_build_succeeded" == "0" ]]; then
      echo "Done" 
   else
      echo "Failed"
   fi
   return $docker_build_succeeded
}

install_binary juno https://github.com/Smartosc-Blockchain/juno.git 67dbd631074aab748d8f7b8e080d11ddc406d1cc

install_binary osmosis https://github.com/osmosis-labs/osmosis.git v13.1.2

install_binary relayer https://github.com/cosmos/relayer.git andrew/client_icq

build_docker juno 67dbd631074aab748d8f7b8e080d11ddc406d1cc
build_docker osmosis v13.1.2
build_docker relayer andrew/client_icq