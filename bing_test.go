package search

import (
	"testing"
)

func TestBingSearch(t *testing.T) {
	cases := []Options{
		// keyword test
		{"test", 30, false},
		// keyword test with tor
		{"test", 20, true},
		// keyword facebook
		{"facebook", 25, false},
		// how to cook with tor
		{"how to cook", 25, true},
	}

	for _, options := range cases {
		results, _ := WebSearch(options)
		if 0 > len(results) {
			t.Errorf("there is error len of results %d != options.count %d", len(results), options.Count)
		}
	}
}

func TestURLInfo(t *testing.T) {
	cases := []struct {
		url string
		tor bool
	}{
		// keyword test
		{"http://wikipedia.org", false},
		{"http://techcrunch.com/2015/05/12/the-ultimate-interface-is-your-brain/", false},
		{"http://facebook.com", true},
	}
	for _, options := range cases {
		info, ok := URLInfo(options.url, options.tor)
		if !ok {
			t.Errorf("found error with connection")
		}

		if info.Description == "" {
			t.Errorf("no description, means failed to fetch url info, or url is invalid")
		}
	}
}
