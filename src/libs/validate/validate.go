package validate

import (
	"fmt"
	"reflect"
	"regexp"
)

// 一些参考的正则匹配表达式
const (
	REmail          string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	RCreditCard     string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
	RISBN10         string = "^(?:[0-9]{9}X|[0-9]{10})$"
	RISBN13         string = "^(?:[0-9]{13})$"
	RUUID3          string = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	RUUID4          string = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	RUUID5          string = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	RUUID           string = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	RAlpha          string = "^[a-zA-Z]+$"
	RAlphanumeric   string = "^[a-zA-Z0-9]+$"
	RNumeric        string = "^[0-9]+$"
	RInt            string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	RFloat          string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	RHexadecimal    string = "^[0-9a-fA-F]+$"
	RHexcolor       string = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	RRGBcolor       string = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	RASCII          string = "^[\x00-\x7F]+$"
	RMultibyte      string = "[^\x00-\x7F]"
	RFullWidth      string = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	RHalfWidth      string = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	RBase64         string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	RPrintableASCII string = "^[\x20-\x7E]+$"
	RDataURI        string = "^data:.+\\/(.+);base64$"
	RLatitude       string = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	RLongitude      string = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	RDNSName        string = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	RIP             string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	RURLSchema      string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	RURLUsername    string = `(\S+(:\S*)?@)`
	RURLPath        string = `((\/|\?|#)[^\s]*)`
	RURLPort        string = `(:(\d{1,5}))`
	RURLIP          string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	RURLSubdomain   string = `((www\.)|([a-zA-Z0-9]([-\.][-\._a-zA-Z0-9]+)*))`
	RURL            string = `^` + RURLSchema + `?` + RURLUsername + `?` + `((` + RURLIP + `|(\[` + RIP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + RURLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + RURLPort + `?` + RURLPath + `?$`
	RSSN            string = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	RWinPath        string = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	RUnixPath       string = `^(/[^/\x00]*)+/?$`
	RSemver         string = "^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"
)

// Validate 验证struct类型的数据
// 两个可选tag:
// 1. validate(正则表达式)
// 2. validate-message(验证未通过的默认提示)
func Validate(v interface{}) error {
	var value = reflect.ValueOf(v)

	if value.Kind() != reflect.Struct {
		return fmt.Errorf("function only accepts structs; got %s", value.Kind())
	}

	for i := 0; i < value.NumField(); i++ {
		var vField = value.Field(i)
		var tField = value.Type().Field(i)

		var tag, found = tField.Tag.Lookup("validate")
		if !found {
			continue
		}

		if tag == "-" || tag == "" {
			continue
		}

		var valStr = fmt.Sprintf("%v", vField.Interface())
		var r, err = regexp.Compile(tag)
		if err != nil {
			return err
		}

		if !r.Match([]byte(valStr)) {
			var tag1, found = tField.Tag.Lookup("validate-message")
			if found {
				return fmt.Errorf(tag1)
			}
			return fmt.Errorf("regex is: %s, value is: %s", tag, valStr)
		}
	}

	return nil
}

// Validate1 (废弃)验证struct类型的数据
// 两个可选tag:
// 1. validate(正则表达式)
// 2. validate-message(验证未通过的默认提示)
func Validate1(v interface{}) error {
	var t = reflect.TypeOf(v)

	if t.Kind() != reflect.Struct {
		return fmt.Errorf(
			"kind value is: %s(%d), but need: %s(%d)",
			t.Kind().String(), t.Kind(),
			reflect.Struct.String(), reflect.Struct,
		)
	}

	var val = reflect.ValueOf(v)

	for i := 0; i < t.NumField(); i++ {
		var field = t.Field(i)

		var tag, found = field.Tag.Lookup("validate")
		if !found {
			continue
		}

		if tag == "-" || tag == "" {
			continue
		}

		var valStr = fmt.Sprintf("%v", val.Field(i).Interface())
		var r, err = regexp.Compile(tag)
		if err != nil {
			return err
		}

		if !r.Match([]byte(valStr)) {
			var tag1, found = field.Tag.Lookup("validate-message")
			if found {
				return fmt.Errorf(tag1)
			}
			return fmt.Errorf("regex is: %s, value is: %s", tag, valStr)
		}
	}

	return nil
}
