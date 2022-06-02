package main

import (
	trippb "coolcar/proto/gen/go"
	"fmt"

)

func main() {
	trip := trippb.Trip{
		Start: "abc",
		End: "def",
		DurationSec: 3600,
		FeeCent: 10000,
	}
	
	fmt.Println(&trip)
}