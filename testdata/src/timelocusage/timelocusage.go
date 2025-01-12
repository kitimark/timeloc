// Copyright (c) Wasutan Kitijerapat.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package timelocusage

import "time"

func setUtcTz(t time.Time) time.Time {
	return t.In(time.UTC)
}

func setAnotherTz(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return t.In(loc)
}

func parseDateInLocation(date string) time.Time {
	t, _ := time.ParseInLocation(time.DateOnly, date, time.UTC)
	return t
}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02")
}

func nestedFormatTime(t time.Time) string {
	return formatTime(t)
}

func nestedFormatTime2(t time.Time) string {
	return nestedFormatTime(t)
}

func nestedFormatTime3(t time.Time) string {
	return nestedFormatTime2(t)
}

func formatAllMethods(t time.Time) {
	_ = t.Format("2006-01-02")
	_ = t.AppendFormat(nil, "2006-01-02")
	_, _, _ = t.Clock()
	_, _, _ = t.Date()
	_ = t.Day()
	_ = t.Hour()
	_, _ = t.ISOWeek()
	_ = t.Minute()
	_ = t.Month()
	_ = t.Second()
	_ = t.Weekday()
	_ = t.Year()
	_ = t.YearDay()
}

func nestedFormatAllMethods(t time.Time) {
	formatAllMethods(t)
}

func nestedFormatAllMethods2(t time.Time) {
	nestedFormatAllMethods(t)
}

func nestedFormatAllMethods3(t time.Time) {
	nestedFormatAllMethods2(t)
}

func NormalUsage() {
	t := time.Now()
	t = t.In(time.UTC)

	_ = t.Format("2006-01-02")
	_ = t.AppendFormat(nil, "2006-01-02")
	_, _, _ = t.Clock()
	_, _, _ = t.Date()
	_ = t.Day()
	_ = t.Hour()
	_, _ = t.ISOWeek()
	_ = t.Minute()
	_ = t.Month()
	_ = t.Second()
	_ = t.Weekday()
	_ = t.Year()
	_ = t.YearDay()

	formatAllMethods(t)
	formatTime(t)
	nestedFormatAllMethods(t)
	nestedFormatTime(t)
	nestedFormatAllMethods2(t)
	nestedFormatTime2(t)
	nestedFormatAllMethods3(t)
	nestedFormatTime3(t)
}

func AdjustTzByFunc() {
	t := time.Now()
	t = setUtcTz(t)

	_ = t.Format("2006-01-02")
	_ = t.AppendFormat(nil, "2006-01-02")
	_, _, _ = t.Clock()
	_, _, _ = t.Date()
	_ = t.Day()
	_ = t.Hour()
	_, _ = t.ISOWeek()
	_ = t.Minute()
	_ = t.Month()
	_ = t.Second()
	_ = t.Weekday()
	_ = t.Year()
	_ = t.YearDay()

	formatAllMethods(t)
	formatTime(t)
	nestedFormatAllMethods(t)
	nestedFormatTime(t)
	nestedFormatAllMethods2(t)
	nestedFormatTime2(t)
	nestedFormatAllMethods3(t)
	nestedFormatTime3(t)
}

func AdjustAnotherTzByFunc() {
	t := time.Now()
	t = setAnotherTz(t)

	_ = t.Format("2006-01-02")
	_ = t.AppendFormat(nil, "2006-01-02")
	_, _, _ = t.Clock()
	_, _, _ = t.Date()
	_ = t.Day()
	_ = t.Hour()
	_, _ = t.ISOWeek()
	_ = t.Minute()
	_ = t.Month()
	_ = t.Second()
	_ = t.Weekday()
	_ = t.Year()
	_ = t.YearDay()

	formatAllMethods(t)
	formatTime(t)
	nestedFormatAllMethods(t)
	nestedFormatTime(t)
	nestedFormatAllMethods2(t)
	nestedFormatTime2(t)
	nestedFormatAllMethods3(t)
	nestedFormatTime3(t)
}

func ParseInLocationUsage() {
	t, _ := time.ParseInLocation(time.DateOnly, "2025-01-01", time.UTC)

	_ = t.Format("2006-01-02")
	_ = t.AppendFormat(nil, "2006-01-02")
	_, _, _ = t.Clock()
	_, _, _ = t.Date()
	_ = t.Day()
	_ = t.Hour()
	_, _ = t.ISOWeek()
	_ = t.Minute()
	_ = t.Month()
	_ = t.Second()
	_ = t.Weekday()
	_ = t.Year()
	_ = t.YearDay()

	formatAllMethods(t)
	formatTime(t)
	nestedFormatAllMethods(t)
	nestedFormatTime(t)
	nestedFormatAllMethods2(t)
	nestedFormatTime2(t)
	nestedFormatAllMethods3(t)
	nestedFormatTime3(t)
}

func ParseLocationFuncUsage() {
	t := parseDateInLocation("2025-01-01")

	_ = t.Format("2006-01-02")
	_ = t.AppendFormat(nil, "2006-01-02")
	_, _, _ = t.Clock()
	_, _, _ = t.Date()
	_ = t.Day()
	_ = t.Hour()
	_, _ = t.ISOWeek()
	_ = t.Minute()
	_ = t.Month()
	_ = t.Second()
	_ = t.Weekday()
	_ = t.Year()
	_ = t.YearDay()

	formatAllMethods(t)
	formatTime(t)
	nestedFormatAllMethods(t)
	nestedFormatTime(t)
	nestedFormatAllMethods2(t)
	nestedFormatTime2(t)
	nestedFormatAllMethods3(t)
	nestedFormatTime3(t)
}

func TimeParseUsage() {
	t, _ := time.Parse(time.DateOnly, "2025-01-01")
	t.In(time.UTC)

	_ = t.Format("2006-01-02")
	_ = t.AppendFormat(nil, "2006-01-02")
	_, _, _ = t.Clock()
	_, _, _ = t.Date()
	_ = t.Day()
	_ = t.Hour()
	_, _ = t.ISOWeek()
	_ = t.Minute()
	_ = t.Month()
	_ = t.Second()
	_ = t.Weekday()
	_ = t.Year()
	_ = t.YearDay()

	formatAllMethods(t)
	formatTime(t)
	nestedFormatAllMethods(t)
	nestedFormatTime(t)
	nestedFormatAllMethods2(t)
	nestedFormatTime2(t)
	nestedFormatAllMethods3(t)
	nestedFormatTime3(t)
}

func badUsageInSubFunc() {
	t := time.Now()

	_ = t.Format("2006-01-02")            // want "\\(t time.Time\\).Format called on t before setting time.Location"
	_ = t.AppendFormat(nil, "2006-01-02") // want "\\(t time.Time\\).AppendFormat called on t before setting time.Location"
	_, _, _ = t.Clock()                   // want "\\(t time.Time\\).Clock called on t before setting time.Location"
	_, _, _ = t.Date()                    // want "\\(t time.Time\\).Date called on t before setting time.Location"
	_ = t.Day()                           // want "\\(t time.Time\\).Day called on t before setting time.Location"
	_ = t.Hour()                          // want "\\(t time.Time\\).Hour called on t before setting time.Location"
	_, _ = t.ISOWeek()                    // want "\\(t time.Time\\).ISOWeek called on t before setting time.Location"
	_ = t.Minute()                        // want "\\(t time.Time\\).Minute called on t before setting time.Location"
	_ = t.Month()                         // want "\\(t time.Time\\).Month called on t before setting time.Location"
	_ = t.Second()                        // want "\\(t time.Time\\).Second called on t before setting time.Location"
	_ = t.Weekday()                       // want "\\(t time.Time\\).Weekday called on t before setting time.Location"
	_ = t.Year()                          // want "\\(t time.Time\\).Year called on t before setting time.Location"
	_ = t.YearDay()                       // want "\\(t time.Time\\).YearDay called on t before setting time.Location"

	formatAllMethods(t)        // want "passing time.Time value without location set to function that may use location-dependent methods"
	formatTime(t)              // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods(t)  // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime(t)        // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods2(t) // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime2(t)       // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods3(t) // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime3(t)       // want "passing time.Time value without location set to function that may use location-dependent methods"
}

func BadUsage() {
	t := time.Now()

	_ = t.Format("2006-01-02")            // want "\\(t time.Time\\).Format called on t before setting time.Location"
	_ = t.AppendFormat(nil, "2006-01-02") // want "\\(t time.Time\\).AppendFormat called on t before setting time.Location"
	_, _, _ = t.Clock()                   // want "\\(t time.Time\\).Clock called on t before setting time.Location"
	_, _, _ = t.Date()                    // want "\\(t time.Time\\).Date called on t before setting time.Location"
	_ = t.Day()                           // want "\\(t time.Time\\).Day called on t before setting time.Location"
	_ = t.Hour()                          // want "\\(t time.Time\\).Hour called on t before setting time.Location"
	_, _ = t.ISOWeek()                    // want "\\(t time.Time\\).ISOWeek called on t before setting time.Location"
	_ = t.Minute()                        // want "\\(t time.Time\\).Minute called on t before setting time.Location"
	_ = t.Month()                         // want "\\(t time.Time\\).Month called on t before setting time.Location"
	_ = t.Second()                        // want "\\(t time.Time\\).Second called on t before setting time.Location"
	_ = t.Weekday()                       // want "\\(t time.Time\\).Weekday called on t before setting time.Location"
	_ = t.Year()                          // want "\\(t time.Time\\).Year called on t before setting time.Location"
	_ = t.YearDay()                       // want "\\(t time.Time\\).YearDay called on t before setting time.Location"

	formatAllMethods(t)        // want "passing time.Time value without location set to function that may use location-dependent methods"
	formatTime(t)              // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods(t)  // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime(t)        // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods2(t) // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime2(t)       // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods3(t) // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime3(t)       // want "passing time.Time value without location set to function that may use location-dependent methods"

	badUsageInSubFunc()
}

// TimeParseBadUsage
// even time.Parse is set default location automatically related from server instance timezone.
// we wouldn't trust default timezone of server instance. we aim to correct timezone in our codebase only.
func TimeParseBadUsage() {
	t, _ := time.Parse(time.DateOnly, "2025-01-01")

	_ = t.Format("2006-01-02")            // want "\\(t time.Time\\).Format called on t before setting time.Location"
	_ = t.AppendFormat(nil, "2006-01-02") // want "\\(t time.Time\\).AppendFormat called on t before setting time.Location"
	_, _, _ = t.Clock()                   // want "\\(t time.Time\\).Clock called on t before setting time.Location"
	_, _, _ = t.Date()                    // want "\\(t time.Time\\).Date called on t before setting time.Location"
	_ = t.Day()                           // want "\\(t time.Time\\).Day called on t before setting time.Location"
	_ = t.Hour()                          // want "\\(t time.Time\\).Hour called on t before setting time.Location"
	_, _ = t.ISOWeek()                    // want "\\(t time.Time\\).ISOWeek called on t before setting time.Location"
	_ = t.Minute()                        // want "\\(t time.Time\\).Minute called on t before setting time.Location"
	_ = t.Month()                         // want "\\(t time.Time\\).Month called on t before setting time.Location"
	_ = t.Second()                        // want "\\(t time.Time\\).Second called on t before setting time.Location"
	_ = t.Weekday()                       // want "\\(t time.Time\\).Weekday called on t before setting time.Location"
	_ = t.Year()                          // want "\\(t time.Time\\).Year called on t before setting time.Location"
	_ = t.YearDay()                       // want "\\(t time.Time\\).YearDay called on t before setting time.Location"

	formatAllMethods(t)        // want "passing time.Time value without location set to function that may use location-dependent methods"
	formatTime(t)              // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods(t)  // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime(t)        // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods2(t) // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime2(t)       // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods3(t) // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime3(t)       // want "passing time.Time value without location set to function that may use location-dependent methods"
}
