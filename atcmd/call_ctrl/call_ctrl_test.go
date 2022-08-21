package call_ctrl_test

import (
	"T700test/modop"
	"T700test/usbcheck"
	"fmt"
	"testing"

	"github.com/tarm/serial"
)

var call_ctrl_com_port string = usbcheck.Getdrives()
var s *serial.Port = modop.Open_serial_port(call_ctrl_com_port)

func Test_CSTA(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSTA", *s))
}

func Test_CHUP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CHUP", *s))
}

func Test_CR(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CR", *s))
}

func Test_CEER(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CEER", *s))
}

func Test_CRC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CRC", *s))
}

func Test_CVHU(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CVHU", *s))
}

func Test_EAIC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EAIC", *s))
}

func Test_ECPI(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECPI", *s))
}

func Test_EALS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EALS", *s))
}

func Test_ESVC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESVC", *s))
}

func Test_FCLASS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+FCLASS", *s))
}

func Test_VGR(t *testing.T) {
	fmt.Println(modop.TestAT("AT+VGR", *s))
}

func Test_VGT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+VGT", *s))
}

func Test_EVOCD(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EVOCD", *s))
}

func Test_EIMSHEADER(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSHEADER", *s))
}

func Test_EIMSCALLMODE(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSCALLMODE", *s))
}

func Test_EIMSCAI(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSCAI", *s))
}
