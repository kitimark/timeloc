// Copyright (c) Wasutan Kitijerapat.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package aliastimepackage

import aliastime "time"

func setUtcTz(t aliastime.Time) aliastime.Time {
	return t.In(aliastime.UTC)
}

func setAnotherTz(t aliastime.Time) aliastime.Time {
	loc, _ := aliastime.LoadLocation("Asia/Bangkok")
	return t.In(loc)
}

func parseDateInLocation(date string) aliastime.Time {
	t, _ := aliastime.ParseInLocation(aliastime.DateOnly, date, aliastime.UTC)
	return t
}

func formatTime(t aliastime.Time) string {
	return t.Format("2006-01-02")
}

func nestedFormatTime(t aliastime.Time) string {
	return formatTime(t)
}

func nestedFormatTime2(t aliastime.Time) string {
	return nestedFormatTime(t)
}

func nestedFormatTime3(t aliastime.Time) string {
	return nestedFormatTime2(t)
}

func formatAllMethods(t aliastime.Time) {
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

func nestedFormatAllMethods(t aliastime.Time) {
	formatAllMethods(t)
}

func nestedFormatAllMethods2(t aliastime.Time) {
	nestedFormatAllMethods(t)
}

func nestedFormatAllMethods3(t aliastime.Time) {
	nestedFormatAllMethods2(t)
}

func NormalUsage() {
	t := aliastime.Now()
	t = t.In(aliastime.UTC)

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
	t := aliastime.Now()
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
	t := aliastime.Now()
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
	t, _ := aliastime.ParseInLocation(aliastime.DateOnly, "2025-01-01", aliastime.UTC)

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
	t, _ := aliastime.Parse(aliastime.DateOnly, "2025-01-01")
	t.In(aliastime.UTC)

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
	t := aliastime.Date(2025, 1, 1, 0, 0, 0, 0, aliastime.UTC)

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
	t := aliastime.Now()

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

func BadUsage() {
	t := aliastime.Now()

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
	t, _ := aliastime.Parse(aliastime.DateOnly, "2025-01-01")

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
	_, _ = aliastime.ParseInLocation(aliastime.DateOnly, "2025-01-01", aliastime.Local) // want "time.Local usage is not allowed as it relies on system timezone"

	_ = aliastime.Date(2025, 1, 1, 0, 0, 0, 0, aliastime.Local) // want "time.Local usage is not allowed as it relies on system timezone"

	t := aliastime.Now()
	_ = t.In(aliastime.Local) // want "time.Local usage is not allowed as it relies on system timezone"
}

func nestedSubFunctionTimeLocalBadUsage() {
	subFunctionTimeLocalBadUsage()
}

func TimeLocalBadUsage() {
	_, _ = aliastime.ParseInLocation(aliastime.DateOnly, "2025-01-01", aliastime.Local) // want "time.Local usage is not allowed as it relies on system timezone"

	_ = aliastime.Date(2025, 1, 1, 0, 0, 0, 0, aliastime.Local) // want "time.Local usage is not allowed as it relies on system timezone"

	t := aliastime.Now()
	_ = t.In(aliastime.Local) // want "time.Local usage is not allowed as it relies on system timezone"

	subFunctionTimeLocalBadUsage()
	nestedSubFunctionTimeLocalBadUsage()
}
