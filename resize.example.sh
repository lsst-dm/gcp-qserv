#!/bin/bash

set -euxo pipefail

PROJECT="qserv-dev-3d7e"
REGION="us-central1"
gcloud config set project $PROJECT

# Size=0
gcloud container clusters resize qserv-dev --quiet --node-pool czar-pool-b1d5  --region=$REGION --num-nodes 0
# Size=5
gcloud container clusters resize qserv-dev --quiet --node-pool worker-pool-ab14  --region=$REGION   --num-nodes 0
