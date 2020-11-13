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
		if consecutiveHanLetterMixedSpace(test[0].(string)) != test[1].(bool) {
			t.Fatal(test)
		}
	}
}

func TestTrim(t *testing.T) {
	tests := [][]interface{}{
		{"我g掉了啊555", "我g掉了"},
		{"我g    掉了啊，", "我g    掉了"},
		{"我g掉    了啊！", "我g掉    了"},
		{"我g掉　　了啊~~", "我g掉　　了"},
		{"我g掉,了啊aaa...", "我g掉,了"},
		{"我g掉,了 啊!!!", "我g掉,了"},
		{"!!!啊啊啊！！！我g掉,了 啊,!!!cao!!!", "我g掉,了"},
	}
	for _, test := range tests {
		if trim(test[0].(string)) != test[1].(string) {
			t.Fatal(test, trim(test[0].(string)))
		}
	}
}

func TestShouldPin(t *testing.T) {
	tests := [][]interface{}{
		{"我g掉了啊555", true},
		{"我g    掉了啊，", false},
		{"我g掉    了啊！", false},
		{"我g掉　　了啊~~", false},
		{"wocao!! 我g掉了啊~~", true},
		{"我g掉,了啊aaa...", false},
		{"我g掉,了 啊!!!", false},
		{"!!!啊啊啊！！！我g掉,了 啊,!!!cao!!!", false},
		{"!!!啊啊啊！！！我g掉了 啊,!!!cao!!!", true},
		{"!!!啊啊啊！！！我this is a test了 啊,!!!cao!!!", true},
	}
	for _, test := range tests {
		if shouldPin(test[0].(string)) != test[1].(bool) {
			t.Fatal(test)
		}
	}
}
