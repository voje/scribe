svc-acc.json:
	gpg --decrypt svc-acc.json.gpg

run: svc-acc.json
	GOOGLE_APPLICATION_CREDENTIALS="./svc-acc.json" \
	go run cmd/scribe/scribe.go

