package main

type Env string

var (
	Staging    Env = "staging"
	Production Env = "production"
)

var (
	// ES ports
	ES_STAGING_LOCAL_PORT  int = 9201
	ES_STAGING_REMOTE_PORT int = 9200

	ES_PRODUCTION_LOCAL_PORT  int = 9202
	ES_PRODUCTION_REMOTE_PORT int = 9200
)
