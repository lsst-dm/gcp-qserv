#!/bin/bash

set -euxo pipefail

DIR=$(cd "$(dirname "$0")"; pwd -P)
. $DIR/env.sh

# Size=0
gcloud container clusters resize qserv-dev --quiet --node-pool czar-pool-b1d5  --region=$REGION --num-nodes 0
# Size=5
gcloud container clusters resize qserv-dev --quiet --node-pool worker-pool-ab14  --region=$REGION   --num-nodes 0
