package main

import (
	"encoding/json"
	"os"
	"testing"
)

func contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func testHelper(t *testing.T, testName string, extracted []string, expected []string) {
	t.Run(testName, func(t *testing.T) {
		if len(extracted) == 0 && len(expected) == 0 {
			return
		}

		for _, answer := range expected {
			if !contains(extracted, answer) {
				t.Fatalf(`"%s" should be in extracted: %v`, answer, extracted)
			}
		}

		for _, e := range extracted {
			if !contains(expected, e) {
				t.Fatalf(`"%s" extracted but not in expected: %v`, e, expected)
			}
		}
	})
}

func TestExtractAddress(t *testing.T) {
	var tests map[string]map[string][]string

	f, err := os.Open("tests.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&tests)

	testTypes := []string{"domains", "ip4s", "urls", "emails"}
	for input, expectedOutputs := range tests {
		for _, testType := range testTypes {
			var extracted []string
			if testType == "domains" {
				extracted = ExtractDomains(input)
			} else if testType == "ip4s" {
				extracted = ExtractIPv4s(input)
			} else if testType == "urls" {
				extracted = ExtractURLs(input)
			} else if testType == "emails" {
				extracted = ExtractEmails(input)
			} else {
				t.Fatal("wat")
			}

			expected, ok := expectedOutputs[testType]
			if !ok {
				expected = []string{}
			}

			testHelper(t, testType+"=>"+input, extracted, expected)
		}
	}
}
