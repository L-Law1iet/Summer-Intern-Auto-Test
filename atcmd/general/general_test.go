package general_test

import (
	"T700test/modop"
	"T700test/usbcheck"
	"fmt"
	"testing"

	"github.com/tarm/serial"
)

var general_com_port string = usbcheck.Getdrives()
var s *serial.Port = modop.Open_serial_port(general_com_port)

func Test_CGMI(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CGMI", *s))
}

func Test_CGMM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CGMM", *s))
}

func Test_CGMR(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CGMR", *s))
}

func Test_CGSN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CGSN", *s))
}

func Test_CSCS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSCS", *s))
}

func Test_CLAC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CLAC", *s))
}

func Test_CIMI(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CIMI", *s))
}
