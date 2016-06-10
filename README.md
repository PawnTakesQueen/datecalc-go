datecalc-go
========

datecalc is created by Vi Grey (https://vigrey.com) <vi@vigrey.com> and is licensed under a BSD 2-Clause License. Read LICENSE.txt for more license text.

Go package to calculate the day of the week of any date

####Dependencies
* Go

####Using the Package
func Date(year int, month, date int8, calType string) (day string, err error)

To calculate the day of the week for any date, use *datecalc.Date(y, m, d, t)* where y is the full year (a negative integer for BC years), m is the month number, d is the day number, and type is the calendar type.  The calendar types you have to chose from are:
* English
* Roman
* Gregorian
* Julian
* CE

English is the calendar system the English speaking western countries are using.  This is a system where the calendar was under the Julian system until 1752, when it switched to the Gregorian Calendar, skipping  September 3rd and going straight to September 14th to offset for the differences in the calendar systems on how they incorporated leap years.

Roman is the calendar system the Roman Empire used.  This was a system where the calendar was under the Julian system until 1582, when it switched to the Gregorian Calendar, skipping October 5th and going straight to October 15th to offset for the differences in the calendar systems on how they incorporated leap yars.

####Example Uses
```
package main

import (
    "datecalc"
    "fmt"
)

func main() {
    fmt.Println(datecalc.Date(2014, 3, 14, "English"))
    fmt.Println(datecalc.Date(2014, 3, 14, "Roman"))
    fmt.Println(datecalc.Date(-2014, 3, 14, "English"))
    fmt.Println(datecalc.Date(-2014, 3, 14, "Julian"))
    fmt.Println(datecalc.Date(2014, 3, 14, "Julian"))
}
```
prints out:
```
Friday
Friday
Wedneday
Wedneday
Thursday
```
