package datecalc

var days []string = []string{
    "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
}
var calTypes []string = []string{
    "GREGORIAN", "CE", "JULIAN", "ENGLISH", "ROMAN",
}
var monthNames []string = []string{
    "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
}

func inRange(x, min, max int8) bool {
    for i := min; i < max; i ++ {
        if i == x {
            return true
        }
    }
    return false
}

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

func julLeapYear(y int) bool {
    if y < 0 {
        y = (((y + 1) % 700) + 700) % 700
    }
    if y % 4 == 0 {
        return true
    }
    return false
}

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

func isRealDate(year int, month, date int8, calType string) int {
    month30 := []int8{4, 6, 9, 11}
    if inRange(month, 1, 13) == false {
        return 0
    }
    if date < 1 || date > 31 {
        return 0
    }
    inTypes := false
    for _, x := range calTypes {
        if calType == x {
            inTypes = true
            break
        }
    }
    if !inTypes {
        return 0
    }
    if year == 0 && calType != "CE" {
        return 0
    } else {
        inMonths := false
        for _, x := range month30 {
            if month == x {
                inMonths = true
                break
            }
        }
        if inMonths && date > 30 {
            return 0
        } else if month == 1 && !(date < 29 || (isLeapYear(year, calType) && date == 29)) {
            return 0
        }
    }
    if calType == "ENGLISH" {
        if month == 9 && date > 2 && date < 15 && year == 1752 {
            return 0
        }
    } else if calType == "ROMAN" {
        if month == 10 && date > 4 && date < 16 && year == 1582 {
            return 0
        }
    }
    return 1
}

func addxxYY(year int, calType string) int8 {
    newYear := ((year % 100) + 100) % 100
    if calType != "CE" && year < 0 {
        newYear = (((year + 1) % 100) + 100) % 100
    }
    return  int8((((newYear / 12 + (((newYear % 12) + 12) % 12) + ((((newYear % 12) / 12) % 12) / 4)) % 7) + 7) % 7)
}

func ceAddYYxx(y int) int8 {
    YYxx := []int8{2, 0, 5, 3}
    return YYxx[(((y / 100) % 4) + 4) % 4]
}

func julAddYYxx(y int) int8 {
    return int8((((7 - y / 100) % 7) + 7) % 7)
}

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

func addYear(year int, month, date int8, calType string) int8{
    return addYYxx(year, month, date, calType) + addxxYY(year, calType)
}

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

func Date(year int, month, date int8, calType string) string {
    check := isRealDate(year, month, date, calType)
    if check == 1 {
        total := (((addYear(year, month, date, calType) + addMonth(year, month, calType) + date) % 7) + 7) % 7
        return days[total]
    }
    return ""
}