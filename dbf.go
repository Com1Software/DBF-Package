package dbf

import (
	//	"bufio"
	"fmt"
	"os"
	//	"strings"
	"io/ioutil"
	"strconv"
)

//------------------ (c) 1992-2022 Com1 Software Development
//
//Checks for full file name. Returns true if the file
//is present and false if it is not.
//
func CheckForFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	file.Close()
	return true
}

//
//Loads a dbf file and returns a byte array.
//
func LoadFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte("")
	}
	return content
}

//
//Returns the record count of the loaded dbf file.
//
func GetRecordCount(content []byte) int {
	rc := 0
	//------------------- Byte 5
	xrc := content[4:5]
	i := fmt.Sprintf("%d", xrc)
	ii := i[1 : len(i)-1]
	iii, _ := strconv.Atoi(ii)
	rc = rc + iii
	//------------------- Byte 6
	xrc = content[5:6]
	i = fmt.Sprintf("%d", xrc)
	ii = i[1 : len(i)-1]
	iii, _ = strconv.Atoi(ii)
	rc = rc + (iii * 256)
	//------------------- Byte 7
	xrc = content[6:7]
	i = fmt.Sprintf("%d", xrc)
	ii = i[1 : len(i)-1]
	iii, _ = strconv.Atoi(ii)
	rc = rc + (iii * 256)
	//------------------- Byte 8
	xrc = content[7:8]
	i = fmt.Sprintf("%d", xrc)
	ii = i[1 : len(i)-1]
	iii, _ = strconv.Atoi(ii)
	rc = rc + (iii * 256)
	return rc
}

//Returns the length of the records in the loaded dbf file.
func GetRecordLength(content []byte) int {
	rc := 0
	//------------------- Byte 11
	xrc := content[10:11]
	i := fmt.Sprintf("%d", xrc)
	ii := i[1 : len(i)-1]
	iii, _ := strconv.Atoi(ii)
	rc = rc + iii
	//------------------- Byte 12
	xrc = content[11:12]
	i = fmt.Sprintf("%d", xrc)
	ii = i[1 : len(i)-1]
	iii, _ = strconv.Atoi(ii)
	rc = rc + (iii * 256)
	return rc
}

//
//Returns the first records position in the loaded dbf file.
//
func GetFirstRecordPosition(content []byte) int {
	rc := 0
	//------------------- Byte 9
	xrc := content[8:9]
	i := fmt.Sprintf("%d", xrc)
	ii := i[1 : len(i)-1]
	iii, _ := strconv.Atoi(ii)
	rc = rc + iii
	//------------------- Byte 10
	xrc = content[9:10]
	i = fmt.Sprintf("%d", xrc)
	ii = i[1 : len(i)-1]
	iii, _ = strconv.Atoi(ii)
	rc = rc + (iii * 256)
	//------------------- Byte 7
	xrc = content[6:7]
	i = fmt.Sprintf("%d", xrc)
	ii = i[1 : len(i)-1]
	iii, _ = strconv.Atoi(ii)
	rc = rc + (iii * 256)
	//------------------- Byte 8
	xrc = content[7:8]
	i = fmt.Sprintf("%d", xrc)
	ii = i[1 : len(i)-1]
	iii, _ = strconv.Atoi(ii)
	rc = rc + (iii * 256)
	return rc
}

//
//Returns a specific record from the loaded dbf file.
//
func GetRecord(content []byte, ctl int) []byte {
	rec := []byte("")
	ii := GetFirstRecordPosition(content)
Record:
	for i := 0; i < GetRecordCount(content); i++ {
		x := GetRecordLength(content)
		s := ii
		e := s + x
		rec = content[s:e]
		ii = ii + x
		if i == ctl {
			break Record
		}
	}
	return rec
}

//
//Returns the field count of the loaded dbf file.
//
func GetFieldCount(content []byte) int {
	rc := GetFirstRecordPosition(content)
	rc = (rc / 32)
	return rc
}

//
//Returns a specified fields name.
//
func GetFieldName(content []byte, ctl int) []byte {
	fieldname := []byte("")
	rc := GetFirstRecordPosition(content)
	rc = (rc / 32)
	ii := 32
FieldName:
	for i := 0; i < rc; i++ {
		s := ii
		e := s + 32
		rec := content[s:e]
		fieldname = rec[0:10]
		ii = ii + 32
		if i == ctl {
			break FieldName
		}
	}
	return fieldname
}

//
//Returns a specied fields lenghth.
//
func GetFieldLength(content []byte, ctl int) int {
	fieldlen := 0
	rc := GetFirstRecordPosition(content)
	rc = (rc / 32)
	ii := 32
FieldLength:
	for i := 0; i < rc; i++ {
		s := ii
		e := s + 32
		rec := content[s:e]
		temp := fmt.Sprintf("%d", rec[16:17])
		tempa := temp[1 : len(temp)-1]
		fieldlen, _ = strconv.Atoi(string(tempa))
		ii = ii + 32
		if i == ctl {
			break FieldLength
		}
	}
	return fieldlen
}

//
//Returns the Data in the specified Field and Record.
//
func GetRecordField(content []byte, rec []byte, ctl int) []byte {
	// fmt.Println(string(rec))
	fld := []byte("")
	fieldlen := 0
	rflen := 1
	rc := GetFirstRecordPosition(content)
	rc = (rc / 32)
	ii := 32
RecordField:
	for i := 0; i < rc; i++ {
		s := ii
		e := s + 32
		trec := content[s:e]
		temp := fmt.Sprintf("%d", trec[16:17])
		tempa := temp[1 : len(temp)-1]
		fieldlen, _ = strconv.Atoi(string(tempa))
		ii = ii + 32
		if i == ctl {
			fld = rec[rflen : rflen+fieldlen]
			break RecordField
		}
		rflen = rflen + fieldlen
	}

	return fld
}
