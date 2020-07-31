package converter

import "strings"

func SanitizeTerraformName(name string) string {
	r := strings.NewReplacer("-", "_", ".", "_", " ", "_")
	return r.Replace(name)
}
