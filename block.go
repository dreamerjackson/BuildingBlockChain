package main

import (
	"bytes"
	"fmt"
	"encoding/binary"
	"log"
	"encoding/hex"
	"crypto/sha256"
)

// IntToHex converts an int32 to a byte array
//将int32转换为字节数组
func IntToHex(num int32) []byte {
	buff := new(bytes.Buffer)
	//还可以大小段模式
	err := binary.Write(buff, binary.LittleEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}


//将32位的int转换为16进制，16进制的每个数字转成4位
func change( num int32) [4]byte{
	var buffle [8]byte
	for i:=0;i<8;i++{
		buff := new(bytes.Buffer)
		binary.Write(buff, binary.BigEndian, num%16)
		//buffle = bytes.Join([][]byte{buffle,buff.Bytes()}, []byte{})
		result := (buff.Bytes())[3]
		buffle[7-i]=result
		num = num >> 4
	}
	fmt.Printf("j = %x\n",buffle)
	var buffle2 [4]byte
		var j uint  = 0;
	for ;j<8;j++{
		buffle2[j/2] = buffle2[j/2] | ( buffle[j]<<(4*((j+1)%2)) );
	}
	return buffle2
}

//大小端转换
func reverse2(data []byte) []byte{
	var s [][]byte
	for i:=len(data);i>0;i--{
		data1 :=data[i-1:i]
		s = append(s,data1)
	}
	sep := []byte("")
	result :=bytes.Join(s,sep)
	return result
}



func main(){
	////版本
	version :=IntToHex(2)

	////前一个区块链的hash
	PrevBlockHash,_:=hex.DecodeString("00000000000000000A2940884E0C3BC96510CAD11912A527E9D15DF42F0E1D67")
	PrevBlockHash  =  reverse2(PrevBlockHash)

	////默克尔根
	////timestamp := []byte(strconv.FormatInt(int64((time-25569)*86400), 16))
	////fmt.Printf("%x\n",[]byte((time-25569)*86400)
	// fmt.Printf("%x\n",PrevBlockHash)
	MerkleRoot,_:=hex.DecodeString("2E99F445C007A9158207CC30CEBAD2B3D26C45FDAB2EBDF50D261335FC00D92C")
	MerkleRoot  =  reverse2(MerkleRoot)

	////时间
	var time2 int32
	time2=1418753140
	var buffletime [4]byte
	buffletime = change(time2)
	reverstime := reverse2(buffletime[:])

	//难度
	var bits int32 = 404454260
	var buff [4]byte
	buff = change(bits)
	reversbuff := reverse2(buff[:])

	//随机Nonce
	var Nonce int32 = 123
	var buffNonce [4]byte
	buffNonce = change(Nonce)
	reversbuffNonce := reverse2(buffNonce[:])

	//字节
	result := bytes.Join([][]byte{version,PrevBlockHash,MerkleRoot,reverstime,reversbuff,reversbuffNonce}, []byte{})

	//双hash
	hash := sha256.Sum256(result)
	rev := sha256.Sum256(hash[:])
	fmt.Printf("%x",rev)
}
