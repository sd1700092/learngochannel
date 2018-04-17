package main

import (
	"os"
	"encoding/binary"
	"fmt"
)

type BitmapInfoHeader struct{
	Size uint32
	Width int32
	Height int32
	Places uint16
	BitCount uint16
	Compression uint32
	SizeImage uint32
	XperlsPerMeter int32
	YperlsPerMeter int32
	ClsrUsed uint32
	ClrImportant uint32
}

func main() {
	file, _:=os.Open("image.bmp")
	defer file.Close()
	var headA, headB byte
	// windows/linux都是小端，而MacOS则是大端
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)
	fmt.Println("%c%c", headA, headB)

	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	var reservedA, reservedB uint16
	binary.Read(file, binary.LittleEndian, &reservedA)
	binary.Read(file, binary.LittleEndian, &reservedB)

	var offbits uint32
	binary.Read(file, binary.LittleEndian, &offbits)

	fmt.Println(headA, headB, size, reservedA, reservedB, offbits)


	infoHeader:=new(BitmapInfoHeader)
	binary.Read(file, binary.LittleEndian, infoHeader)
	fmt.Println(infoHeader)
}
