package modop

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tarm/serial"
)

var atcmds = make(map[string]string)

const (
	baud = 9600
	//interval of writing each AT command
	write_interval = 1
)

//Execute an AT command
func ExecAT(constcmd string, input string, s serial.Port) string {
	var response string
	//read regexp of AT command and match
	regexps := read_regexps(constcmd)
	isMatch := match(input, regexps)
	if !isMatch {
		log.Fatal("The AT command is not regular.")
	}
	atcmds[constcmd] = constcmd + input + "\n"
	response = write_and_check(s, write_interval, atcmds[constcmd])
	return response
}

//Execute tests of an AT command
func TestAT(value interface{}, s serial.Port) string {
	defer s.Close()
	v := value
	var response string
	cmds, args := read_tests(v)
	for i, atcmd := range cmds {
		response += ExecAT(atcmd, args[i], s)
	}
	return response
}

func Open_serial_port(serial_port string) *serial.Port {
	c := &serial.Config{Name: serial_port, Baud: baud}
	s, err := serial.OpenPort(c)
	check(err)
	return s
}

//Pass an AT command to device and receive response
func write_and_check(s serial.Port, write_interval int, cmd string) string {
	s.Write([]byte(cmd))
	time.Sleep(time.Duration(write_interval) * time.Second)
	buf := make([]byte, 1024)
	n, err := s.Read(buf)
	check(err)
	response := string(buf[:n])
	response += "\n"
	return response
}

//Decide the AT command is regular or not
func match(input string, reg []string) bool {
	var match []bool
	for _, i := range reg {
		imatch, _ := regexp.MatchString(i, input)
		match = append(match, imatch)
	}
	ismatch := false
	for _, matches := range match {
		ismatch = ismatch || matches
	}
	return ismatch
}

//Read data from database
func read_db(value interface{}) (int, string, string, string) {
	var db_test_id int
	var db_name, db_data, db_category string
	dirname, err := os.Getwd()
	check(err)
	var db_file sql.DB
	if strings.Contains(dirname, "xs_golang\\atcmd\\") {
		length := strings.LastIndex(dirname, "\\")
		runes := []rune(dirname)
		if length > len(runes) {
			length = len(runes)
		}
		dirname = string(runes[0:length])
		db, err := sql.Open("sqlite3", dirname+"\\atcmd.db")
		check(err)
		db_file = *db
	} else if strings.Contains(dirname, "xs_golang\\atcmd") {
		db, err := sql.Open("sqlite3", "atcmd.db")
		check(err)
		db_file = *db
	} else {
		db, err := sql.Open("sqlite3", dirname+"\\atcmd\\atcmd.db")
		check(err)
		db_file = *db
	}
	defer db_file.Close()
	switch v := value.(type) {
	case string:
		rows, err := db_file.Query("SELECT * FROM atcmd WHERE atcmd_name = '" + v + "'")
		check(err)
		for rows.Next() {
			err = rows.Scan(&db_test_id, &db_name, &db_data, &db_category)
			check(err)
		}
	case int:
		tcid := strconv.Itoa(v)
		rows, err := db_file.Query("SELECT * FROM atcmd WHERE atcmd_test_id =" + tcid)
		check(err)
		for rows.Next() {
			err = rows.Scan(&db_test_id, &db_name, &db_data, &db_category)
			check(err)
		}
	}
	return db_test_id, db_name, db_data, db_category
}

//Parse json and read regexp
func read_regexps(atcmd string) []string {
	_, _, db_data, _ := read_db(atcmd)
	byte_data := []byte(db_data)
	var data Data
	var regexps []string
	json.Unmarshal(byte_data, &data)
	for _, val := range data.List1 {
		regexps = append(regexps, val.Regexp)
	}
	return regexps
}

//Parse json and read tests of an AT command
func read_tests(value interface{}) ([]string, []string) {
	v := value
	_, _, db_data, _ := read_db(v)
	byte_data := []byte(db_data)
	var data Data
	var cmds, args []string
	json.Unmarshal(byte_data, &data)
	for _, val := range data.List2 {
		cmds = append(cmds, val.Atcmd)
		args = append(args, val.Arg)
	}
	return cmds, args
}

/*json structure
e.g AT+CPAS
{
  "regexps": [
    {
      "regexp": "^=[?]$"
    },
    {
      "regexp": ""
    }
  ],
  "tests": [
    {
      "atcmd": "AT+CPAS",
      "arg": "=?"
    },
    {
      "atcmd": "AT+CPAS",
      "arg": ""
    }
  ]
}
*/
type Regexps struct {
	Regexp string `json:"regexp"`
}

type Tests struct {
	Atcmd string `json:"atcmd"`
	Arg   string `json:"arg"`
}

type Data struct {
	List1 []Regexps `json:"regexps"`
	List2 []Tests   `json:"tests"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
