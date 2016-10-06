#!/bin/bash

set -ex

HERE=$(cd $(dirname $0); pwd)

mkdir -p ${HERE}/layers/tfdemo/bin

GOBIN=${HERE}/layers/tfdemo/bin go install ./cmd/tfwebapp

JUJU_REPOSITORY=${HERE} LAYER_PATH=${HERE}/layers INTERFACE_PATH=${HERE}/interfaces \
	charm build ${HERE}/layers/tfdemo

