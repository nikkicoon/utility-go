package pkg

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// GetFQDN uses the `hostname` binary to look up the
// fully qualified domain name.
// It expects to find aaa.bbb.tld or something longer than
// 1 element if split by dots in the name.
// The expected length will be bypassed for mocking.
func GetFQDN() (string, error) {
	// this sucks, but there's no standard way and
	// not calling an external binary involves a lot
	// of code
	// TODO: solve how to do this without the binary (os.Hostname() won't do)
	// one possible package to vendor: https://github.com/Showmax/go-fqdn/blob/master/fqdn.go
	path, err := exec.LookPath("hostname")
	if err != nil {
		return "", fmt.Errorf("hostname binary not found: %w", err)
	}
	cmd := exec.Command(path, "-f")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("getFQDN error: %w", err)
	}
	fqdn := out.String()
	fqdn = fqdn[:len(fqdn)-1] // remove EOL
	if len(strings.Split(fqdn, ".")) <= 1 {
		return "", fmt.Errorf("fqdn is shorter than expected (1 or less elements separated by dots), looks like a hostname")
	}
	return fqdn, nil
}
