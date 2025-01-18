// Copyright (c) Wasutan Kitijerapat.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package timelocusage

import "time"

var bkkLoc *time.Location

func init() {
	bkkLoc, _ = time.LoadLocation("Asia/Bangkok")
}

func setUtcTz(t time.Time) time.Time {
	return t.In(time.UTC)
}

func setAnotherTz(t time.Time) time.Time {
	return t.In(bkkLoc)
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

func TimeDateUsage() {
	t := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

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

func TimeMethodChainingUsage() {
	_ = time.Now().In(bkkLoc).Format("2006-01-02")
	_ = time.Now().In(bkkLoc).AppendFormat(nil, "2006-01-02")
	_, _, _ = time.Now().In(bkkLoc).Clock()
	_, _, _ = time.Now().In(bkkLoc).Date()
	_ = time.Now().In(bkkLoc).Day()
	_ = time.Now().In(bkkLoc).Hour()
	_, _ = time.Now().In(bkkLoc).ISOWeek()
	_ = time.Now().In(bkkLoc).Minute()
	_ = time.Now().In(bkkLoc).Month()
	_ = time.Now().In(bkkLoc).Second()
	_ = time.Now().In(bkkLoc).Weekday()
	_ = time.Now().In(bkkLoc).Year()
	_ = time.Now().In(bkkLoc).YearDay()

	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).Format("2006-01-02")
	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).AppendFormat(nil, "2006-01-02")
	_, _, _ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).Clock()
	_, _, _ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).Date()
	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).Day()
	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).Hour()
	_, _ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).ISOWeek()
	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).Minute()
	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).Month()
	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).Second()
	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).Weekday()
	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).Year()
	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, bkkLoc).YearDay()
}

func badUsageInSubFunc() {
	t := time.Now()

	_ = t.Format("2006-01-02")            // want "\\(t time.Time\\).Format called before setting time.Location"
	_ = t.AppendFormat(nil, "2006-01-02") // want "\\(t time.Time\\).AppendFormat called before setting time.Location"
	_, _, _ = t.Clock()                   // want "\\(t time.Time\\).Clock called before setting time.Location"
	_, _, _ = t.Date()                    // want "\\(t time.Time\\).Date called before setting time.Location"
	_ = t.Day()                           // want "\\(t time.Time\\).Day called before setting time.Location"
	_ = t.Hour()                          // want "\\(t time.Time\\).Hour called before setting time.Location"
	_, _ = t.ISOWeek()                    // want "\\(t time.Time\\).ISOWeek called before setting time.Location"
	_ = t.Minute()                        // want "\\(t time.Time\\).Minute called before setting time.Location"
	_ = t.Month()                         // want "\\(t time.Time\\).Month called before setting time.Location"
	_ = t.Second()                        // want "\\(t time.Time\\).Second called before setting time.Location"
	_ = t.Weekday()                       // want "\\(t time.Time\\).Weekday called before setting time.Location"
	_ = t.Year()                          // want "\\(t time.Time\\).Year called before setting time.Location"
	_ = t.YearDay()                       // want "\\(t time.Time\\).YearDay called before setting time.Location"

	_ = time.Now().Format("2006-01-02")            // want "\\(t time.Time\\).Format called before setting time.Location"
	_ = time.Now().AppendFormat(nil, "2006-01-02") // want "\\(t time.Time\\).AppendFormat called before setting time.Location"
	_, _, _ = time.Now().Clock()                   // want "\\(t time.Time\\).Clock called before setting time.Location"
	_, _, _ = time.Now().Date()                    // want "\\(t time.Time\\).Date called before setting time.Location"
	_ = time.Now().Day()                           // want "\\(t time.Time\\).Day called before setting time.Location"
	_ = time.Now().Hour()                          // want "\\(t time.Time\\).Hour called before setting time.Location"
	_, _ = time.Now().ISOWeek()                    // want "\\(t time.Time\\).ISOWeek called before setting time.Location"
	_ = time.Now().Minute()                        // want "\\(t time.Time\\).Minute called before setting time.Location"
	_ = time.Now().Month()                         // want "\\(t time.Time\\).Month called before setting time.Location"
	_ = time.Now().Second()                        // want "\\(t time.Time\\).Second called before setting time.Location"
	_ = time.Now().Weekday()                       // want "\\(t time.Time\\).Weekday called before setting time.Location"
	_ = time.Now().Year()                          // want "\\(t time.Time\\).Year called before setting time.Location"
	_ = time.Now().YearDay()                       // want "\\(t time.Time\\).YearDay called before setting time.Location"

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

	_ = t.Format("2006-01-02")            // want "\\(t time.Time\\).Format called before setting time.Location"
	_ = t.AppendFormat(nil, "2006-01-02") // want "\\(t time.Time\\).AppendFormat called before setting time.Location"
	_, _, _ = t.Clock()                   // want "\\(t time.Time\\).Clock called before setting time.Location"
	_, _, _ = t.Date()                    // want "\\(t time.Time\\).Date called before setting time.Location"
	_ = t.Day()                           // want "\\(t time.Time\\).Day called before setting time.Location"
	_ = t.Hour()                          // want "\\(t time.Time\\).Hour called before setting time.Location"
	_, _ = t.ISOWeek()                    // want "\\(t time.Time\\).ISOWeek called before setting time.Location"
	_ = t.Minute()                        // want "\\(t time.Time\\).Minute called before setting time.Location"
	_ = t.Month()                         // want "\\(t time.Time\\).Month called before setting time.Location"
	_ = t.Second()                        // want "\\(t time.Time\\).Second called before setting time.Location"
	_ = t.Weekday()                       // want "\\(t time.Time\\).Weekday called before setting time.Location"
	_ = t.Year()                          // want "\\(t time.Time\\).Year called before setting time.Location"
	_ = t.YearDay()                       // want "\\(t time.Time\\).YearDay called before setting time.Location"

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

	_ = t.Format("2006-01-02")            // want "\\(t time.Time\\).Format called before setting time.Location"
	_ = t.AppendFormat(nil, "2006-01-02") // want "\\(t time.Time\\).AppendFormat called before setting time.Location"
	_, _, _ = t.Clock()                   // want "\\(t time.Time\\).Clock called before setting time.Location"
	_, _, _ = t.Date()                    // want "\\(t time.Time\\).Date called before setting time.Location"
	_ = t.Day()                           // want "\\(t time.Time\\).Day called before setting time.Location"
	_ = t.Hour()                          // want "\\(t time.Time\\).Hour called before setting time.Location"
	_, _ = t.ISOWeek()                    // want "\\(t time.Time\\).ISOWeek called before setting time.Location"
	_ = t.Minute()                        // want "\\(t time.Time\\).Minute called before setting time.Location"
	_ = t.Month()                         // want "\\(t time.Time\\).Month called before setting time.Location"
	_ = t.Second()                        // want "\\(t time.Time\\).Second called before setting time.Location"
	_ = t.Weekday()                       // want "\\(t time.Time\\).Weekday called before setting time.Location"
	_ = t.Year()                          // want "\\(t time.Time\\).Year called before setting time.Location"
	_ = t.YearDay()                       // want "\\(t time.Time\\).YearDay called before setting time.Location"

	formatAllMethods(t)        // want "passing time.Time value without location set to function that may use location-dependent methods"
	formatTime(t)              // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods(t)  // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime(t)        // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods2(t) // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime2(t)       // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatAllMethods3(t) // want "passing time.Time value without location set to function that may use location-dependent methods"
	nestedFormatTime3(t)       // want "passing time.Time value without location set to function that may use location-dependent methods"
}

func subFunctionTimeLocalBadUsage() {
	_, _ = time.ParseInLocation(time.DateOnly, "2025-01-01", time.Local) // want "time.Local usage is not allowed as it relies on system timezone"

	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local) // want "time.Local usage is not allowed as it relies on system timezone"

	t := time.Now()
	_ = t.In(time.Local) // want "time.Local usage is not allowed as it relies on system timezone"
}

func nestedSubFunctionTimeLocalBadUsage() {
	subFunctionTimeLocalBadUsage()
}

func TimeLocalBadUsage() {
	_, _ = time.ParseInLocation(time.DateOnly, "2025-01-01", time.Local) // want "time.Local usage is not allowed as it relies on system timezone"

	_ = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local) // want "time.Local usage is not allowed as it relies on system timezone"

	t := time.Now()
	_ = t.In(time.Local) // want "time.Local usage is not allowed as it relies on system timezone"

	subFunctionTimeLocalBadUsage()
	nestedSubFunctionTimeLocalBadUsage()
}

func BadUsageTimeMethodChainingUsage() {
	_ = time.Now().Format("2006-01-02")            // want "\\(t time.Time\\).Format called before setting time.Location"
	_ = time.Now().AppendFormat(nil, "2006-01-02") // want "\\(t time.Time\\).AppendFormat called before setting time.Location"
	_, _, _ = time.Now().Clock()                   // want "\\(t time.Time\\).Clock called before setting time.Location"
	_, _, _ = time.Now().Date()                    // want "\\(t time.Time\\).Date called before setting time.Location"
	_ = time.Now().Day()                           // want "\\(t time.Time\\).Day called before setting time.Location"
	_ = time.Now().Hour()                          // want "\\(t time.Time\\).Hour called before setting time.Location"
	_, _ = time.Now().ISOWeek()                    // want "\\(t time.Time\\).ISOWeek called before setting time.Location"
	_ = time.Now().Minute()                        // want "\\(t time.Time\\).Minute called before setting time.Location"
	_ = time.Now().Month()                         // want "\\(t time.Time\\).Month called before setting time.Location"
	_ = time.Now().Second()                        // want "\\(t time.Time\\).Second called before setting time.Location"
	_ = time.Now().Weekday()                       // want "\\(t time.Time\\).Weekday called before setting time.Location"
	_ = time.Now().Year()                          // want "\\(t time.Time\\).Year called before setting time.Location"
	_ = time.Now().YearDay()                       // want "\\(t time.Time\\).YearDay called before setting time.Location"
}
