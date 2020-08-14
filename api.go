package txlbs

import "fmt"

const (
	base = "https://apis.map.qq.com"
)

// API represents a relative path
type API string

const (
	districtList        API = "/ws/district/v1/list"
	districtGetChildren API = "/ws/district/v1/getchildren"
	districtSearch      API = "/ws/district/v1/search"
	ipLocation          API = "/ws/location/v1/ip"
)

// Full returns the full path.
func (a API) Full(p string) string {
	return fmt.Sprintf("%s?%s", a, p)
}
