#!/bin/bash
set -eux

NAVIGATOR_NAMESPACE="navigator"
USER_NAMESPACE="navigator-e2e-database1"
RELEASE_NAME="nav-e2e"

ROOT_DIR="$(git rev-parse --show-toplevel)"
SCRIPT_DIR="$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"
CONFIG_DIR=$(mktemp -d -t navigator-e2e.XXXXXXXXX)
mkdir -p $CONFIG_DIR
CERT_DIR="$CONFIG_DIR/certs"
mkdir -p $CERT_DIR
TEST_DIR="$CONFIG_DIR/tmp"
mkdir -p $TEST_DIR

source "${SCRIPT_DIR}/libe2e.sh"

kube_delete_namespace_and_wait "${USER_NAMESPACE}"

function debug_navigator_start() {
    kubectl api-versions
    kubectl get pods --all-namespaces
    kubectl describe deploy
    kubectl describe pod
}

function helm_install() {
    if [ "${CHART_VALUES}" == "" ]; then
        echo "CHART_VALUES must be set";
        exit 1
    fi
    helm delete --purge "${RELEASE_NAME}" || true
    echo "Installing navigator..."
    if helm --debug install --wait --name "${RELEASE_NAME}" contrib/charts/navigator \
         --values ${CHART_VALUES}
    then
        return 0
    fi
    return 1
}

# Retry helm install to work around intermittent API server availability.
# See https://github.com/jetstack/navigator/issues/118
if ! retry helm_install; then
    debug_navigator_start
    echo "ERROR: Failed to install Navigator"
    exit 1
fi

# Wait for navigator API to be ready
function navigator_ready() {
    local replica_count_controller=$(
        kubectl get deployment ${RELEASE_NAME}-navigator-controller \
                --output 'jsonpath={.status.readyReplicas}' || true)
    if [[ "${replica_count_controller}" -eq 0 ]]; then
        return 1
    fi
    local replica_count_apiserver=$(
        kubectl get deployment ${RELEASE_NAME}-navigator-apiserver \
                --output 'jsonpath={.status.readyReplicas}' || true)
    if [[ "${replica_count_apiserver}" -eq 0 ]]; then
        return 1
    fi
    if ! kubectl api-versions | grep 'navigator.jetstack.io'; then
        return 1
    fi
    # Even after the API appears in api-versions, it takes a short time for API
    # server to recognise navigator API types.
    if ! kubectl get esc; then
        return 1
    fi
    if ! kube_event_exists "kube-system" \
         "navigator-controller:Endpoints:Normal:LeaderElection"
    then
        return 1
    fi
    return 0
}

echo "Waiting for Navigator to be ready..."
if ! retry navigator_ready; then
    debug_navigator_start
    echo "ERROR: Timeout waiting for Navigator API"
    exit 1
fi

kubectl create namespace "${USER_NAMESPACE}"

FAILURE_COUNT=0

function fail_test() {
    FAILURE_COUNT=$(($FAILURE_COUNT+1))
    echo "TEST FAILURE: $1"
}

function test_elasticsearchcluster() {
    echo "Testing ElasticsearchCluster"
    if ! kubectl get esc; then
        fail_test "Failed to use shortname to get ElasticsearchClusters"
    fi
    # Create and delete an ElasticSearchCluster
    if ! kubectl create \
            --namespace "${USER_NAMESPACE}" \
            --filename "${SCRIPT_DIR}/testdata/es-cluster-test.yaml"; then
        fail_test "Failed to create elasticsearchcluster"
    fi
    if ! kubectl get \
            --namespace "${USER_NAMESPACE}" \
            ElasticSearchClusters; then
        fail_test "Failed to get elasticsearchclusters"
    fi
    if ! retry kubectl get \
         --namespace "${USER_NAMESPACE}" \
         service es-test; then
        fail_test "Navigator controller failed to create elasticsearchcluster service"
    fi
    if ! retry kube_event_exists "${USER_NAMESPACE}" \
         "navigator-controller:ElasticsearchCluster:Normal:SuccessSync"
    then
        fail_test "Navigator controller failed to create SuccessSync event"
    fi
    if ! retry TIMEOUT=300 stdout_gt 0 kubectl \
         --namespace "${USER_NAMESPACE}" \
         get pilot \
         "es-test-mixed-0" \
         "-o=go-template={{.status.elasticsearch.documents}}"
    then
        fail_test "Elasticsearch pilot did not update the document count"
    fi
    if ! kubectl delete \
            --namespace "${USER_NAMESPACE}" \
            ElasticSearchClusters \
            --all; then
        fail_test "Failed to delete elasticsearchcluster"
    fi
    if ! kube_delete_namespace_and_wait "${USER_NAMESPACE}"; then
        fail_test "Failed to delete test namespace"
    fi
}

test_elasticsearchcluster

function ignore_expected_controller_errors() {
    # Ignore the following error types:
    # E1103 14:58:06.819858       1 reflector.go:205] github.com/jetstack/navigator/pkg/client/informers/externalversions/factory.go:68: Failed to list *v1alpha1.Pilot: the server could not find the requested resource (get pilots.navigator.jetstack.io)
    # E1108 14:18:37.610718       1 reflector.go:205] github.com/jetstack/navigator/pkg/client/informers/externalversions/factory.go:68: Failed to list *v1alpha1.Pilot: an error on the server ("Error: 'dial tcp 10.0.0.233:443: getsockopt: connection refused'\nTrying to reach: 'https://10.0.0.233:443/apis/navigator.jetstack.io/v1alpha1/pilots?resourceVersion=0'") has prevented the request from succeeding (get pilots.navigator.jetstack.io)
    egrep --invert-match \
          -e 'Failed to list \*v1alpha1\.\w+:\s+the server could not find the requested resource\s+\(get \w+\.navigator\.jetstack\.io\)$' \
          -e 'Failed to list \*v1alpha1\.\w+:\s+an error on the server \([^)]+\) has prevented the request from succeeding\s+\(get \w+\.navigator\.jetstack\.io\)$'
}

function test_logged_errors() {
    if kubectl logs -c controller -l app=navigator,component=controller \
            | egrep '^E[0-9]{4} ' \
            | ignore_expected_controller_errors
    then
        fail_test "Unexpected errors in controller logs"
    fi
    if kubectl logs -c apiserver -l app=navigator,component=apiserver \
            | egrep '^E[0-9]{4} '
    then
        fail_test "Unexpected errors in apiserver logs"
    fi
}

test_logged_errors

# TODO Run integration tests in a separate script
function test_integration() {
    if ! go test \
         -timeout 60s \
         -v \
         ./test/integration/... \
         -args \
         -v 6 \
         -alsologtostderr \
         -logtostderr
    then
        fail_test "Integration test failure"
    fi
}

test_integration

kubectl api-versions
kubectl get apiservice -o yaml
kubectl cluster-info dump --all-namespaces || true

exit $FAILURE_COUNT
