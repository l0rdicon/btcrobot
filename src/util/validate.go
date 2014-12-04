package util

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// Verify form data
func Validate(data url.Values, rules map[string]map[string]map[string]string) (errMsg string) {
	for field, rule := range rules {
		val := data.Get(field)
		// Check for required info
		if requireInfo, ok := rule["require"]; ok {
			if val == "" {
				errMsg = requireInfo["error"]
				return
			}
		}
		// Check the length
		if lengthInfo, ok := rule["length"]; ok {
			valLen := len(val)
			if lenRange, ok := lengthInfo["range"]; ok {
				errMsg = checkRange(valLen, lenRange, lengthInfo["error"])
				if errMsg != "" {
					return
				}
			}
		}
		// Check for type int
		if intInfo, ok := rule["int"]; ok {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				errMsg = field + "Type error！"
				return
			}
			if intRange, ok := intInfo["range"]; ok {
				errMsg = checkRange(valInt, intRange, intInfo["error"])
				if errMsg != "" {
					return
				}
			}
		}
		// Check email address
		if emailInfo, ok := rule["email"]; ok {
			validEmail := regexp.MustCompile(`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$`)
			if !validEmail.MatchString(val) {
				errMsg = emailInfo["error"]
				return
			}
		}
		// Check values match 
		if compareInfo, ok := rule["compare"]; ok {
			compared := compareInfo["field"] // field to be compaired
			// Comparison rules
			switch compareInfo["rule"] {
			case "=":
				if val != data.Get(compared) {
					errMsg = compareInfo["error"]
					return
				}
			case "<":
			case ">":
			default:

			}
		}
	}
	return
}

// checkRange of values is good
// src value to be checked; destRange target range; msg error message when parameters
func checkRange(src int, destRange string, msg string) (errMsg string) {
	parts := strings.SplitN(destRange, ",", 2)
	parts[0] = strings.TrimSpace(parts[0])
	parts[1] = strings.TrimSpace(parts[1])
	min, max := 0, 0
	if parts[0] == "" {
		max = MustInt(parts[1])
		if src > max {
			errMsg = fmt.Sprintf(msg, max)
			return
		}
	}
	if parts[1] == "" {
		min = MustInt(parts[0])
		if src < min {
			errMsg = fmt.Sprintf(msg, min)
			return
		}
	}
	if min == 0 {
		min = MustInt(parts[0])
	}
	if max == 0 {
		max = MustInt(parts[1])
	}
	if src < min || src > max {
		errMsg = fmt.Sprintf(msg, min, max)
		return
	}
	return
}
