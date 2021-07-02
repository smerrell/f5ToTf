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
	destination = "10.0.0.1"
	port        = "443"
}`
	server := &bigip.VirtualServer{
		Name:        "myvirtual_server",
		FullPath:    "/Common/myvirtual_server",
		Description: "This is a virtual server description",
		Destination: "/Common/10.0.0.1:443",
	}
	vs := ExportVirtualServer(server)
	if vs != template {
		t.Errorf("Expected: \n%v\nGot:\n%v", template, vs)
	}
}
