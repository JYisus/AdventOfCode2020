package day05

import (
	"errors"
)

func CheckRow(seatCode string) (int, error) {
	totalSeats := 128

	minSeat := 0
	maxSeat := 127
	for i := 0; i < len(seatCode)-3; i++ {
		switch string(seatCode[i]) {
		case "F":
			totalSeats /= 2
			maxSeat = maxSeat - totalSeats
		case "B":
			totalSeats /= 2
			minSeat = minSeat + totalSeats

		}
	}

	if maxSeat == minSeat {
		return maxSeat, nil
	}

	return 0, errors.New("invalid code")
}

func CheckColumn(seatCode string) (int, error) {
	aux := 8

	minSeat := 0
	maxSeat := 7
	for i := len(seatCode) - 3; i < len(seatCode); i++ {
		switch string(seatCode[i]) {
		case "L":
			aux = aux / 2
			maxSeat = maxSeat - aux
			// fmt.Println("kmax seat")

			// fmt.Println(maxSeat)

		case "R":
			aux = aux / 2
			minSeat = minSeat + aux
			// fmt.Println("min seat")
			// fmt.Println(minSeat)

		}
	}

	if maxSeat == minSeat {
		return maxSeat, nil
	}

	return 0, errors.New("invalid code")
}

func SeatID(seatCode string) (int, error) {
	row, _ := CheckRow(seatCode)
	column, _ := CheckColumn(seatCode)

	return row*8 + column, nil
}

func AoC05(codes []string) (int, error) {
	biggestVal := 0
	for _, code := range codes {
		id, _ := SeatID(code)
		if id > biggestVal {
			biggestVal = id
		}
	}
	return biggestVal, nil
}
