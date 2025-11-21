package allinkl

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	allinkl "github.com/libdns/all-inkl"
)

func TestUnmarshalCaddyFile(t *testing.T) {
	tests := []string{
		`allinkl {
			kas_username username_test
			kas_password password_test
		}`}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			// given
			dispenser := caddyfile.NewTestDispenser(tc)
			p := Provider{&allinkl.Provider{}}
			// when
			err := p.UnmarshalCaddyfile(dispenser)
			// then
			if err != nil {
				t.Errorf("UnmarshalCaddyfile failed with %v", err)
				return
			}

			expectedUsername := "username_test"
			actualUsername := p.Provider.KasUsername
			if expectedUsername != actualUsername {
				t.Errorf("Expected Username to be '%s' but got '%s'", expectedUsername, actualUsername)
			}

			expectedAPIPassword := "password_test"
			actualAPIPassword := p.Provider.KasPassword
			if expectedAPIPassword != actualAPIPassword {
				t.Errorf("Expected APIPassword to be '%s' but got '%s'", expectedAPIPassword, actualAPIPassword)
			}
		})
	}
}
