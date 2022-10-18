gcf_region = asia-northeast1
gcf_function_name = kamikou-twitter-notification
gcf_entry_point = HelloWorld

run:
	cd gcf && go run local/main.go

deploy:
	cd gcf && \
	gcloud functions deploy $(gcf_function_name) \
		--gen2 \
		--runtime=go116 \
		--region=$(gcf_region) \
		--source=. \
		--entry-point=$(gcf_entry_point) \
		--trigger-http \
		--allow-unauthenticated

url:
	gcloud functions describe $(gcf_function_name) \
		--gen2 \
		--region=$(gcf_region) \
		--format="value(serviceConfig.uri)"

delete:
	gcloud functions delete $(gcf_function_name) \
		--gen2 \
		--region=$(gcf_region) 
