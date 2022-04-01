#!/bin/bash

set -euxo pipefail

DIR=$(cd "$(dirname "$0")"; pwd -P)
. $DIR/env.sh

gcloud container clusters get-credentials qserv-dev --region="$REGION"
