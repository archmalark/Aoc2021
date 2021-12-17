package main

import (
	"bufio"
	"fmt"
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

func ParsePacket(binaryString string) (int64, int64) {
	version, packetType := GetPacketInfo(binaryString)

	if packetType == 4 {
		_, size := ParsePacketType4(binaryString, 0)

		return size, version
	} else {
		lengthType := GetPacketLengthType(binaryString)
		if lengthType == 0 {
			subPacketLength, _ := strconv.ParseInt(binaryString[7:22], 2, 64)

			subPacketStart := 22
			subBitsParsed := 0
			versionSum := 0
			for true {

				subPacketSize, subVersion := ParsePacket(binaryString[subPacketStart:])
				subBitsParsed += int(subPacketSize)

				subPacketStart += int(subPacketSize)
				versionSum += int(subVersion)
				if subBitsParsed == int(subPacketLength) {
					break
				}
			}

			return int64(subBitsParsed) + 22, version + int64(versionSum)
		} else {
			numberOfSubPackets, _ := strconv.ParseInt(binaryString[7:18], 2, 64)
			subBitsParsed := 0
			subPacketStart := 18
			versionSum := 0

			for i := int64(0); i < numberOfSubPackets; i++ {
				subPacketSize, subVersion := ParsePacket(binaryString[subPacketStart:])
				subBitsParsed += int(subPacketSize)
				subPacketStart += int(subPacketSize)
				versionSum += int(subVersion)

			}
			return int64(subBitsParsed) + 18, version + int64(versionSum)
		}

	}

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

	_, versionSum := ParsePacket(binaryString)

	fmt.Println("VersionSum: ", versionSum)

}
