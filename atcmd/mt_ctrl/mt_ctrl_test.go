package mt_ctrl_test

import (
	"T700test/modop"
	"T700test/usbcheck"
	"fmt"
	"testing"

	"github.com/tarm/serial"
)

var mt_ctrl_com_port string = usbcheck.Getdrives()
var s *serial.Port = modop.Open_serial_port(mt_ctrl_com_port)

func Test_CPAS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CPAS", *s))
}

func Test_CFUN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CFUN", *s))
}

func Test_EFUN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EFUN", *s))
}

func Test_CPIN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CPIN", *s))
}

func Test_EPIN1(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EPIN1", *s))
}

func Test_EPIN2(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EPIN2", *s))
}

func Test_EPINC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EPINC", *s))
}

func Test_ESIMS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESIMS", *s))
}

func Test_CESQ(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CESQ", *s))
}

func Test_CMEC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMEC", *s))
}

func Test_CIND(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CIND", *s))
}

func Test_CMER(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMER", *s))
}

func Test_CPBS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CPBS", *s))
}

func Test_EPBUM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EPBUM", *s))
}

func Test_CSIM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSIM", *s))
}

func Test_CRSM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CRSM", *s))
}

func Test_CACM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CACM", *s))
}

func Test_CAMM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CAMM", *s))
}

func Test_CPUC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CPUC", *s))
}

func Test_CTZR(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CTZR", *s))
}

func Test_CGLA(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CGLA", *s))
}

func Test_CRLA(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CRLA", *s))
}

func Test_CCHO(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CCHO", *s))
}

func Test_CCHC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CCHC", *s))
}

func Test_CSUS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSUS", *s))
}

func Test_ECRSM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECRSM", *s))
}

func Test_ECPIN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECPIN", *s))
}

func Test_ICCID(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ICCID", *s))
}

func Test_CNUM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CNUM", *s))
}

func Test_ESIMAPP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESIMAPP", *s))
}

func Test_ECIMI(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECIMI", *s))
}

func Test_ESIMTC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESIMTC", *s))
}

func Test_ESIMAUTH(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESIMAUTH", *s))
}

func Test_USIMSMT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+USIMSMT", *s))
}

func Test_CAVIMS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CAVIMS", *s))
}

func Test_CASIMS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CASIMS", *s))
}

func Test_CEN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CEN", *s))
}

func Test_CISRVCC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CISRVCC", *s))
}

func Test_EGDM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EGDM", *s))
}

func Test_ESLOTSINFO(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESLOTSINFO", *s))
}

func Test_CUAD(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CUAD", *s))
}
