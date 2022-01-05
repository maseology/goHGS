package readers

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func readInt32(b *bytes.Reader) (i int32) {
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		log.Fatalf("(readInt32) %v", err)
	}
	return
}

func ReadHGSvector(fp string, nps, epl, nly int) (float64, [][3]float32) {
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	flen := len(b)
	buf := bytes.NewReader(b)

	// get timestep
	recb := readInt32(buf)
	tstxt := make([]byte, recb)
	if err := binary.Read(buf, binary.LittleEndian, &tstxt); err != nil {
		log.Fatalln("readHGS read timestep failed: ", err)
	}
	if readInt32(buf) != recb {
		log.Fatal("timestep read error")
	}
	t, err := strconv.ParseFloat(strings.TrimSpace(string(tstxt)), 64)
	if err != nil {
		log.Fatalln("readHGS read timestep failed: ", err)
	}

	if (flen-88)%(12+8) != 0 { ////////// HARD-CODED "12" //////////
		log.Fatalln("readHGS file size error: ", err)
	}
	nnds := (flen - 88) / (12 + 8) ////////// HARD-CODED "12" //////////
	vlist := make([]float32, nnds)
	if err := binary.Read(buf, binary.LittleEndian, &vlist); err != nil {
		log.Fatalln("readHGS read vector failed: ", err)
	}

	vs := make([][3]float32, nnds)
	nv := (12 + 8) / 4 ////////// HARD-CODED "12" //////////
	bn := math.Float32frombits(12)
	for i := 0; i < int(nnds)/nv; i++ {
		if vlist[i*nv] != bn && vlist[i*nv+nv-1] != bn {
			log.Fatal("vector read error")
		}
		vs[i][0] = vlist[i*nv+1] ////////// HARD-CODED "12" //////////
		vs[i][1] = vlist[i*nv+2] ////////// HARD-CODED "12" //////////
		vs[i][2] = vlist[i*nv+3] ////////// HARD-CODED "12" //////////
	}

	// fmt.Printf("     %f %d %d  %s\n", t, 12, (flen-88)/(12+8), fp)

	return t, vs
}
