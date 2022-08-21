package atcmd

import (
	"T700test/modop"
	"fmt"

	"github.com/tarm/serial"
)

func registIP(s serial.Port) {
	res += modop.ExecAT("AT", "", s)
	res += modop.ExecAT("AT+EGMR", "=0,5", s)
	res += modop.ExecAT("AT+EGMR", "=0,7", s)
	res += modop.ExecAT("AT+EIAAPN", "=\"internet\"", s)
	res += modop.ExecAT("AT+CGDCONT", "=1,\"ip\",\"internet\"", s)
	res += modop.ExecAT("AT+CFUN", "=1", s)
	res += modop.ExecAT("AT+CREG", "?", s)
	res += modop.ExecAT("AT+CGACT", "=1,1", s)
	res += modop.ExecAT("AT+CGPADDR", "=1", s)
	res += modop.ExecAT("AT+EDMFAPP", "=6,4", s)
}

func getband(s serial.Port) {
	res += modop.ExecAT("AT+EDMFAPP", "=6,4", s)
	res += modop.ExecAT("AT+ECSQ", "", s)
	res += modop.ExecAT("AT+CGSN", "", s)
}

func shutdown(s serial.Port) {
	res += modop.ExecAT("AT+CFUN", "=0", s)
}

func turnOnOff(s serial.Port) {
	res += modop.ExecAT("AT+CFUN", "=0", s)
	res += modop.ExecAT("AT+CFUN", "=1", s)
	res += modop.ExecAT("AT+CREG", "?", s)
}

var res string

func Send_at_cmd(mode string, serial_port string) string {
	s := modop.Open_serial_port(serial_port)
	defer s.Close()
	switch {
	case mode == "registIP":
		registIP(*s)
	case mode == "getband":
		getband(*s)
	case mode == "shutdown":
		shutdown(*s)
	case mode == "turnOnOff":
		turnOnOff(*s)
	default:
		fmt.Println("no match mode")
	}
	return res
}
