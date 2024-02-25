svc-acc.json:
	gpg --decrypt svc-acc.json.gpg

# run: svc-acc.json
# SCRIBE_SVCACC_JSON_GPG_PASS="test-gpg-pass"
# SCRIBE_SMTP_PASS="test-smtp-pass"
run:
	SCRIBE_SMTP_PASS="test-smtp-pass" \
	SCRIBE_SVCACC_JSON_GPG="./svc-acc.json.gpg" \
	go run cmd/scribe/scribe.go list

