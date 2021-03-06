/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"testing"
)

const (
	failed  = "\u2717"
	succeed = "\u2713"
)

func TestRemovePortFromHost(t *testing.T) {
	tests := []struct {
		name     string
		origHost string
		expect   string
	}{
		{"Domain&Port", "localhost:8080", "localhost"},
		{"Domain", "localhost", "localhost"},
		{"IPv4&Port", "192.168.0.1:8080", "192.168.0.1"},
		{"ShortestIPv6", "::", "::"},
		{"IPv6", "9878::7675:1292:9183:7562", "9878::7675:1292:9183:7562"},
		{"IPv6&Port", "[alsk:1204:1020::1292]:8080", "alsk:1204:1020::1292"},
		{"FQDN", " www.kubernetes.test:8080", " www.kubernetes.test"},
	}

	for _, tt := range tests {
		st := tt
		tf := func(t *testing.T) {
			t.Parallel()
			t.Logf("\tTestCase: %s", st.name)
			{
				get := RemovePortFromHost(st.origHost)
				if get != st.expect {
					t.Fatalf("\t%s\texpect %v, but get %v", failed, st.expect, get)
				}
				t.Logf("\t%s\texpect %v, get %v", succeed, st.expect, get)

			}
		}
		t.Run(st.name, tf)
	}
}
