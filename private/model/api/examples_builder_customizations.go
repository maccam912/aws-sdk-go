// +build codegen

package api

import (
	"bytes"
	"fmt"
)

type wafregionalExamplesBuilder struct {
	defaultExamplesBuilder
}

func (builder wafregionalExamplesBuilder) Imports(a *API) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`"fmt"
	"strings"
	"time"

	"github.com/maccam912/aws-sdk-go/aws"
	"github.com/maccam912/aws-sdk-go/aws/awserr"
	"github.com/maccam912/aws-sdk-go/aws/session"
	"github.com/maccam912/aws-sdk-go/service/waf"
	`)

	buf.WriteString(fmt.Sprintf("\"%s/%s\"", "github.com/maccam912/aws-sdk-go/service", a.PackageName()))
	return buf.String()
}
