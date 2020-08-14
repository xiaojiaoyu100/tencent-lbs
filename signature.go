package txlbs

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strings"
)

func (c *Client) signatureQueryParameter(v url.Values, path string) (url.Values, error) {
	v.Set("key", c.Config.Key)
	toSig := string(path) + "?" + sortQueryParameter(v) + c.SecretKey
	sig, err := Md5(toSig)
	if err != nil {
		return nil, err
	}
	v.Set("sig", sig)
	return v, nil
}

func sortQueryParameter(v url.Values) string {
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var ret string
	for _, k := range keys {
		vs := v[k]
		if len(vs) == 0 {
			continue
		}
		for _, s := range vs {
			ret += k + "=" + s + "&"
		}
	}
	ret = strings.TrimSuffix(ret, "&")
	return ret
}

// Md5 calculates a string digest.
func Md5(str string) (string, error) {
	w := md5.New()
	_, err := io.WriteString(w, str)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", w.Sum(nil)), nil
}
