#!/bin/bash
set -eo pipefail

#goals
#1. start a timer and stop when the pods are done deploying
#2. Be able to set the number of pods to deploy
#3. The pods use whereabouts

#should probably check that oc is installed here
#should check here that a kind cluster is running 

#start a timer to record how long the pods take to spin up
start=$SECONDS
HERE="$(dirname "$(readlink --canonicalize ${BASH_SOURCE[0]})")"
ROOT="$(readlink --canonicalize "$HERE/..")"
WHEREABOUTSNAD="$ROOT/yamls/whereaboutsNAD.yaml"
SCALEDEPLOYMENT="$ROOT/yamls/scale-deployment.yaml"

retry() {
  local status=0
  local retries=${RETRY_MAX:=5}
  local delay=${INTERVAL:=5}
  local to=${TIMEOUT:=20}
  cmd="$*"

  while [ $retries -gt 0 ]
  do
    status=0
    timeout $to bash -c "echo $cmd && $cmd" || status=$?
    if [ $status -eq 0 ]; then
      break;
    fi
    echo "Exit code: '$status'. Sleeping '$delay' seconds before retrying"
    sleep $delay
    let retries--
  done
  return $status
}


#create the whereabouts nad
oc apply -f "$WHEREABOUTSNAD"
#create the deployment 
oc apply -f "$SCALEDEPLOYMENT"
retry kubectl wait --for=condition=available deploy/scale-deployment
#wait for all pods to be deployed

#Log the amount of time it took the pods to create
duration=$(( SECONDS - start ))
echo Pod creation duration:"$duration"
echo Here:"$HERE"
echo Root:"$ROOT"
echo WhereaboutsNAD:"$WHEREABOUTSNAD"
echo deployment:"$SCALEDEPLOYMENT"