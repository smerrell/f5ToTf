package converter

import (
	"strings"
	"text/template"

	"github.com/f5devcentral/go-bigip"
)

func ExportVirtualServer(vs *bigip.VirtualServer) string {
	funcMap := template.FuncMap{
		"sanitize":  SanitizeTerraformName,
		"ipaddress": IPAddress,
		"ipport":    IPPort,
	}
	vstemplate := `resource "bigip_ltm_virtual_server" "{{.Name | sanitize}}" {
	name        = "{{.FullPath}}"
	description = "{{.Description}}"
	destination = "{{.Destination | ipaddress}}"
	port        = "{{.Destination | ipport}}"
}`

	var builder strings.Builder

	if vs == nil {
		return ""
	}

	tmpl, err := template.New("vs").Funcs(funcMap).Parse(vstemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&builder, vs)
	if err != nil {
		panic(err)
	}
	return builder.String()
}
