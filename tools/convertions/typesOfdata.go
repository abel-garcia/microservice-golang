package convertions

import "strconv"

import "log"

func StringToInt64(date string) int64 {
	if number, err := strconv.Atoi(date); err == nil {
		return int64(number)
	} else if number > 0 {
		log.Panic("Error conver to number 64 :", err)
	}

	return 0
}
