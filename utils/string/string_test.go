package stringformat

import (
	"strings"
	"testing"
)

type TestDataItem struct {
	value        string
	replacements map[string]string
	output       string
}

type TestLogDataItem struct {
	msg            string
	status         string
	statusExpected string
}

func TestReplaceStringSubvals(t *testing.T) {
	dataItems := []TestDataItem{
		{"R$7,00", map[string]string{"R$": "", ",": ".", " ": ""}, "7.00"},
		{"", map[string]string{"R$": "", ",": "", " ": ""}, ""},
		{"Removing spaces from sentence", map[string]string{" ": ""}, "Removingspacesfromsentence"},
		{"Removing spaces from sentence and \nbreaklines\n", map[string]string{}, "Removingspacesfromsentenceandbreaklines"},
	}

	for _, item := range dataItems {
		val := ReplaceStringSubvals(item.value, item.replacements)

		if val == item.output {
			t.Logf("ReplaceStringSubvals() with args (value: %v, replacements: %v) PASSED, expected output '%v' and got output '%v'", item.value, item.replacements, item.output, val)
		} else {
			t.Errorf("ReplaceStringSubvals() with args (value: %v, replacements: %v) FAILED, expected output '%v' but got output '%v'", item.value, item.replacements, item.output, val)
		}
	}
}

func TestLogMessage(t *testing.T) {
	dataItems := []TestLogDataItem{
		{"Error happened!", "", "INFO"},
		{"New status created!", "ERROR", "ERROR"},
	}

	for _, item := range dataItems {
		msg := LogMessage(item.msg, item.status)

		if strings.Contains(msg, item.msg) && strings.Contains(msg, item.statusExpected) {
			t.Logf("LogMessage() with args (msg: %v, status: %v) PASSED, expected status '%v' and got output '%v'", item.msg, item.status, item.statusExpected, msg)
		} else {
			t.Errorf("LogMessage() with args (msg: %v, status: %v) FAILED, expected status '%v' but got output '%v'", item.msg, item.status, item.statusExpected, msg)
		}
	}
}
