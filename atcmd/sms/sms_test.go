package sms_test

import (
	"T700test/modop"
	"T700test/usbcheck"
	"fmt"
	"testing"

	"github.com/tarm/serial"
)

var sms_com_port string = usbcheck.Getdrives()
var s *serial.Port = modop.Open_serial_port(sms_com_port)

func Test_CSMS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSMS", *s))
}

func Test_CPMS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CPMS", *s))
}

func Test_CMGF(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGF", *s))
}

func Test_CSCA(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSCA", *s))
}

func Test_CSMP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSMP", *s))
}

func Test_CSDH(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSDH", *s))
}

func Test_CSCB(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSCB", *s))
}

func Test_CSAS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSAS", *s))
}

func Test_CRES(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CRES", *s))
}

func Test_CNMI(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CNMI", *s))
}

func Test_CMGL_TEXT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGL_TEXT", *s))
}

func Test_CMGL_PDU(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGL_PDU", *s))
}

func Test_CMGR_TEXT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGR_TEXT", *s))
}

func Test_CMGR_PDU(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGR_PDU", *s))
}

func Test_CNMA_TEXT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CNMA_TEXT", *s))
}

func Test_CNMA_PDU(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CNMA_PDU", *s))
}

func Test_CMGS_TEXT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGS_TEXT", *s))
}

func Test_CMGS_PDU(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGS_PDU", *s))
}

func Test_CMSS_TEXT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMSS_TEXT", *s))
}

func Test_CMSS_PDU(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMSS_PDU", *s))
}

func Test_CMGW_TEXT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGW_TEXT", *s))
}

func Test_CMGW_PDU(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGW_PDU", *s))
}

func Test_CMGD(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGD", *s))
}

func Test_CMGC_TEXT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGC_TEXT", *s))
}

func Test_CMGC_PDU(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMGC_PDU", *s))
}

func Test_CMMS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CMMS", *s))
}

func Test_EQSI(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EQSI", *s))
}

func Test_ESMSS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESMSS", *s))
}

func Test_EMEMS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EMEMS", *s))
}

func Test_ETWS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ETWS", *s))
}

func Test_ECBMR(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECBMR", *s))
}

func Test_EMGS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EMGS", *s))
}

func Test_EMGW(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EMGW", *s))
}

func Test_EMGC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EMGC", *s))
}

func Test_ENMA(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ENMA", *s))
}

func Test_EMGL(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EMGL", *s))
}

func Test_ECSCBSW(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECSCBSW", *s))
}

func Test_ECSCBCFG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECSCBCFG", *s))
}

func Test_CPNER(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CPNER", *s))
}

func Test_EWAC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EWAC", *s))
}
