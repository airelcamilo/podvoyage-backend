# Retrieve Cloud Run service URL -
# Cloud Run URLs are not deterministic.
get_url() {
    gcloud run services describe ${_SERVICE_NAME}-${BUILD_ID} \
        --format 'value(status.url)' \
        --region ${_DEPLOY_REGION} \
        --platform managed
}

# Retrieve Id token to make an aunthenticated request -
# Impersonate service account, token-creator@, since
# Cloud Build does not natively mint identity tokens.
get_idtoken() {
    curl -X POST -H "content-type: application/json" \
        -H "Authorization: Bearer $(gcloud auth print-access-token)" \
        -d "{\"audience\": \"$(cat _service_url)\"}" \
        "https://iamcredentials.googleapis.com/v1/projects/-/serviceAccounts/token-creator@${PROJECT_ID}.iam.gserviceaccount.com:generateIdToken" |
        python3 -c "import sys, json; print(json.load(sys.stdin)['token'])"
}