package codec

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Decode /解码
func Decode(reader *bufio.Reader) (string, error) {
	lengthBytes, err := reader.Peek(4) //4字节的消息长度
	if err != nil {
		return "", err
	}
	lengthBuf := bytes.NewBuffer(lengthBytes)
	var length int32
	err = binary.Read(lengthBuf, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	//出现了拆包
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	msgByte := make([]byte, length+4)
	reader.Read(msgByte)
	return string(msgByte[4:]), nil
}

// Encode /编码
func Encode(message string) ([]byte, error) {
	length := int32(len(message))
	pkg := new(bytes.Buffer)
	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息实体
	err2 := binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err2
	}
	return pkg.Bytes(), nil
}
