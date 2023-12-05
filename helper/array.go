package helper

import (
	"bytes"
	"strconv"
	"strings"
)

func Intersection(a []string, b []string) []string {
	m := make(map[string]bool)
	r := make([]string, 0	)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b{
		if _, ok := m[v]; ok {
			r = append(r, v)
		}
	}
	return r
}

func StrLineToIntArray(a string) ([]int, error) {
	ret := make([]int, len(a))
	for i, v := range a {
		sv, err := strconv.Atoi(string(v))
		if err != nil {
			return ret, err
		}
		ret[i] = sv
	}
	return ret, nil
}

func StrArrayToInt(arr []string) ([]int, error) {
	ret := make([]int, len(arr))
	for i, v := range arr {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ret[i] = n
	}
	return ret, nil
}

func StrCsvToIntArray(s string, sep string) ([]int, error) {
	ret := make([]int, 0)
	nums := strings.Split(s, sep)
	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		ret = append(ret, n)
	}
	return ret, nil
}

func ContainsInt(i int, a []int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

func ContainsString(i string, a []string) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

func KeysStr(m map[string]int) (r []string) {
	r = make([]string, len(m))
	i := 0
	for k := range m {
		r[i] = k
	}
	return r
}

func PermuteStrings(parts ...[]string) (ret []string) {
	{
		var n = 1
		for _, ar := range parts {
			n *= len(ar)
		}
		ret = make([]string, 0, n)
	}
	var at = make([]int, len(parts))
	var buf bytes.Buffer
loop:
	for {
		// increment position counters
		for i := len(parts) - 1; i >= 0; i-- {
			if at[i] > 0 && at[i] >= len(parts[i]) {
				if i == 0 || (i == 1 && at[i-1] == len(parts[0])-1) {
					break loop
				}
				at[i] = 0
				at[i-1]++
			}
		}
		// construct permutated string
		buf.Reset()
		for i, ar := range parts {
			var p = at[i]
			if p >= 0 && p < len(ar) {
				buf.WriteString(ar[p])
			}
		}
		ret = append(ret, buf.String())
		at[len(parts)-1]++
	}
	return ret
}