/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either extress or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package list

import (
	"reflect"
	"testing"
)

func TestDiff_same(t *testing.T) {
	left := []string{"elsa", "anna", "olaf", "kristoff"}
	right := []string{"elsa", "anna", "olaf", "kristoff"}
	extraLeft, extraRight := Diff(left, right)

	if !reflect.DeepEqual(extraLeft, []string{}) {
		t.Errorf("Didn't expect extra strings in left list but got %v", extraLeft)
	}

	if !reflect.DeepEqual(extraRight, []string{}) {
		t.Errorf("Didn't expect extra strings in right list but got %v", extraRight)
	}
}

func TestDiff_extraLeft(t *testing.T) {
	left := []string{"elsa", "anna", "olaf", "kristoff", "hans"}
	right := []string{"elsa", "anna", "olaf", "kristoff"}
	extraLeft, extraRight := Diff(left, right)

	if !reflect.DeepEqual(extraLeft, []string{"hans"}) {
		t.Errorf("Should have identified extra string in left list but got %v", extraLeft)
	}

	if !reflect.DeepEqual(extraRight, []string{}) {
		t.Errorf("Didn't expect extra strings in right list but got %v", extraRight)
	}
}

func TestDiff_extraRight(t *testing.T) {
	left := []string{"elsa", "anna", "olaf", "kristoff"}
	right := []string{"elsa", "anna", "olaf", "kristoff", "hans"}
	extraLeft, extraRight := Diff(left, right)
	if !reflect.DeepEqual(extraLeft, []string{}) {
		t.Errorf("Didn't expect extra strings in left list but got %v", extraLeft)
	}

	if !reflect.DeepEqual(extraRight, []string{"hans"}) {
		t.Errorf("Should have identified extra string in right list but got %v", extraRight)
	}
}

func TestDiff_extraBoth(t *testing.T) {
	left := []string{"elsa", "anna", "olaf", "kristoff", "pabbie"}
	right := []string{"elsa", "anna", "olaf", "kristoff", "hans"}
	extraLeft, extraRight := Diff(left, right)

	if !reflect.DeepEqual(extraLeft, []string{"pabbie"}) {
		t.Errorf("Should have identified extra string in left list but got %v", extraLeft)
	}

	if !reflect.DeepEqual(extraRight, []string{"hans"}) {
		t.Errorf("Should have identified extra string in right list but got %v", extraRight)
	}
}
