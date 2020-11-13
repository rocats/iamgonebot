package main

import "testing"

func TestConsecutiveHanMixedSpace(t *testing.T) {
	tests := [][]interface{}{
		{"我g掉了啊", false},
		{"我g    掉了啊", false},
		{"我g掉    了啊", true},
		{"我g掉　　了啊", true},
		{"我g掉,了啊", false},
		{"我g掉,了 啊", true},
	}
	for _, test := range tests {
		if consecutiveHanMixedSpace(test[0].(string)) != test[1].(bool) {
			t.Fatal(test)
		}
	}
}
