package main

import (
	"T700test/graphqlapi"
	"T700test/modop"
	"T700test/usbcheck"
	"database/sql"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	b64 "encoding/base64"

	"fmt"
	"regexp"
	"strings"

	"github.com/briandowns/spinner"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tarm/serial"
)

var (
	s                *serial.Port
	testexecid       string // test execution id e.g: 13495
	testcase_issueid string
	com_port         string
	DUTnumber        string
	date             string
	exec_count       int
	atcmd_test_id    int
	atcmd_name       string
	atcmd_data       string
	atcmd_category   string
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkDUTnumber() string {
	defer s.Close()
	str := modop.ExecAT("AT+EGMR", "=0,5", *s)
	re := regexp.MustCompile(`.*EGMR:.*`)
	matches := re.FindStringSubmatch(str)
	words := strings.Split(matches[0], " ")
	serialnumber := strings.Replace(words[1], `"`, "", -1)
	serialnumber = serialnumber[0 : len(serialnumber)-1]
	return serialnumber
}

func read_db_id(instr string) {
	MutiCols := 0
	db, err := sql.Open("sqlite3", "atcmd\\atcmd.db")
	check(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM atcmd WHERE atcmd_name = '" + instr + "'")
	check(err)
	for rows.Next() {
		MutiCols++
		err = rows.Scan(&atcmd_test_id, &atcmd_name, &atcmd_data, &atcmd_category)
		check(err)
		fmt.Printf("Test case id : %d , Category : %s\n", atcmd_test_id, atcmd_category)
	}
	if MutiCols > 1 {
		fmt.Println("This AT command have mutiple Test Cases, please choose one and enter the Test case id :")
		fmt.Scanln(&testcase_issueid)
	} else {
		tcid := strconv.Itoa(atcmd_test_id)
		testcase_issueid = tcid
	}
}

func init() {
	// progress indicator
	sp := spinner.New(spinner.CharSets[9], 50*time.Millisecond)
	sp.Start()
	exec_count = 1
	// 1. Read Serial Number from device
	com_port = usbcheck.Getdrives()
	s = modop.Open_serial_port(com_port)
	DUTnumber = checkDUTnumber()
	timeLocal := time.Now().Local()
	date = timeLocal.Format("0102")
	sp.Stop()
	fmt.Println("S/N : " + DUTnumber)
}

func main() {
	var instr string
	var cmd_res string
	bar := spinner.New(spinner.CharSets[9], 50*time.Millisecond)
	// 2. Enter test execution id
ENTER_EXEC_ID:
	fmt.Println("Please enter the Test Execution id e.g. 13481:")
	fmt.Scanln(&testexecid)
	for {
		var mutation_res string
		fmt.Printf("=============Start No.%d Test==============\n", exec_count)
	ENTER_ITEM:
		// 3. Enter test item (AT command)
		fmt.Println("Please enter the Test item e.g. AT+CFUN:")
		fmt.Scanln(&instr)
		instr = strings.ToUpper(instr)
		read_db_id(instr)
		if testcase_issueid == "0" {
			fmt.Println("Couldn't find test case, please enter a correct AT command.")
			goto ENTER_ITEM
		}
		// 4. Get testrun id from Jira
		bar.Start()
		token := graphqlapi.Getauthorize_api()
		data := make(map[string]string)
		data["testexecid"] = testexecid
		testjson := graphqlapi.Query(testcase_issueid, data, "getTestRun", token)
		testrunid := graphqlapi.Load_json(testjson, "testrunid")
		bar.Stop()
		if testrunid == "" {
			fmt.Println("Test case id doesn't match.\nPlease enter a correct Test Execution id or Test case id again.")
			goto ENTER_EXEC_ID
		}
		fmt.Println("testrunid:", testrunid)
		// 5. Run unit test for AT command
		os.Setenv("testcase_id", testcase_issueid)
		bar.Start()
		if runtime.GOOS == "windows" {
			cmd := exec.Command("cmd", "/c", "go", "test", "atcmd/atcmd_test.go", "-v")
			out, _ := cmd.CombinedOutput()
			cmd_res = string(out)
		} else { // Linux
			cmd := exec.Command("bash", "-c", "go", "test", "atcmd/atcmd_test.go", "-v")
			out, _ := cmd.CombinedOutput()
			cmd_res = string(out)
		}
		bar.Stop()
		fmt.Println(cmd_res)
		// 6. Check unit test status
		isPass := strings.Contains(cmd_res, "PASS")
		// 7. Update testrun status
		if isPass {
			data["status"] = "PASSED"
		} else {
			data["status"] = "FAILED"
		}
		bar.Start()
		mutation_res += graphqlapi.Mutation(testrunid, data, "updateTestRunStatus", token)
		mutation_res += "\n"
		// 8. Update testrun evidence
		sEnc := b64.StdEncoding.EncodeToString([]byte(cmd_res))
		data["file"] = instr + "_" + date + "_" + DUTnumber + ".log"
		data["textdata"] = sEnc
		mutation_res += graphqlapi.Mutation(testrunid, data, "addEvidenceToTestRun", token)
		bar.Stop()
		fmt.Println(mutation_res)
		exec_count++
		fmt.Println("=============End of test=============")
	}
}
