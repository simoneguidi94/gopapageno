package generator

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringSet(t *testing.T) {
	set := newStringSet()
	set.Add("hello")
	set.Add("hello")

	expected := []string{"hello"}
	if !set.Equals(expected) {
		t.Error(fmt.Sprintf("Error: set is [%s], should be [%s]", strings.Join(set, ","), strings.Join(expected, ",")))
	}

	set.Remove("hello")

	expected = []string{}
	if !set.Equals(expected) {
		t.Error(fmt.Sprintf("Error: set is [%s], should be [%s]", strings.Join(set, ","), strings.Join(expected, ",")))
	}

	set = newStringSet()
	set.Add("hello")
	set.Add("cat")

	set2 := newStringSet()
	set.Add("hello")
	set.Add("dog")
	unionSet := set.Union(set2)

	expected = []string{"cat", "dog", "hello"}
	if !unionSet.Equals(expected) {
		t.Error(fmt.Sprintf("Error: set is [%s], should be [%s]", strings.Join(unionSet, ","), strings.Join(expected, ",")))
	}

	set = newStringSet()
	set.Add("hello")
	set.Add("cat")
	set.Add("dog")

	set2 = newStringSet()
	set2.Add("dog")
	set2.Add("cat")

	intersectionSet := set.Intersection(set2)
	expected = []string{"cat", "dog"}
	if !intersectionSet.Equals(expected) {
		t.Error(fmt.Sprintf("Error: set is [%s], should be [%s]", strings.Join(intersectionSet, ","), strings.Join(expected, ",")))
	}

	set = newStringSet()
	set.Add("hello")
	set.Add("cat")
	set.Add("dog")

	set2 = newStringSet()
	set2.Add("dog")
	set2.Add("cat")

	differenceSet := set.Difference(set2)
	expected = []string{"hello"}
	if !differenceSet.Equals(expected) {
		t.Error(fmt.Sprintf("Error: set is [%s], should be [%s]", strings.Join(differenceSet, ","), strings.Join(expected, ",")))
	}
}
