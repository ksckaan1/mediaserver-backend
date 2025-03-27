#!/bin/sh

/couchbase-cli/couchbase-cli bucket-list -u ${COUCHBASE_USER} -p ${COUCHBASE_PASSWORD} -c couchbase

if [ $? -eq 0 ]; then
  echo "CLUSTER ALREADY INITIALIZED"
  exit 0
fi

/couchbase-cli/couchbase-cli cluster-init \
  --cluster-username ${COUCHBASE_USER} \
  --cluster-password ${COUCHBASE_PASSWORD} \
  -c ${COUCHBASE_HOST} \
  --services index,data,query \
  --cluster-name my-cluster

/couchbase-cli/couchbase-cli bucket-create \
  -u ${COUCHBASE_USER} \
  -p ${COUCHBASE_PASSWORD} \
  -c ${COUCHBASE_HOST} \
  --bucket ${COUCHBASE_BUCKET} \
  --bucket-type couchbase \
  --bucket-ramsize 1392 \
  --bucket-replica 1