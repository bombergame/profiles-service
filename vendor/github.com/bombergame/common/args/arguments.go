package args

import (
	"os"
	"regexp"
)

var (
	osArgs = os.Args
)

func GetString(name, defaultValue string) string {
	r, _ := regexp.Compile(`^--(?P<name>\w+)=(?P<value>.+)$`)
	for _, arg := range osArgs {
		mp := parseGroups(r, arg)
		if mp == nil {
			continue
		}

		if mp["name"] == name {
			return mp["value"]
		}
	}
	return defaultValue
}

func GetFlag(name string, defaultValue bool) bool {
	r, _ := regexp.Compile(`^--(?P<name>\w+)$`)
	for _, arg := range osArgs {
		mp := parseGroups(r, arg)
		if mp == nil {
			continue
		}

		if mp["name"] == name {
			return true
		}
	}
	return defaultValue
}

func parseGroups(r *regexp.Regexp, str string) map[string]string {
	n := len(r.SubexpNames())

	m := r.FindStringSubmatch(str)
	if len(m) != n {
		return nil
	}

	mp := make(map[string]string)
	for i, k := range r.SubexpNames() {
		mp[k] = m[i]
	}

	return mp
}
