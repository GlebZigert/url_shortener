package convertIPtoCIDR

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test(t *testing.T) {

	tests := []struct {
		name    string
		ipStart string
		ipEnd   string
		ip      string
		check   bool
	}{
		{
			name:    "1",
			ipStart: "192.168.1.10",
			ipEnd:   "192.168.1.17",
			ip:      "192.168.1.12",
			check:   true,
		},
		{
			name:    "1",
			ipStart: "192.168.1.10",
			ipEnd:   "192.168.1.17",
			ip:      "192.168.1.18",
			check:   false,
		},
		{
			name:    "2",
			ipStart: "192.168.1.10",
			ipEnd:   "192.168.2.17",
			ip:      "192.168.1.12",
			check:   true,
		},
		{
			name:    "2",
			ipStart: "192.167.1.10",
			ipEnd:   "192.168.2.17",
			ip:      "192.168.1.12",
			check:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ipStart := test.ipStart
			ipEnd := test.ipEnd

			cidr, err := iPv4RangeToCIDR(ipStart, ipEnd)
			if err != nil {
				t.Error(err.Error())
			}
			t.Log("cidr", cidr)

			resStart, resEnd, err := CIDRRangeToIPv4Range(cidr)
			if err != nil {
				t.Error(err.Error())
			}
			assert.Equal(t, ipStart, resStart)
			assert.Equal(t, ipEnd, resEnd)

			assert.Equal(t, test.check, check(test.ip, cidr))

		})
	}
}
