
package main

//import all needed functionality
import (
	"fmt"
	"regexp"
	"io/ioutil"
	"strings"
	"os"
	"strconv"
	"os/exec"
	"time"
)

//execute all method calls using multiple go routines
func main(){
	go extract_useful()
	time.Sleep(time.Millisecond * 1000)
	ip_pairs := organize_ipd()
	go extract_ipd(ip_pairs)
	time.Sleep(time.Millisecond * 1000)
}

//extract all the timestamps for each IP pair and write them to the corrispondig file
func organize_ipd() []string{
	inputfile := "input.useful"
	buffer, _ := ioutil.ReadFile(inputfile)

	lines := strings.Split(string(buffer), "\n")
	ip_formula := `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`
	ip_regular_exp := regexp.MustCompile(ip_formula)

	timestamp_pattern := `\s\d+\.\d+\s`
	timestamp_regularExpression := regexp.MustCompile(timestamp_pattern)

	ip_pairs := make([]string, 0)
	isFound := false

	for _, line := range lines{
		if line == ""{continue}

		ips := ip_regular_exp.FindAllString(line, -1)
		source_ip := ips[0]
		destination_ip := ips[1]
		timestamp := strings.TrimSpace(timestamp_regularExpression.FindString(line))

		//replace all "." with "_"
		pair := strings.Replace(source_ip + "_" + destination_ip, ".", "_", -1)

		//write to a file
		output_filename := pair + "_ipd_input.txt"
		f, err := os.OpenFile(output_filename, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil{
			fmt.Println("Error, CANT OPEN FILE")
		}

		f.WriteString(timestamp + "\n")

		f.Close()

		search:
		for _, p := range ip_pairs{
			if p == pair{
			isFound = true
			break search
			}
		}
		if !isFound{
			ip_pairs = append(ip_pairs, pair)
		}
	}
	return ip_pairs
}

//runs all needed terminal commands
func extract_useful(){
	rm_textfile_cmd := "rm *.txt"
	_, err := exec.Command("sh", "-c", rm_textfile_cmd).Output()
	if err != nil{}

	dump_file := "input.dump"

	tshark_command := "/usr/bin/tshark -Y 'ip' -r " + dump_file + " > input.useful"

	_, err2 := exec.Command("sh","-c",tshark_command).Output()
	if err2 != nil{}
}

//calculate the ipd and write them to a file
func extract_ipd(ip_pairs []string){
	for _, pair := range ip_pairs{
		filename := pair + "_ipd_input.txt"

		//create a slice of file contents
		timestamps := make([]string, 0)

		//read in file
		file_text, err := ioutil.ReadFile(filename)
		if err != nil {fmt.Println("ERROR, COULD NOT FIND FILE")}

		//split text into slices
		timestamps = strings.Split(string(file_text), "\n")

		//create and open a file for output
		output_filename := pair + "_ipd_output.txt"
		o, err := os.OpenFile(output_filename, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil{fmt.Println("Error, cant open file")}

		//calculate the packet delay and write that delay to the outfile
		for i:=1; i < len(timestamps)-1; i++{
			f1, err1 := strconv.ParseFloat(timestamps[i], 64)
			f2, err2 := strconv.ParseFloat(timestamps[i-1], 64)
			if err1 != nil || err2 != nil{fmt.Println("Error, could not convert float to string")}
			delay := f1 - f2
			o.WriteString(fmt.Sprintf("%.6f" , delay) + "\n")
		}

		//close the outfile
		o.Close()
	}
}
