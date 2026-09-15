package domain

import (
	"strings"
	"unicode"
)

func ClassifyRDAP(httpStatus int) Status {
	switch httpStatus {
	case 200:
		return StatusRegistered
	case 404:
		return StatusAvailable
	default:
		return StatusUnknown
	}
}

func ClassifyWhois(tld, fqdn, text string) Status {
	switch tld {
	case "cn":
		return classifyCN(fqdn, text)
	case "io", "me":
		return classifyNotFoundLine(text, "Domain not found.")
	case "co":
		return classifyCO(text)
	default:
		return StatusUnknown
	}
}

func classifyCN(fqdn, text string) Status {
	substance := substanceLines(text)
	joined := strings.Join(substance, "\n")
	low := strings.ToLower(joined)

	hasNMR := false
	hasDomain := false
	hasPending := false
	hasRestricted := strings.Contains(low, "can not be registered online")

	for _, line := range substance {
		l := strings.ToLower(line)
		if line == "No matching record." {
			hasNMR = true
		}
		if strings.HasPrefix(l, "domain name:") {
			hasDomain = true
		}
		if strings.Contains(l, "pendingdelete") {
			hasPending = true
		}
	}

	if hasRestricted && !hasNMR {
		return StatusReserved
	}
	if hasNMR && hasDomain {
		return StatusUnknown
	}
	if hasNMR && len(substance) == 1 {
		return StatusAvailable
	}
	if hasNMR && !hasDomain {
		// banners already stripped; leftover comments without Domain Name
		onlyNMR := true
		for _, line := range substance {
			if line != "No matching record." {
				onlyNMR = false
				break
			}
		}
		if onlyNMR {
			return StatusAvailable
		}
		return StatusUnknown
	}
	if hasPending && hasDomain {
		return StatusPendingDelete
	}
	if hasDomain {
		return StatusRegistered
	}
	_ = fqdn
	return StatusUnknown
}

func classifyNotFoundLine(text, emptyLine string) Status {
	substance := substanceLines(text)
	hasEmpty := false
	hasDomain := false
	for _, line := range substance {
		if line == emptyLine {
			hasEmpty = true
		}
		if strings.HasPrefix(strings.ToLower(line), "domain name:") {
			hasDomain = true
		}
	}
	if hasEmpty && hasDomain {
		return StatusUnknown
	}
	if hasEmpty {
		return StatusAvailable
	}
	if hasDomain {
		return StatusRegistered
	}
	return StatusUnknown
}

func classifyCO(text string) Status {
	substance := substanceLines(text)
	hasEmpty := false
	hasDomain := false
	for _, line := range substance {
		l := strings.ToLower(line)
		if strings.Contains(l, "domain not found") {
			hasEmpty = true
		}
		if strings.HasPrefix(l, "domain name:") {
			hasDomain = true
		}
	}
	if hasEmpty && hasDomain {
		return StatusUnknown
	}
	if hasEmpty {
		return StatusAvailable
	}
	if hasDomain {
		return StatusRegistered
	}
	return StatusUnknown
}

func substanceLines(text string) []string {
	var out []string
	for _, raw := range strings.Split(text, "\n") {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "%") || strings.HasPrefix(line, ">>>") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			continue
		}
		if isQueryBanner(line) {
			continue
		}
		out = append(out, line)
	}
	return out
}

func isQueryBanner(line string) bool {
	if !strings.Contains(line, "Querying") && !strings.Contains(strings.ToLower(line), "whois.") {
		return false
	}
	for _, r := range line {
		if unicode.IsLetter(r) && r > 127 {
			return false
		}
	}
	return strings.HasPrefix(line, "[") || strings.HasPrefix(strings.ToLower(line), "whois")
}
