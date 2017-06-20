package gonzh

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	// regular expression for scientific notification
	sciRegEx = regexp.MustCompile(`^(?P<sign>[+-])?0*(?P<integer>\d+)(?P<decimal>\.(\d+))?e((?P<esign>[+-])?(?P<exp>\d+))$`)
	// regular expression for ordinary number
	numRegEx   = regexp.MustCompile(`^(?P<sign>[+-])?0*(?P<integer>\d+)(?P<decimal>\.(\d+))?$`)
	numCapital = map[string]string{
		"0": "零",
		"1": "壹",
		"2": "贰",
		"3": "叁",
		"4": "肆",
		"5": "伍",
		"6": "陆",
		"7": "柒",
		"8": "捌",
		"9": "玖",
	}
	numLower = map[string]string{
		"0": "零",
		"1": "一",
		"2": "二",
		"3": "三",
		"4": "四",
		"5": "五",
		"6": "六",
		"7": "七",
		"8": "八",
		"9": "九",
	}
	siCapital = []string{"个", "拾", "佰", "仟", "万", "亿", "兆", "京", "垓", "秭", "穰", "沟", "涧", "正", "载"}
	siLower   = []string{"个", "十", "百", "千", "万", "亿", "兆", "京", "垓", "秭", "穰", "沟", "涧", "正", "载"}
)

// EncodeString converts a number into Chinese representation, if asMoney true
// returns in RMB style
func EncodeString(num string, asMoney, isCapital bool) string {
	parts := parseNumber(num)
	e := convert(parts)
	switch e := e.(type) {
	case sciNotation:
		return encodeSciNotation(e, asMoney, isCapital)
	case number:
		return encodeNumber(e, asMoney, isCapital)
	default:
		return num
	}
}

type number struct {
	minus   bool
	integer string
	decimal string
}

type sciNotation struct {
	minus   bool
	eminus  bool
	integer string
	decimal string
	exp     int
}

func parseNumber(num string) map[string]string {
	var parts map[string]string
	var names []string

	match := sciRegEx.FindStringSubmatch(num)
	if match != nil {
		names = sciRegEx.SubexpNames()
	} else {
		match = numRegEx.FindStringSubmatch(num)
		if match != nil {
			names = numRegEx.SubexpNames()
		}
	}

	if names != nil {
		parts = make(map[string]string)
	}

	for i, n := range names {
		if i > 0 && i <= len(match) && n != "" {
			parts[n] = match[i]
		}
	}

	return parts
}

func convert(parts map[string]string) interface{} {
	e, ok := parts["exp"]
	if ok {
		exp, _ := strconv.Atoi(e)
		return sciNotation{
			minus:   parts["sign"] == "-",
			eminus:  parts["esign"] == "-",
			integer: parts["integer"],
			decimal: parts["decimal"],
			exp:     exp,
		}
	}

	return number{
		minus:   parts["sign"] == "-",
		integer: parts["integer"],
		decimal: parts["decimal"],
	}
}

func encodeSciNotation(n sciNotation, asMoney, isCapital bool) string {
	var result number
	decimal := n.decimal
	if decimal != "" {
		decimal = decimal[1:]
	}
	if n.eminus {
		diff := n.exp - len(n.integer)
		if diff >= 0 {
			var zeros string
			for i := 0; i < diff; i++ {
				zeros += "0"
			}
			result = number{
				minus:   n.minus,
				integer: "0",
				decimal: "." + zeros + n.integer + decimal,
			}
		} else {
			result = number{
				minus:   n.minus,
				integer: n.integer[:-diff],
				decimal: "." + n.integer[-diff:] + decimal,
			}
		}
	} else {
		diff := n.exp - len(decimal)
		if diff >= 0 {
			var zeros string
			for i := 0; i < diff; i++ {
				zeros += "0"
			}
			var integer string
			for _, i := range n.integer {
				if string(i) != "0" {
					integer += string(i)
				}
			}
			result = number{
				minus:   n.minus,
				integer: integer + zeros,
			}
		} else {
			result = number{
				minus:   n.minus,
				integer: n.integer + decimal[:n.exp],
				decimal: "." + decimal[n.exp:],
			}
		}
	}
	return encodeNumber(result, asMoney, isCapital)
}

func encodeNumber(n number, asMoney, isCapital bool) string {
	num := numLower
	si := siLower
	if isCapital {
		num = numCapital
		si = siCapital
	}

	var sign string
	if n.minus {
		sign = "负"
	}
	integer := encodeInteger(n.integer, num, si, 3)
	decimal := encodeDecimal(n.decimal, num)
	// TODO asMoney
	return sign + integer + decimal
}

func encodeInteger(num string, numMap map[string]string, si []string, index int) string {
	if len(num) <= 4 {
		var result string
		for i := range num {
			result += (numMap[string(num[i])] + si[len(num)-1-i])
		}
		if index >= 4 {
			result = strings.Replace(result, "个", si[index], -1)
		} else {
			result = strings.Replace(result, "个", "", -1)
		}
		return result
	}
	return encodeInteger(num[:len(num)-4], numMap, si, index+1) + encodeInteger(num[len(num)-4:], numMap, si, index)
}

func encodeDecimal(decimal string, numMap map[string]string) string {
	if decimal == "" {
		return decimal
	}

	result := "点"
	decimal = decimal[1:]
	for i := range decimal {
		result += numMap[string(decimal[i])]
	}
	return result
}
