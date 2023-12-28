saladcloud.json:
	curl https://docs.salad.com/reference/saladcloud-openapi-spec?json=on | jq .doc.body -r | head --lines=-1 | tail --lines=+2

saladcloud.yaml: saladcloud.json patch.sh
	cat $< | go run github.com/brancz/gojsontoyaml@v0.1.0 | sh patch.sh > $@

$(SPEAKEASY_FILES) files.gen &: saladcloud.yaml $(SPEAKEASY_BINARY)
	$(SPEAKEASY_BINARY) generate sdk --lang terraform --schema $< --out .
	$(MAKE) fmt
	$(MAKE) deps
	grep "some_saladcloud_api_key" --files-without-match $$(grep "# Configuration options" --recursive --exclude-dir bin --files-with-matches .) | xargs sed -i '/^\s*# Configuration options/a \ \ api_key = "some_saladcloud_api_key"'
