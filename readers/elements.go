package readers

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
)

func GetElementsPM(fp string, nly int) ([][]int, int, int) {
	bfile, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	b := bytes.NewReader(bfile)

	readInt32 := func(b *bytes.Reader) (i int32) {
		if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
			log.Fatalf("(readInt32) %v", err)
		}
		return
	}

	recb := readInt32(b)
	nln := int(readInt32(b)) // number of nodes in a single element (4: tetrahedron; 6: triangular prism; 8: hexahedron)
	if nln == 4 {
		log.Fatal("getElemPM read error: incompatible with tetrahedra")
	}
	if readInt32(b) != recb {
		log.Fatal("getElemPM read error1")
	}

	recb = readInt32(b)
	ne := int(readInt32(b)) // number of (3D) elements (prisms)
	if ne%nly != 0 {
		log.Fatal("getElemPM read error2")
	}
	epl := ne / nly
	if readInt32(b) != recb {
		log.Fatal("getElemPM read error3")
	}

	// elements (no re-ordering)
	recb = readInt32(b)
	elms := make([][]int, ne)
	// nxr := hgs.bottomUpNodeXRef()
	// exr := hgs.bottomUpElementXRef()
	if int(recb)/4/nln != ne {
		log.Fatal("getElemPM read error4")
	}
	for i := 0; i < ne; i++ {
		cns, lst := make([]int32, nln), make([]int, nln)
		if err := binary.Read(b, binary.LittleEndian, &cns); err != nil {
			log.Fatalln("getElemPM read elements failed: ", err)
		}
		if nln == 6 { // triangular prism (VTK order)
			for ii := 0; ii < nln/2; ii++ {
				// lst[ii] = nxr[int(cns[ii+nln/2])]
				lst[ii] = int(cns[ii+nln/2]) - 1
			}
			for ii := nln / 2; ii < nln; ii++ {
				// lst[ii] = nxr[int(cns[ii-nln/2])]
				lst[ii] = int(cns[ii-nln/2]) - 1
			}
		} else if nln == 8 { // quadrilateral prism
			for ii := 0; ii < nln/2; ii++ {
				// lst[ii] = nxr[int(cns[ii+nln/2])]
				lst[ii] = int(cns[ii+nln/2]) - 1
			}
			for ii := nln / 2; ii < nln; ii++ {
				// lst[ii] = nxr[int(cns[ii-nln/2])]
				lst[ii] = int(cns[ii-nln/2]) - 1
			}
		} else {
			log.Fatalf("getElemPM read todo, nln = %d", nln)
		}
		// elms[exr[i]] = lst
		elms[i] = lst
	}

	if readInt32(b) != recb {
		log.Fatal("getElemPM read error5")
	}

	/////////////////////////////////////////
	////  Note: skipping remaining info  ////
	/////////////////////////////////////////

	// fmt.Printf(" %d %d  %s\n", nln, hgs.Epl, fp)
	return elms, ne, epl
}
