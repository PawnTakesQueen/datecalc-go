/* Copyright (C) 2012-2014, PariahVi
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *     1. Redistributions of source code must retain the above copyright
 *        notice, this list of conditions and the following disclaimer.
 *     2. Redistributions in binary form must reproduce the above copyright
 *        notice, this list of conditions and the following disclaimer in the
 *        documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY AUTHOR AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL AUTHOR OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */
 
/* A Library to Calculate the Day of the Week of Any Date */

/* Version 1.0.3.4 */

package datecalc

import (
    "errors"
    "strconv"
    "strings"
)

var days []string = []string{
    "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
}
var calTypes []string = []string{
    "GREGORIAN", "CE", "JULIAN", "ENGLISH", "ROMAN",
}
var monthNames []string = []string{
    "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
}

// See if x is in range min, max
func inRange(x, min, max int8) bool {
    for i := min; i < max; i ++ {
        if i == x {
            return true
        }
    }
    return false
}

// Figure out if leap year for CE style dates.
func ceLeapYear(y int) bool {
    if y % 100 == 0 {
        if y % 400 == 0 {
            return true
        }
    } else if y % 4 == 0 {
        return true
    }
    return false
}
// Figure out if leap year for Julian style dates.
func julLeapYear(y int) bool {
    if y < 0 {
        y = (((y + 1) % 700) + 700) % 700
    }
    if y % 4 == 0 {
        return true
    }
    return false
}

// Figure out if year is a leap year for cal_type.
func isLeapYear(year int, calType string) bool {
    if calType == "GREGORIAN" {
        newYear := year
        if year < 0 {
            newYear = (((year + 1) % 400) + 400) % 400
        }
        return ceLeapYear(newYear)
    } else if calType == "CE" {
        return ceLeapYear(year)
    } else if calType == "JULIAN" {
        return julLeapYear(year)
    } else if calType == "ENGLISH" {
        if year >= 1800 {
            return ceLeapYear(year)
        } else {
            return julLeapYear(year)
        }
    } else if calType == "ROMAN" {
        if year >= 1700 {
            return ceLeapYear(year)
        } else {
            return julLeapYear(year)
        }
    }
    return false
}

// Check if the date (year, month, date) exists in cal_type.
func isRealDate(year int, month, date int8, calType string) bool {
    month30 := []int8{4, 6, 9, 11}
    if inRange(month, 1, 13) == false {
        return false
    }
    if date < 1 || date > 31 {
        return false
    }
    inTypes := false
    for _, x := range calTypes {
        if calType == x {
            inTypes = true
            break
        }
    }
    if !inTypes {
        return false
    }
    if year == 0 && calType != "CE" {
        return false
    } else {
        inMonths := false
        for _, x := range month30 {
            if month == x {
                inMonths = true
                break
            }
        }
        if inMonths && date > 30 {
            return false
        } else if month == 2 && !(date < 29 || (isLeapYear(year, calType) &&
                date == 29)) {
            return false
        }
    }
    if calType == "ENGLISH" {
        if month == 9 && date > 2 && date < 15 && year == 1752 {
            return false
        }
    } else if calType == "ROMAN" {
        if month == 10 && date > 4 && date < 16 && year == 1582 {
            return false
        }
    }
    return true
}

// Figures out value to add from last two digits of year.
func addxxYY(year int, calType string) int8 {
    newYear := ((year % 100) + 100) % 100
    if calType != "CE" && year < 0 {
        newYear = (((year + 1) % 100) + 100) % 100
    }
    return int8((((newYear / 12 + (((newYear % 12) + 12) % 12) +
	        ((((newYear % 12) + 12) % 12) / 4)) % 7) + 7) % 7)
}

// Returns value calculated from every digit of the year besides the last 2
// digits for CE style dates.
func ceAddYYxx(y int) int8 {
    YYxx := []int8{2, 0, 5, 3}
    return YYxx[(((y / 100) % 4) + 4) % 4]
}

// Returns value calculated from every digit of the year besides the last 2
// digits for Julian style dates.
func julAddYYxx(y int) int8 {
    return int8((((7 - y / 100) % 7) + 7) % 7)
}

// Figures out value to add from every digit of the year besides the last 2
// digits.
func addYYxx(year int, month, date int8, calType string) int8{
    if calType == "GREGORIAN" {
        newYear := year
        if year < 0 {
            newYear = (((year + 1) % 400) + 400) % 400
        newYear = ((newYear % 400) + 400) % 400
        return ceAddYYxx(newYear)
        }
    } else if calType == "CE" {
        newYear := ((year % 400) + 400) % 400
        return ceAddYYxx(newYear)
    } else if calType == "JULIAN" {
        newYear := year
        if year < 0 {
            newYear = (((year + 1) % 700) + 700) % 700
        }
        newYear = ((newYear % 700) + 700) % 700
        return julAddYYxx(newYear)
    } else if calType == "ENGLISH" {
        if year >= 1752 {
            if year == 1752 {
                if inRange(month, 9, 13) {
                    if month == 9 {
                        if date >= 14 {
                            return 0
                        } else {
                            return 4
                        }
                    } else {
                        return 0
                    }
                } else {
                    return 4
                }
            } else {
                newYear := ((year % 400) + 400) % 400
                return ceAddYYxx(newYear)
            }
        } else {
            newYear := year
            if year < 0 {
                newYear = (((year + 1) % 700) + 700) % 700
            newYear = ((newYear % 700) + 700) % 700
            return julAddYYxx(newYear)
            }
        }
    } else if calType == "ROMAN" {
        if year >= 1582 {
            if year == 1582 {
                if inRange(month, 10, 13) {
                    if month == 10 {
                        if date >= 15 {
                            return 3
                        } else {
                            return 6
                        }
                    } else {
                        return 3
                    }
                } else {
                    return 6
                }
            } else {
                newYear := ((year % 400) + 400) % 400
                return ceAddYYxx(newYear)
            }
        } else {
            newYear := ((year % 400) + 400) % 400
            if year < 0 {
                newYear = (((year + 1) % 700) + 700) % 700
            newYear = ((newYear % 700) + 700) % 700
            return julAddYYxx(newYear)
            }
        }
    }
    return 0
}

// Add value calculated from the year.
func addYear(year int, month, date int8, calType string) int8 {
    return addYYxx(year, month, date, calType) + addxxYY(year, calType)
}

// Add value for the month based on the year and cal_type.
func addMonth(year int, month int8, calType string) int8 {
    monthOffsetKey := []int8{0, 0, 0, 3, 5, 1, 3, 6, 2, 4, 0, 2}
    monthOffset := make([][]int8, 12)
    for i, value := range monthOffsetKey {
        if i == 0 {
            monthOffset[i] = []int8{4, 3}
        } else if i == 1 {
            monthOffset[i] = []int8{0, 6}
        } else {
            monthOffset[i] = []int8{value}
        }
    }
    if isLeapYear(year, calType) && len(monthOffset[month - 1]) > 1 {
        return monthOffset[month - 1][1]
    }
    return monthOffset[month - 1][0]
}

// Returns the day of the week or raises error if a date can't be calculated.
func Date(year int, month, date int8, calType string) (day string, err error) {
    calType = strings.ToUpper(calType)
    check := isRealDate(year, month, date, calType)
    if check {
        total := (((addYear(year, month, date, calType) +
            addMonth(year, month, calType) + date) % 7) + 7) % 7
        return days[total], nil
    }
    calType = strings.Title(strings.ToLower(calType))
    return "", errors.New("Cannot Calculate Date " + strconv.Itoa(int(year)) +
                          ", " + strconv.Itoa(int(month)) + ", " +
                          strconv.Itoa(int(date)) + ", " + calType)
}
