package sqs

import "github.com/maccam912/aws-sdk-go/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		setupChecksumValidation(r)
	}
}
