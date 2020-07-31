package converter

import (
	"regexp"
	"strings"
)

func SanitizeTerraformName(name string) string {
	r := strings.NewReplacer("-", "_", ".", "_", " ", "_")
	return r.Replace(name)
}

func ipport(ipport string) []string {
	regex := regexp.MustCompile(`(\/.+\/)((?:[0-9]{1,3}\.){3}[0-9]{1,3})(\%\d+)?\:(\d+)`)
	return regex.FindStringSubmatch(ipport)
}

func IPAddress(address string) string {
	addr := ipport(address)
	return addr[2]
}

func IPPort(address string) string {
	addr := ipport(address)
	return addr[4]
}
