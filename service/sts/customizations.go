package sts

import "github.com/maccam912/aws-sdk-go/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		switch r.Operation.Name {
		case opAssumeRoleWithSAML, opAssumeRoleWithWebIdentity:
			r.Handlers.Sign.Clear() // these operations are unsigned
		}
	}
}
