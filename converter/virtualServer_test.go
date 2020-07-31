package converter

import (
	"testing"

	"github.com/f5devcentral/go-bigip"
)

func TestNilVirtualServerReturnsEmptyString(t *testing.T) {
	vs := ExportVirtualServer(nil)
	if len(vs) > 0 {
		t.Errorf("\nExpected empty string: got\n'%v'", vs)
	}
}

func TestVirtualServerExport(t *testing.T) {
	template := `resource "bigip_ltm_virtual_server" "myvirtual_server" {
	name        = "/Common/myvirtual_server"
	description = "This is a virtual server description"
}`
	server := &bigip.VirtualServer{
		Name:        "myvirtual_server",
		FullPath:    "/Common/myvirtual_server",
		Description: "This is a virtual server description",
	}
	vs := ExportVirtualServer(server)
	if vs != template {
		t.Errorf("Expected: \n%v\nGot:\n%v", template, vs)
	}
}
