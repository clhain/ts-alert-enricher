# TS Alert Enrichment Webhook Receiver

Simple webhook receiver that catches incoming TS Alert WebHook API requests, and enriches alert data
before forwarding to an end destination. Built with [GCP Functions Framerwork for Go](https://github.com/GoogleCloudPlatform/functions-framework-go).

Incoming events are parsed for alert ID, and subsequent requests are made to the TS API for
alert context, and the first 50 related events. Due to limited TS Webhook authentication options, the
function expects and arbitrary api_key query parameter that must match the value supplied via the API_KEY 
environment variable (see env_example.yaml file for examples).
