package main

import (
	"fmt"
	"time"
)

func doSomething() {
	time.Sleep(250 * time.Millisecond)
}

func main() {
	now := time.Now()

	timeText := fmt.Sprintf(
		"now: %v\n"+
			"year: %v\n"+
			"month: %v\n"+
			"day: %v\n"+
			"hour: %v\n"+
			"minute: %v\n"+
			"second: %v\n"+
			"nanosecond: %v\n"+
			"weekday: %v\n"+
			"year day: %v\n"+
			"location: %v\n"+
			"utc: %v\n"+
			"local: %v\n"+
			"formatted date: %v\n"+
			"formatted time: %v\n"+
			"formatted datetime: %v\n"+
			"rfc3339: %v\n"+
			"unix: %v\n"+
			"unix milli: %v\n"+
			"unix micro: %v\n"+
			"unix nano: %v\n"+
			"is zero: %v\n"+
			"after one hour: %v\n"+
			"after one day: %v\n"+
			"after one month: %v\n"+
			"after one year: %v\n"+
			"truncated hour: %v\n"+
			"rounded hour: %v\n",

		now,
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond(),
		now.Weekday(),
		now.YearDay(),
		now.Location(),
		now.UTC(),
		now.Local(),
		now.Format("2006-01-02"),
		now.Format("15:04:05"),
		now.Format("2006-01-02 15:04:05"),
		now.Format(time.RFC3339),
		now.Unix(),
		now.UnixMilli(),
		now.UnixMicro(),
		now.UnixNano(),
		now.IsZero(),
		now.Add(time.Hour),
		now.AddDate(0, 0, 1),
		now.AddDate(0, 1, 0),
		now.AddDate(1, 0, 0),
		now.Truncate(time.Hour),
		now.Round(time.Hour),
	)
	fmt.Println(timeText)

	var t time.Time
	fmt.Println(t.IsZero()) // true

	var d time.Duration = 5 * time.Second
	fmt.Printf("Time Duration: %v\n", d)

	/*
		location, err := time.LoadLocation("Europe/Berlin")
		if err != nil {
			panic(err)
		}

		fmt.Println(location)

		Europe/Berlin
		America/New_York
		America/Los_Angeles
		Asia/Tehran
		Asia/Tokyo
		UTC
	*/

	start := time.Now()

	// Code to measure.
	doSomething()

	elapsed := time.Since(start)

	fmt.Println("Elapsed:", elapsed)

	deadline := time.Now().Add(10 * time.Second)

	remaining := time.Until(deadline)

	fmt.Println("Remaining:", remaining)
}

/*
time.January
time.February
time.March
time.April
time.May
time.June
time.July
time.August
time.September
time.October
time.November
time.December

time.Sunday
time.Monday
time.Tuesday
time.Wednesday
time.Thursday
time.Friday
time.Saturday
*/

/*
Creating a Specific Date with time.Date
Use time.Date when you know the individual components:

go
t := time.Date(
	2026,
	time.September,
	27,
	12,
	30,
	0,
	0,
	time.UTC,
)

fmt.Println(t)
Signature:

time.Date(
	year int,
	month time.Month,
	day int,
	hour int,
	min int,
	sec int,
	nsec int,
	loc *time.Location,
) time.Time
*/

/*
Formatted Time:

time.Layout
time.ANSIC
time.UnixDate
time.RubyDate
time.RFC822
time.RFC822Z
time.RFC850
time.RFC1123
time.RFC1123Z
time.RFC3339
time.RFC3339Nano
time.Kitchen
time.Stamp
time.StampMilli
time.StampMicro
time.StampNano
time.DateTime
time.DateOnly
time.TimeOnly

t := time.Date(
	2026,
	time.September,
	27,
	14,
	5,
	9,
	123456789,
	time.UTC,
)

fmt.Println(t.Format(time.RFC3339))
fmt.Println(t.Format(time.DateOnly))
fmt.Println(t.Format(time.TimeOnly))
fmt.Println(t.Format(time.DateTime))
*/

/*
Parsing Strings into Time
Use time.Parse to convert a string into time.Time.


layout := "2006-01-02"

t, err := time.Parse(layout, "2026-09-27")
if err != nil {
	fmt.Println("Parse error:", err)
	return
}

fmt.Println(t)
Important: if the input contains no time zone, time.Parse interprets it as UTC.


t, err := time.Parse(
	"2006-01-02 15:04:05",
	"2026-09-27 14:30:00",
)

if err != nil {
	panic(err)
}

fmt.Println(t.Location()) // UTC
Parsing RFC3339
For APIs and JSON, RFC3339 is usually the best choice:


value := "2026-09-27T14:30:00+02:00"

t, err := time.Parse(time.RFC3339, value)
if err != nil {
	panic(err)
}

fmt.Println(t)
Example UTC value:


value := "2026-09-27T12:30:00Z"

t, err := time.Parse(time.RFC3339, value)
if err != nil {
	panic(err)
}
Parsing custom formats

layout := "02/01/2006 15:04"

t, err := time.Parse(layout, "27/09/2026 14:30")
if err != nil {
	panic(err)
}

fmt.Println(t)
Remember:

text
02/01/2006
means:

text
day/month/year
because the layout is based on the reference date.

ParseInLocation
Use ParseInLocation when the input has no time zone and should be interpreted in a specific location.


berlin, err := time.LoadLocation("Europe/Berlin")
if err != nil {
	panic(err)
}

t, err := time.ParseInLocation(
	"2006-01-02 15:04:05",
	"2026-09-27 14:30:00",
	berlin,
)

if err != nil {
	panic(err)
}

fmt.Println(t)
Difference:


time.Parse(...)
without a zone uses UTC.


time.ParseInLocation(..., berlin)
uses Berlin time.

This distinction is important for appointments, business hours, and user-entered local times.

Always check parsing errors
Avoid:


t, _ := time.Parse(time.RFC3339, input)
Prefer:


t, err := time.Parse(time.RFC3339, input)
if err != nil {
	return err
}
An invalid date can otherwise silently become a zero-value time in your program logic.

Parsing durations
Use time.ParseDuration:


d, err := time.ParseDuration("2h30m")
if err != nil {
	panic(err)
}

fmt.Println(d)
Supported units:

text
ns   nanoseconds
us   microseconds
µs   microseconds
ms   milliseconds
s    seconds
m    minutes
h    hours
Examples:


time.ParseDuration("500ms")
time.ParseDuration("1.5s")
time.ParseDuration("2h45m")
time.ParseDuration("-10m")
Complete example:


inputs := []string{
	"500ms",
	"1.5s",
	"2h45m",
	"-10m",
}

for _, input := range inputs {
	d, err := time.ParseDuration(input)
	if err != nil {
		fmt.Println("Error:", err)
		continue
	}

	fmt.Println(input, "=", d)
}
A duration cannot directly parse "2 days" because the package only supports units through hours.
*/

/*
duration.Hours()        // float64
duration.Minutes()      // float64
duration.Seconds()      // float64
duration.Milliseconds() // int64
duration.Microseconds() // int64
duration.Nanoseconds()  // int64
*/
