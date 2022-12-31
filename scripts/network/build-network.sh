#!/bin/bash

ROOT=$(pwd)
mkdir build

install_fa_binary() {
   cd "fa-chain"
   GOBIN=$ROOT/build/binary go install -mod=readonly -trimpath ./...
   local_build_succeeded=${PIPESTATUS[0]}

   if [[ "$local_build_succeeded" == "0" ]]; then
      echo "Done" 
   else
      echo "Failed"
   fi

   cd ..

   return $local_build_succeeded
}

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

install_fa_binary
# install_binary juno https://github.com/CosmosContracts/juno.git 5875239f4e2d20646024a2c7f4a383fa45081e81
install_binary osmosis https://github.com/osmosis-labs/osmosis.git v13.1.2
install_binary relayer https://github.com/cosmos/relayer.git andrew/client_icq

build_docker fachain nil
# build_docker juno 5875239f4e2d20646024a2c7f4a383fa45081e81
build_docker osmosis v13.1.2
build_docker relayer andrew/client_icq