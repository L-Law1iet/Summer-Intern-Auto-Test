package ims_test

import (
	"T700test/modop"
	"T700test/usbcheck"
	"fmt"
	"testing"

	"github.com/tarm/serial"
)

var ims_com_port string = usbcheck.Getdrives()
var s *serial.Port = modop.Open_serial_port(ims_com_port)

func Test_CIREG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CIREG", *s))
}

func Test_CSSAC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSSAC", *s))
}

func Test_ECFGGET(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECFGGET", *s))
}

func Test_ECFGSET(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECFGSET", *s))
}

func Test_EPVSGET(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EPVSGET", *s))
}

func Test_EPVSSET(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EPVSSET", *s))
}

func Test_EIREG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIREG", *s))
}

func Test_EIMSDEREG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSDEREG", *s))
}

func Test_ENWSEL(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ENWSEL", *s))
}

func Test_EIRS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIRS", *s))
}

func Test_EIDRS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIDRS", *s))
}

func Test_EVVS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EVVS", *s))
}

func Test_EIMSCI(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSCI", *s))
}

func Test_EMOBD(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EMOBD", *s))
}

func Test_ESSAC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESSAC", *s))
}

func Test_EEMCINFO(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EEMCINFO", *s))
}

func Test_ESRVCCTFR(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESRVCCTFR", *s))
}

func Test_ESRVSTATE(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESRVSTATE", *s))
}

func Test_EAPMODE(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EAPMODE", *s))
}

func Test_EAPPROVE(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EAPPROVE", *s))
}

func Test_ESMSMOIP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESMSMOIP", *s))
}

func Test_ESMSMTIP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESMSMTIP", *s))
}

func Test_ESMMA(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESMMA", *s))
}

func Test_ESMSMAIP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESMSMAIP", *s))
}

func Test_CDU(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CDU", *s))
}

func Test_EVADSMOD(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EVADSMOD", *s))
}

func Test_EVADSREP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EVADSREP", *s))
}

func Test_ESIPHEADER(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESIPHEADER", *s))
}

func Test_EIMSRCS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSRCS", *s))
}

func Test_EIMSCFG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSCFG", *s))
}

func Test_EIMSPDIS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSPDIS", *s))
}

func Test_EIMSPCSCF(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSPCSCF", *s))
}

func Test_EIMSPDN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSPDN", *s))
}

func Test_EIMSCP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSCP", *s))
}

func Test_EIMSESS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSESS", *s))
}

func Test_EIWLPLEN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIWLPLEN", *s))
}

func Test_EIMSGEO(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSGEO", *s))
}

func Test_EIMSRTT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSRTT", *s))
}

func Test_PRTTCALL(t *testing.T) {
	fmt.Println(modop.TestAT("AT+PRTTCALL", *s))
}
