package util

const (
	CACHE_TURNOFF_THRESHOLD = 65536
)

var stringMatcherCache map[string]*PathStringMatcher

type PathStringMatcher struct {
	Pattern string
}

func GetStringMatcher(pattern string) *PathStringMatcher {
	if stringMatcherCache == nil {
		stringMatcherCache = make(map[string]*PathStringMatcher)
	}
	if matcher, ok := stringMatcherCache[pattern]; ok {
		return matcher
	}
	matcher := &PathStringMatcher{Pattern: pattern}
	if len(stringMatcherCache) > CACHE_TURNOFF_THRESHOLD {
		stringMatcherCache = make(map[string]*PathStringMatcher)
		return matcher
	}
	stringMatcherCache[pattern] = matcher
	return matcher
}

func (m *PathStringMatcher) MatchString(str string) (bool, string, string) {
	if len(m.Pattern) <= 1 || m.Pattern[0] != ':' {
		return m.Pattern == str, "", ""
	}
	if m.Pattern[0] == ':' {
		return true, m.Pattern, str
	}
	return false, "", ""
}
