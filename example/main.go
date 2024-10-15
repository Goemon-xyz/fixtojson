package main

import (
	"fmt"

	fixjson "github.com/goemon-xyz/fixtojson"
)

func main() {
	converter, err := fixjson.NewConverter("spec/FIX44-PT.xml")
	if err != nil {
		fmt.Printf("Error creating converter: %v\n", err)
		return
	}

	// Use the ASCII character 1 (SOH) as the delimiter
	soh := string([]byte{1})

	// Construct the message body
	messageBody := "35=D" + soh +
		"49=SENDER" + soh +
		"56=TARGET" + soh +
		"34=1" + soh +
		"52=20230101-12:00:00" + soh +
		"11=OrderID" + soh +
		"55=SYMBOL" + soh +
		"54=1" + soh +
		"38=100" + soh +
		"40=2" + soh

	// Calculate the correct message length
	bodyLength := len(messageBody)

	// Construct the full message with correct length
	rawFIXMessage := "8=FIX.4.4" + soh +
		fmt.Sprintf("9=%d", bodyLength) + soh +
		messageBody +
		"10=000" + soh

	jsonData, err := converter.FIXToJSON([]byte(rawFIXMessage))
	if err != nil {
		fmt.Printf("Error converting FIX to JSON: %v\n", err)
		return
	}

	fmt.Printf("FIX Message as JSON:\n%s\n", jsonData)
}
