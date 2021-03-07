package filter

import "testing"

func Test_MatchDefaultUrl(t *testing.T) {
	config := Config{}
	config = configDefault(config)

	if matched := match(config.UrlPattern, "/", config.CaseSensitive); !matched {
		t.Error("default url not match /")
	}
	if matched := match(config.UrlPattern, "/api", config.CaseSensitive); !matched {
		t.Error("default url not match /api")
	}
	if matched := match(config.UrlPattern, "/api/v1", config.CaseSensitive); !matched {
		t.Error("default url not match /api/v1")
	}
}

func Test_MatchUrlRegexOneStar1(t *testing.T) {
	config := Config{
		UrlPattern: []string{"/api/*"},
	}
	config = configDefault(config)

	if config.UrlPattern[0] != "/api/*" {
		t.Error("config url invalid")
	}
	// case sensitive true
	// must match is correct
	if matched := match(config.UrlPattern, "/api/insert", true); !matched {
		t.Error("url /api/* match /api/insert with case sensitive true")
	}
	// case sensitive false
	// must not match is correct
	if matched := match(config.UrlPattern, "/Api/insert", false); matched {
		t.Error("url /api/* match /Api/insert with case sensitive false")
	}
	// other prefix not startWith /api
	// must not match is correct
	if matched := match(config.UrlPattern, "/checking", false); matched {
		t.Error("url /api/* match /checking")
	}
}

func Test_MatchUrlRegexOneStar2(t *testing.T) {
	config := Config{
		UrlPattern: []string{"/api/*/v1"},
	}
	config = configDefault(config)

	if matched := match(config.UrlPattern, "/api/insert/v1", false); !matched {
		t.Error("url /api/*/v1 not match /api/insert/v1")
	}
	if matched := match(config.UrlPattern, "/api/update/v1", false); !matched {
		t.Error("url /api/*/v1 not match /api/update/v1")
	}
}
