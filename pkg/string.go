package pkg

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

// ConvertLineToCRLF converts a line ending of \n or \r to
// \r\n (CR LF)
func ConvertLineToCRLF(s string) string {
	start := time.Now()
	// ldif reads a line until `lf` and trims right `sep` from line
	var cr byte = '\x0D'
	var lf byte = '\x0A'
	var sep = string([]byte{cr, lf})
	// CR LF \r\n (windows)
	var windowsEOLActual = "\r\n"
	var windowsEOLEscaped = `\r\n`
	// CF \r (mac)
	var macEOLActual = "\r"
	var macEOLEscaped = `\r`
	// LF \n (unix)
	var unixEOLActual = "\n"
	var unixEOLEscaped = `\n`
	r := strings.NewReplacer(windowsEOLActual, sep, windowsEOLEscaped, sep, macEOLActual, sep, macEOLEscaped, sep, unixEOLActual, sep, unixEOLEscaped, sep)
	res := r.Replace(s)
	fmt.Fprintf(os.Stderr, "execution time of %s: %s\n", GetCurrentFuncName(), time.Since(start).String())
	return res
}

// DissolveEmptyValues removes all empty value lines in `s`.
// This is a work-around for https://github.com/go-ldap/ldif/issues/21.
func DissolveEmptyValues(s string) string {
	start := time.Now()
	re := regexp.MustCompile("(?m)^[^#][a-zA-Z]+:$[\n\r]")
	res := re.ReplaceAllString(s, "")
	fmt.Fprintf(os.Stderr, "execution time of %s: %s\n", GetCurrentFuncName(), time.Since(start).String())
	return res
}

// DissolveDoubleColon replaces all `t::` with `t:` in `s`.
// This is a work-around for https://github.com/go-ldap/ldif/issues/23.
func DissolveDoubleColon(s string) string {
	start := time.Now()
	re := regexp.MustCompile("(?m)(^[^#][a-zA-Z]+:):")
	res := re.ReplaceAllString(s, "$1")
	fmt.Fprintf(os.Stderr, "execution time of %s: %s\n", GetCurrentFuncName(), time.Since(start).String())
	return res
}

func SplitMailString(s string) (string, string) {
	u := strings.Split(s, "@")
	return u[0], u[1]
}

// ContainsMultiple checks if the string `s` contains any of the substrings in
// the string slice `input`. The function returns true for the first match
// which returned true, else it returns false.
func ContainsMultiple(s string, input []string) bool {
	res := false
	for _, v := range input {
		res = strings.Contains(s, v)
		if res {
			return res
		}
	}
	return res
}

// HasSuffixMultiple checks if the string `s` has any of the suffixes
// in the string slice `input`. The function returns true for the first
// match which returned true, else it returns false.
func HasSuffixMultiple(s string, input []string) bool {
	res := false
	for _, v := range input {
		res = strings.HasSuffix(s, v)
		if res {
			return res
		}
	}
	return res
}
