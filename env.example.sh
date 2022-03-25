
# gcloud projects list
PROJECT="qserv-dev-3d7e"
#PROJECT="qserv-int-8069"
REGION="us-central1"
ZONE="${REGION}-c"
CLUSTER="qserv-dev"
gcloud config set project $PROJECT
gcloud config set compute/zone $ZONE
gcloud config get-value project
