#!/bin/bash

set -euxo pipefail

gcloud container clusters list

CLUSTER_NAME="qserv-dev"

gcloud container clusters describe "$CLUSTER_NAME" --region=us-central1 | grep masterIpv4CidrBlock

gcloud compute firewall-rules list \
    --filter "name~^gke-${CLUSTER_NAME}" \
    --format 'table(
        name,
        network,
        direction,
        sourceRanges.list():label=SRC_RANGES,
        allowed[].map().firewall_rule().list():label=ALLOW,
        targetTags.list():label=TARGET_TAGS
    )'