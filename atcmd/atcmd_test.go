package atcmd

import (
	"T700test/modop"
	"T700test/usbcheck"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/tarm/serial"
)

var atcmd_com_port string = usbcheck.Getdrives()
var s *serial.Port = modop.Open_serial_port(atcmd_com_port)
var id_env string
var testcase_id int

func init() {
	id_env = os.Getenv("testcase_id")
	testcase_id, _ = strconv.Atoi(id_env)
}
func Test_AT_Command(t *testing.T) {
	test_res := modop.TestAT(testcase_id, *s)
	fmt.Println(test_res)
}
