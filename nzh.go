package gonzh

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	// regular expression for scientific notification
	sciRegEx = regexp.MustCompile(`^(?P<sign>[+-])?0*(?P<integer>\d+)(\.(?P<decimal>\d+))?e((?P<esign>[+-])?(?P<exp>\d+))$`)
	// regular expression for ordinary number
	numRegEx   = regexp.MustCompile(`^(?P<sign>[+-])?0*(?P<integer>\d+)(\.(?P<decimal>\d+))?$`)
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
	if len(parts) == 0 {
		return nil
	}

	e, ok := parts["exp"]
	if ok {
		exp, _ := strconv.Atoi(e)
		return sciNotation{
			minus:   parts["sign"] == "-",
			eminus:  parts["esign"] == "-",
			integer: parts["integer"],
			decimal: strings.TrimRight(parts["decimal"], "0"),
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
				decimal: zeros + n.integer + n.decimal,
			}
		} else {
			result = number{
				minus:   n.minus,
				integer: n.integer[:-diff],
				decimal: n.integer[-diff:] + n.decimal,
			}
		}
	} else {
		diff := n.exp - len(n.decimal)
		if diff >= 0 {
			var zeros string
			for i := 0; i < diff; i++ {
				zeros += "0"
			}
			integer := n.integer
			for _, i := range n.decimal {
				integer += string(i)
			}
			result = number{
				minus:   n.minus,
				integer: integer + zeros,
			}
		} else {
			result = number{
				minus:   n.minus,
				integer: n.integer + n.decimal[:n.exp],
				decimal: n.decimal[n.exp:],
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
	if integer == "" {
		integer = "零"
	}
	decimal := encodeDecimal(n.decimal, num)
	var point string
	if asMoney {
		integer, decimal = treatAsMoney(integer, decimal, isCapital)
	} else if len(decimal) != 0 {
		point = "点"
	}
	return sign + integer + point + decimal
}

func treatAsMoney(integer, decimal string, capital bool) (string, string) {
	unit := map[bool]string{
		false: "元",
		true:  "圆",
	}
	integer += unit[capital]
	if utf8.RuneCountInString(decimal) == 0 {
		integer += "整"
	} else {
		if utf8.RuneCountInString(decimal) > 1 {
			decimal = string([]rune(decimal)[0]) + "角" + string([]rune(decimal)[1]) + "分"
		} else {
			decimal = string([]rune(decimal)[0]) + "角"
		}
	}
	return integer, decimal
}

func encodeInteger(num string, numMap map[string]string, si []string, index int) string {
	if len(num) <= 4 {
		var result string
		for i := range num {
			result += (numMap[string(num[i])] + si[len(num)-1-i])
		}
		result = dealWithIntegerZeros(result)
		if index >= 4 {
			result = strings.Replace(result, "个", si[index], -1)
		}
		if utf8.RuneCountInString(result) == 2 && string([]rune(result)[0]) == "零" {
			result = ""
		}
		return strings.Replace(result, "个", "", -1)
	}
	return encodeInteger(num[:len(num)-4], numMap, si, index+1) + encodeInteger(num[len(num)-4:], numMap, si, index)
}

func dealWithIntegerZeros(chn string) string {
	chn = strings.Replace(chn, "零个", "个", -1)
	chn = strings.Replace(chn, "零十", "零", -1)
	chn = strings.Replace(chn, "零百", "零", -1)
	chn = strings.Replace(chn, "零千", "零", -1)
	chn = strings.Replace(chn, "零万", "零", -1)
	chn = strings.Replace(chn, "零亿", "零", -1)
	chn = strings.Replace(chn, "零兆", "零", -1)
	chn = strings.Replace(chn, "零京", "零", -1)
	chn = strings.Replace(chn, "零垓", "零", -1)
	chn = strings.Replace(chn, "零秭", "零", -1)
	chn = strings.Replace(chn, "零穰", "零", -1)
	chn = strings.Replace(chn, "零沟", "零", -1)
	chn = strings.Replace(chn, "零涧", "零", -1)
	chn = strings.Replace(chn, "零正", "零", -1)
	chn = strings.Replace(chn, "零载", "零", -1)
	var (
		filtered string
		found    bool
	)
	for _, c := range chn {
		if string(c) == "零" {
			if found {
				continue
			} else {
				found = true
				filtered += string(c)
			}
		} else {
			found = false
			filtered += string(c)
		}
	}
	return filtered
}

func encodeDecimal(decimal string, numMap map[string]string) string {
	if decimal == "" {
		return decimal
	}
	decimal = strings.TrimRight(decimal, "0")
	var result string
	for i := range decimal {
		result += numMap[string(decimal[i])]
	}
	return result
}
