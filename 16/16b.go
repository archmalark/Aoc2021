package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ParsePacketType4(binaryString string, position int64) (int64, int64) {

	packetBitSize := int64(6)
	var buffer uint64
	if len(binaryString) >= 64 {
		buffer, _ = strconv.ParseUint(binaryString[position:position+64], 2, 64)
	} else {
		buffer, _ = strconv.ParseUint(binaryString[position:position+int64(len(binaryString))], 2, 64)
		buffer = buffer << (64 - len(binaryString))
	}

	buffer = buffer << 6
	literalNumber := uint64(0)

	for true {
		literalNumber = literalNumber << 4

		group := buffer >> 59
		groupNumber := (group & 0b1111)

		literalNumber = literalNumber | groupNumber

		packetBitSize += 5
		if (buffer & 0x8000000000000000) == 0 {
			break
		}
		buffer = buffer << 5
	}

	return int64(literalNumber), packetBitSize

}

func GetPacketInfo(binaryString string) (int64, int64) {

	version, _ := strconv.ParseInt(binaryString[0:3], 2, 64)
	packetType, _ := strconv.ParseInt(binaryString[3:6], 2, 64)

	return version, packetType
}

func GetPacketLengthType(binaryString string) int64 {

	lengthType, _ := strconv.ParseInt(string(binaryString[6]), 2, 64)

	return lengthType
}

func ParsePacket(binaryString string, functionMap map[int64]func(int64, int64, bool) int64) (int64, int64) {
	_, packetType := GetPacketInfo(binaryString)

	if packetType == 4 {
		number, size := ParsePacketType4(binaryString, 0)

		return size, number
	} else {
		lengthType := GetPacketLengthType(binaryString)
		if lengthType == 0 {
			subPacketLength, _ := strconv.ParseInt(binaryString[7:22], 2, 64)

			subPacketStart := 22
			subBitsParsed := 0
			returnValue := int64(0)
			first := true
			for true {

				subPacketSize, number := ParsePacket(binaryString[subPacketStart:], functionMap)

				returnValue = functionMap[packetType](returnValue, number, first)
				first = false

				subBitsParsed += int(subPacketSize)

				subPacketStart += int(subPacketSize)
				if subBitsParsed == int(subPacketLength) {
					break
				}
			}

			return int64(subBitsParsed) + 22, returnValue
		} else {
			numberOfSubPackets, _ := strconv.ParseInt(binaryString[7:18], 2, 64)
			subBitsParsed := 0
			subPacketStart := 18
			returnValue := int64(0)
			first := true

			for i := int64(0); i < numberOfSubPackets; i++ {
				subPacketSize, number := ParsePacket(binaryString[subPacketStart:], functionMap)

				returnValue = functionMap[packetType](returnValue, number, first)
				first = false

				subBitsParsed += int(subPacketSize)
				subPacketStart += int(subPacketSize)

			}
			return int64(subBitsParsed) + 18, returnValue
		}

	}

}

func Sum(one int64, two int64, first bool) int64 {
	return one + two
}

func Product(one int64, two int64, first bool) int64 {
	if first {
		one = 1
	}
	return one * two
}

func Min(one int64, two int64, first bool) int64 {
	if first {
		return two
	}

	return int64(math.Min(float64(one), float64(two)))
}

func Max(one int64, two int64, first bool) int64 {
	if first {
		return two
	}

	return int64(math.Max(float64(one), float64(two)))
}

func Gt(one int64, two int64, first bool) int64 {
	if first {
		return two
	}
	if one > two {
		return 1
	}
	return 0
}

func Lt(one int64, two int64, first bool) int64 {
	if first {
		return two
	}
	if one < two {
		return 1
	}
	return 0
}

func Eq(one int64, two int64, first bool) int64 {
	if first {
		return two
	}
	if one == two {
		return 1
	}
	return 0
}

func main() {

	// Parse input

	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	binaryString := ""
	for scanner.Scan() {

		inputString := scanner.Text()

		for _, char := range inputString {
			nibbleInt, _ := strconv.ParseInt(string(char), 16, 64)
			nibbleString := strconv.FormatInt(nibbleInt, 2)
			binaryString += fmt.Sprintf("%04s", nibbleString)
		}

	}

	functionMap := map[int64]func(int64, int64, bool) int64{}
	functionMap[0] = Sum
	functionMap[1] = Product
	functionMap[2] = Min
	functionMap[3] = Max
	functionMap[5] = Gt
	functionMap[6] = Lt
	functionMap[7] = Eq
	_, packetValue := ParsePacket(binaryString, functionMap)

	fmt.Println("packetValue: ", packetValue)
	// fmt.Println(binaryString)

}
