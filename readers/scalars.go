package readers

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadHGSscalar(fp string) (float64, []float64) {
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalf("ReadHGSscalar failed: %v\n", err)
	}
	reclen := len(b) - 96
	if reclen%4 != 0 {
		log.Fatalln("ReadHGSscalar failed: invalid file size")
	}

	buf := bytes.NewReader(b)
	readBuf := func(v interface{}) {
		err := binary.Read(buf, binary.LittleEndian, v)
		if err != nil {
			log.Fatalf("ReadHGSscalar ReadBinary failed: %v\n", err)
		}
	}

	var recb, rece int32
	testRec := func() {
		readBuf(&rece)
		if recb != rece {
			log.Fatalf("ReadHGSscalar file read error recb != rece\n")
		}
	}

	readBuf(&recb)
	bts := make([]byte, recb)
	readBuf(&bts)
	ts, err := strconv.ParseFloat(strings.TrimSpace(string(bts)), 64)
	if err != nil {
		log.Fatalf("ReadHGSscalar parse float not happening: %v\n", err)
	}
	testRec()

	readBuf(&recb)
	if recb%8 != 0 {
		log.Fatalf("ReadHGSscalar file read error 1\n")
	}
	vals := make([]float64, recb/8)
	readBuf(&vals)
	testRec()

	if !reachedEOF(buf) {
		log.Fatalf("ReadHGSscalar more file to be read")
	}

	return ts, vals
}
