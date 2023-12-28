saladcloud.json:
	curl https://docs.salad.com/reference/saladcloud-openapi-spec?json=on | jq .doc.body -r | head --lines=-1 | tail --lines=+2

saladcloud.yaml: saladcloud.json patch.sh
	cat $< | go run github.com/brancz/gojsontoyaml@v0.1.0 | sh patch.sh > $@
