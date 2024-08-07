// SiYuan - Refactor your thinking
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package util

import (
	"math"
	"text/template"
	"time"

	// edit_itmikno
	"strconv"
	"reflect"
	"fmt"


	"github.com/88250/go-humanize"
	"github.com/Masterminds/sprig/v3"
	"github.com/araddon/dateparse"
	"github.com/siyuan-note/logging"
	// "github.com/spf13/cast" // edit_itmikno
)

func BuiltInTemplateFuncs() (ret template.FuncMap) {
	ret = sprig.TxtFuncMap()
	ret["Weekday"] = Weekday
	ret["WeekdayCN"] = WeekdayCN
	ret["WeekdayCN2"] = WeekdayCN2
	ret["ISOWeek"] = ISOWeek
	ret["pow"] = pow
	ret["powf"] = powf
	ret["log"] = log
	ret["logf"] = logf
	ret["parseTime"] = parseTime
	ret["FormatFloat"] = FormatFloat
	return
}

// edit_itmikno
var floatType = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf("")
func getFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		 return i, nil
	case float32:
		 return float64(i), nil
	case int64:
		 return float64(i), nil
	case int32:
		 return float64(i), nil
	case int:
		 return float64(i), nil
	case uint64:
		 return float64(i), nil
	case uint32:
		 return float64(i), nil
	case uint:
		 return float64(i), nil
	case string:
		 return strconv.ParseFloat(i, 64)
	default:
		 v := reflect.ValueOf(unk)
		 v = reflect.Indirect(v)
		 if v.Type().ConvertibleTo(floatType) {
			  fv := v.Convert(floatType)
			  return fv.Float(), nil
		 } else if v.Type().ConvertibleTo(stringType) {
			  sv := v.Convert(stringType)
			  s := sv.String()
			  return strconv.ParseFloat(s, 64)
		 } else {
			  return math.NaN(), fmt.Errorf("Can't convert %v to float64", v.Type())
		 }
	}
}

func toFloat(x interface {}) float64 {
	f, _ := getFloat(x)
	return f
}

func pow(a, b interface{}) int64    { return int64(math.Pow(toFloat(a), toFloat(b))) }
func powf(a, b interface{}) float64 { return math.Pow(toFloat(a), toFloat(b)) }
func log(a, b interface{}) int64 {
	return int64(math.Log(toFloat(a)) / math.Log(toFloat(b)))
}
func logf(a, b interface{}) float64 { return math.Log(toFloat(a)) / math.Log(toFloat(b)) }

// func pow(a, b interface{}) int64    { return int64(math.Pow(cast.ToFloat64(a), cast.ToFloat64(b))) }
// func powf(a, b interface{}) float64 { return math.Pow(cast.ToFloat64(a), cast.ToFloat64(b)) }
// func log(a, b interface{}) int64 {
// 	return int64(math.Log(cast.ToFloat64(a)) / math.Log(cast.ToFloat64(b)))
// }
// func logf(a, b interface{}) float64 { return math.Log(cast.ToFloat64(a)) / math.Log(cast.ToFloat64(b)) }

func parseTime(dateStr string) time.Time {
	now := time.Now()
	retTime, err := dateparse.ParseIn(dateStr, now.Location())
	if nil != err {
		logging.LogWarnf("parse date [%s] failed [%s], return current time instead", dateStr, err)
		return now
	}
	return retTime
}

func FormatFloat(format string, n float64) string {
	return humanize.FormatFloat(format, n)
}
