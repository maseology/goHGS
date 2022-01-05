package readers

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
)

func ReadHGSstructure(fp string) (coords [][3]float64, nps, eps, nsl int) {
	if !fileExists(fp) {
		log.Fatalf("ReadHGSstructure error: file %s does not exist\n", fp)
	}

	b, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalf("ReadHGSstructure readHGSscalar failed: %v\n", err)
	}

	buf := bytes.NewReader(b)
	readBuf := func(v interface{}) {
		err := binary.Read(buf, binary.LittleEndian, v)
		if err != nil {
			log.Fatalf("ReadHGSstructure ReadBinary failed: %v\n", err)
		}
	}

	var recb, rece, itmp int32
	testRec := func(recb int32) {
		readBuf(&rece)
		if recb != rece {
			log.Fatalf("ReadHGSstructure file read error recb != rece\n")
		}
	}
	chkRec := func(against int) {
		if recb != int32(against) {
			log.Fatalf("ReadHGSstructure file read error recb not as expected\n")
		}
	}

	var nn int32 // total number of nodes
	readBuf(&recb)
	readBuf(&nn)
	testRec(recb)

	readBuf(&recb)
	chkRec(8 * 3 * int(nn))
	coords = make([][3]float64, nn)
	if err := binary.Read(buf, binary.LittleEndian, &coords); err != nil {
		log.Fatalln("ReadHGSstructure read coordinates failed: ", err)
	}
	// var c [3]float64
	// for i := int32(0); i < nn; i++ {
	// 	readBuf(&c)
	// 	coords[i] = c
	// }
	testRec(recb)

	readBuf(&recb)
	var nx, ny, nz, nsptot int32
	readBuf(&nx)     // number of grid lines in x coordinates (for unstructured grid nx is neglected)
	readBuf(&ny)     // number of grid lines in y coordinates (for unstructured grid ny is neglected)
	readBuf(&nz)     // number of grid lines in z coordinates
	readBuf(&nsptot) // number of species for transport
	testRec(recb)

	readBuf(&recb)
	var epsl int32
	readBuf(&epsl) // number of elements in 2D node sheet
	testRec(recb)

	readBuf(&recb)
	readBuf(&itmp) // logical switch (.true. if it is a tetrahedral mesh)
	testRec(recb)

	readBuf(&recb)
	readBuf(&itmp) // logical switch (.true. if the Galerkin method is used for tetramesh)
	testRec(recb)

	readBuf(&recb)
	readBuf(&itmp) // logical switch (should always be .false.)
	testRec(recb)

	readBuf(&recb)
	readBuf(&itmp) // maximum number of nodes connected to a node for the 2D triangular mesh
	testRec(recb)

	if !reachedEOF(buf) {
		log.Fatalf("ReadHGSstructure more file to be read")
	}
	if nn%nz != 0 {
		log.Fatalf("ReadHGSstructure file error n nodes (%d) does not divide into %d slices\n", nn, nz)
	}

	nps = int(nn / nz)
	eps = int(epsl)
	nsl = int(nz)

	return
}
