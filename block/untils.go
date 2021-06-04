package block

import (
"bytes"
"encoding/binary"
"log"
)

//实现Int64转byte[]
func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Panicf("int trans to []byte failed! %v\n", err)
	}
	//	fmt.Print(utf8.DecodeRuneInString(buffer.String()))
	return buffer.Bytes()
}
