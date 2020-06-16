package converter

import (
	"strings"
	"text/template"

	"github.com/f5devcentral/go-bigip"
)

func Export(vs *bigip.VirtualServer) string {
	vstemplate := `resource "bigip_ltm_virtual_server" "{{.Name}}" {
	name        = "{{.FullPath}}"
	description = "{{.Description}}"
}`

	var builder strings.Builder

	if vs == nil {
		return ""
	}

	tmpl, err := template.New("vs").Parse(vstemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&builder, vs)
	if err != nil {
		panic(err)
	}
	return builder.String()
}
