package generator

import (
	"fmt"
	"strings"
)

type stringSet []string

func newStringSet() stringSet {
	return stringSet(make([]string, 0))
}

func (set *stringSet) Contains(s string) bool {
	for _, cur := range *set {
		if cur == s {
			return true
		}
	}
	return false
}

func (set *stringSet) Add(s string) {
	for i, cur := range *set {
		if cur == s {
			return
		} else if cur > s {
			*set = append(*set, "")
			copy((*set)[i+1:], (*set)[i:])
			(*set)[i] = s
			return
		}
	}
	*set = append(*set, s)
}

func (set *stringSet) Remove(s string) {
	for i, cur := range *set {
		if cur == s {
			(*set) = append((*set)[:i], (*set)[i+1:]...)
			return
		}
	}
}

func (set *stringSet) Equals(set2 stringSet) bool {
	if len(*set) != len(set2) {
		return false
	}

	for i, cur := range *set {
		if cur != set2[i] {
			return false
		}
	}
	return true
}

func (set *stringSet) Union(set2 stringSet) stringSet {
	unionSet := newStringSet()

	unionSet = append(unionSet, *set...)

	for _, cur := range set2 {
		unionSet.Add(cur)
	}

	return unionSet
}

func (set *stringSet) Intersection(set2 stringSet) stringSet {
	intersectionSet := newStringSet()

	for _, cur := range *set {
		if set2.Contains(cur) {
			intersectionSet.Add(cur)
		}
	}

	return intersectionSet
}

func (set *stringSet) Difference(set2 stringSet) stringSet {
	differenceSet := newStringSet()

	for _, cur := range *set {
		if !set2.Contains(cur) {
			differenceSet.Add(cur)
		}
	}

	return differenceSet
}

func (set *stringSet) Copy() stringSet {
	newSet := make([]string, len(*set))

	copy(newSet, *set)

	return newSet
}

func (set *stringSet) String() string {
	return fmt.Sprintf("[%s]", strings.Join(*set, " "))
}
