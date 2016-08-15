/*
Copyright 2016 Skippbox, Ltd All rights reserved.

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

package app

import (
	"fmt"
	"testing"

	"github.com/skippbox/kompose/pkg/transformer"
)

func TestParseVolume(t *testing.T) {
	name1 := "datavolume"
	host1 := "./cache"
	host2 := "~/configs"
	container1 := "/tmp/cache"
	container2 := "/etc/configs/"
	mode := "rw"

	tests := []struct {
		test, volume, name, host, container, mode string
	}{
		{
			"name:host:container:mode",
			fmt.Sprintf("%s:%s:%s:%s", name1, host1, container1, mode),
			name1,
			host1,
			container1,
			mode,
		},
		{
			"host:container:mode",
			fmt.Sprintf("%s:%s:%s", host2, container2, mode),
			"",
			host2,
			container2,
			mode,
		},
		{
			"name:container:mode",
			fmt.Sprintf("%s:%s:%s", name1, container1, mode),
			name1,
			"",
			container1,
			mode,
		},
		{
			"name:host:container",
			fmt.Sprintf("%s:%s:%s", name1, host1, container1),
			name1,
			host1,
			container1,
			"",
		},
		{
			"host:container",
			fmt.Sprintf("%s:%s", host1, container1),
			"",
			host1,
			container1,
			"",
		},
		{
			"container:mode",
			fmt.Sprintf("%s:%s", container2, mode),
			"",
			"",
			container2,
			mode,
		},
		{
			"name:container",
			fmt.Sprintf("%s:%s", name1, container1),
			name1,
			"",
			container1,
			"",
		},
		{
			"container",
			fmt.Sprintf("%s", container2),
			"",
			"",
			container2,
			"",
		},
	}

	for _, test := range tests {
		name, host, container, mode, err := transformer.ParseVolume(test.volume)
		if err != nil {
			t.Errorf("In test case %q, returned unexpected error %v", test.test, err)
		}
		if name != test.name {
			t.Errorf("In test case %q, returned volume name %s, expected %s", name, test.name)
		}
		if host != test.host {
			t.Errorf("In test case %q, returned host path %s, expected %s", host, test.host)
		}
		if container != test.container {
			t.Errorf("In test case %q, returned container path %s, expected %s", container, test.container)
		}
		if mode != test.mode {
			t.Errorf("In test case %q, returned access mode %s, expected %s", mode, test.mode)
		}
	}
}
