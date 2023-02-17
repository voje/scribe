svc-acc.json:
	gpg --decrypt svc-acc.json.gpg

# run: svc-acc.json
run:
	SCRIBE_SVCACC_JSON_GPG_PASS="test-gpg-pass" \
	SCRIBE_SMTP_PASS="test-smtp-pass" \
	go run cmd/scribe/scribe.go \
		--svcacc-json-gpg ./svc-acc.json.gpg

