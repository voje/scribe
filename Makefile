svc-acc.json:
	gpg --decrypt svc-acc.json.gpg

# SCRIBE_SVCACC_JSON_GPG_PASS="test-gpg-pass"
run:
	SCRIBE_SVCACC_JSON_GPG="./svc-acc.json.gpg" \
	go run cmd/scribe/scribe.go sync

install:
	go install cmd/scribe/scribe.go
	cp systemd/* ~/.config/systemd/user/.
	systemctl --user daemon-reload
	systemctl --user enable scribe.timer
	-systemctl --user start scribe.service
	systemctl --user status scribe.service
