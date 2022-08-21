package nw_serv_test

import (
	"T700test/modop"
	"T700test/usbcheck"
	"fmt"
	"testing"

	"github.com/tarm/serial"
)

var nw_serv_com_port string = usbcheck.Getdrives()
var s *serial.Port = modop.Open_serial_port(nw_serv_com_port)

func Test_CNUM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CNUM", *s))
}

func Test_CREG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CREG", *s))
}

func Test_COPS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+COPS", *s))
}

func Test_CLCK(t *testing.T) {
	fmt.Println(modop.TestAT("AT+COPS", *s))
}

func Test_ECLCK(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECLCK", *s))
}

func Test_CPWD(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CPWD", *s))
}

func Test_CLIP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CLIP", *s))
}

func Test_CLIR(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CLIR", *s))
}

func Test_COLP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+COLP", *s))
}

func Test_COLR(t *testing.T) {
	fmt.Println(modop.TestAT("AT+COLR", *s))
}

func Test_CNAP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CNAP", *s))
}

func Test_CCUG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CCUG", *s))
}

func Test_CCFC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CCFC", *s))
}

func Test_CCWA(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CCWA", *s))
}

func Test_CHLD(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CHLD", *s))
}

func Test_CUSD(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CUSD", *s))
}

func Test_ECUSD(t *testing.T) {
	fmt.Println("\nExecute SS/USSD operation (Proprietary command)\n# The MTK documents have this command, but XS5G03 at_cmd has no.")
	t.Fail()
}

func Test_CAOC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CAOC", *s))
}

func Test_CSSN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSSN", *s))
}

func Test_CLCC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CLCC", *s))
}

func Test_CLCCS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CLCCS", *s))
}

func Test_CPOL(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CPOL", *s))
}

func Test_EPOL(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EPOL", *s))
}

func Test_CPLS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CPLS", *s))
}

func Test_COPN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+COPN", *s))
}

func Test_CAEMLPP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CAEMLPP", *s))
}

func Test_EPBSE(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EPBSE", *s))

}

func Test_EPBSEH(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EPBSEH", *s))
}

func Test_EOPN(t *testing.T) {
	//FAIL
	fmt.Println(modop.TestAT("AT+EOPN", *s))
}

func Test_ECSQ(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECSQ", *s))
}

func Test_EFD(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EFD", *s))
}

func Test_ESCRI(t *testing.T) {
	fmt.Println("Empty content in the document")
	t.Fail()
}

func Test_ECSG(t *testing.T) {
	//FAIL
	fmt.Println(modop.TestAT("AT+ECSG", *s))
}

func Test_ERAT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ERAT", *s))
}

func Test_EPRATL(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EPRATL", *s))
}

func Test_EGTYPE(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EGTYPE", *s))
}

func Test_EHSM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EHSM", *s))
}

func Test_ECELL(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECELL", *s))
}

func Test_ESSP(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ESSP", *s))
}

func Test_CCBS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CCBS", *s))
}

func Test_EREGCHK(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EREGCHK", *s))
}

func Test_CSRA(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSRA", *s))
}

func Test_ECSRA(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECSRA", *s))
}

func Test_ERMS(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ERMS", *s))
}

func Test_EMODCFG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EMODCFG", *s))
}

func Test_EREGINFO(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EREGINFO", *s))
}

func Test_ERPRAT(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ERPRAT", *s))
}

func Test_CSCM(t *testing.T) {
	fmt.Println(modop.TestAT("AT+CSCM", *s))
}

func Test_ECREG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECREG", *s))
}

func Test_ECGREG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECGREG", *s))
}

func Test_ECEREG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECEREG", *s))
}

func Test_ECASW(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ECASW", *s))
}

func Test_EREG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EREG", *s))
}

func Test_EGREG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EGREG", *s))
}

func Test_ELCE(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ELCE", *s))
}

func Test_ETM9(t *testing.T) {
	fmt.Println(modop.TestAT("AT+ETM9", *s))
}

func Test_EWIFIRSSITHRCFG(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EWIFIRSSITHRCFG", *s))
}

func Test_EWIFIRSSIVER(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EWIFIRSSIVER", *s))
}

func Test_EWIFISIGLVL(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EWIFISIGLVL", *s))
}

func Test_EIMSRCSCONN(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EIMSRCSCONN", *s))
}

func Test_EWIFIASC(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EWIFIASC", *s))
}

func Test_EWIFIADDR(t *testing.T) {
	fmt.Println(modop.TestAT("AT+EWIFIADDR", *s))
}
