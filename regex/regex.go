package gregex

import "regexp"

var PhoneRegex *regexp.Regexp
var IdentityCardRegex *regexp.Regexp
var SqlColumn *regexp.Regexp

func init() {
	PhoneRegex = regexp.MustCompile("^1(3\\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\\d|9[0-35-9])\\d{8}$")
	IdentityCardRegex = regexp.MustCompile("^[1-9]\\d{5}(18|19|([23]\\d))\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$")
	SqlColumn = regexp.MustCompile("[a-zA-Z]{1,1}[a-zA-Z0-9_]*")
}
