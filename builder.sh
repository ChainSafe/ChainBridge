#!/bin/bash

GIT_TOKEN=$1
GIT_KEY=$2
SPECFILE=$3
PROJECT_ID=$4
CLUSTER_NAME=$5
REPO_NAME=$6
SECRET_PROJECT_ID=$7
IMPERSONATE_ACCOUNT=$8

### Export GIT TOKEN ###
export GITHUB_OAUTH_TOKEN=$(gcloud secrets versions access latest --secret=$GIT_TOKEN --project $SECRET_PROJECT_ID|tr -d '')

### Update the SSH keys in the path ###
gcloud secrets versions access latest --secret=$GIT_KEY --project $SECRET_PROJECT_ID > /root/.ssh/github && chmod 600 /root/.ssh/github
eval "$(ssh-agent -s)"
ssh-add /root/.ssh/github

### Clear if repo exists ###
rm -rf $REPO_NAME

### Fix for HEX dependecy ###
mix local.hex --force && mix local.rebar

### Clone the repository ###
git clone --recursive git@github.com:fortelabsinc/$REPO_NAME.git
a=$?

if [ $a -eq 0 ]
 then
   cd $REPO_NAME && git submodule foreach git pull origin main
   a=$?
fi

if [ $a -eq 0 ]
 then
   gcloud config set project $PROJECT_ID
   gcloud config set auth/impersonate_service_account $IMPERSONATE_ACCOUNT
   gcloud auth configure-docker

   ### Application BUILD ###
   cd fc_devops && ./ops --build --spec $SPECFILE && ./ops --publish --spec $SPECFILE
   b=$?

   if [ $b -eq 0 ]
    then
     ### Application DEPLOYMENT ###
     gcloud container clusters get-credentials $CLUSTER_NAME --project $PROJECT_ID  --region us-central1
     ./ops --deploy --spec $SPECFILE
     c=$?

     if [ $c -eq 1 ]
      then
        echo "Deployment has failed"
        exit 1
     fi
   else
     echo "Build has failed"
     exit 1
   fi
else
  echo "Git clone has failed"
  exit 1
fi