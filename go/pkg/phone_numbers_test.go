package pkg

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"testing"
)

func TestPhoneNumbers(t *testing.T) {
	phoneNumberStr := "+1671-6882898"
	phoneNumber, err := phonenumbers.Parse(phoneNumberStr, "")
	if err != nil {
		fmt.Println("无法解析电话号码:", err)
		return
	}

	isValid := phonenumbers.IsValidNumber(phoneNumber)
	fmt.Println(isValid)
}
