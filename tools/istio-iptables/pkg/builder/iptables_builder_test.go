// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builder

import (
	"testing"
)

func Test_BuildV6Restore(t *testing.T) {
	iptables := NewIptablesBuilder()
	expected := ""
	actual := iptables.BuildV6Restore()
	if expected != actual {
		t.Errorf("Output didn't match: Got: %s, Expected: %s", actual, expected)
	}
}

func Test_BuildV4Restore(t *testing.T) {
	iptables := NewIptablesBuilder()
	expected := ""
	actual := iptables.BuildV4Restore()
	if expected != actual {
		t.Errorf("Output didn't match: Got: %s, Expected: %s", actual, expected)
	}
}

func Test_BuildV4(t *testing.T) {
	iptables := NewIptablesBuilder()
	iptables.AppendRuleV4("chain", "table", "-f", "foo", "-b", "bar")
	iptables.InsertRuleV4("chain", "table", 2, "-f", "foo", "-b", "bar")
	if err := len(iptables.rules.rulesv6) != 0; err {
		t.Errorf("Expected rulesV6 to be empty; but got %v", iptables.rules.rulesv6)
	}
	if err := len(iptables.rules.rulesv4) != 2; err {
		t.Errorf("Expected rulesV4 to be not empty; but got %v", iptables.rules.rulesv4)
	}
	actual := iptables.BuildV4()
	if err := len(actual) != 2; err {
		t.Errorf("Expected actual to have single element; but got %v", actual)
	}
	expected := []string{"iptables -t table -A chain -f foo -b bar", "iptables -t table -I chain 2 -f foo -b bar"}
	for index := range expected {
		if actual[index] != expected[index] {
			t.Errorf("Output mismatch. Actual: %v, Expected: %v", actual, expected)
		}
	}
}

func Test_BuildV6(t *testing.T) {
	iptables := NewIptablesBuilder()
	iptables.AppendRuleV6("chain", "table", "-f", "foo", "-b", "bar")
	iptables.InsertRuleV6("chain", "table", 2, "-f", "foo", "-b", "bar")
	if err := len(iptables.rules.rulesv4) != 0; err {
		t.Errorf("Expected rulesV4 to be empty; but got %v", iptables.rules.rulesv4)
	}
	if err := len(iptables.rules.rulesv6) != 2; err {
		t.Errorf("Expected rulesV6 to be not empty; but got %v", iptables.rules.rulesv6)
	}
	actual := iptables.BuildV6()
	if err := len(actual) != 2; err {
		t.Errorf("Expected actual to have single element; but got %v", actual)
	}
	expected := []string{"ip6tables -t table -A chain -f foo -b bar", "ip6tables -t table -I chain 2 -f foo -b bar"}
	for index := range expected {
		if actual[index] != expected[index] {
			t.Errorf("Output mismatch. Actual: %v, Expected: %v", actual, expected)
		}
	}
}
