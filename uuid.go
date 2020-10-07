package arand

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// UUID will return hex of UUID (UUID: 4-2-2-2-6)
func UUID() []byte {
	// Prepare space
	outs := make([]byte, 36)
	outs[8], outs[13], outs[18], outs[23] = '-', '-', '-', '-'

	if outb := UUIDb(); outb == nil { // if failed to get UUIDb, then return nil
		return nil
	} else {
		// ==================================================================== String
		outs := make([]byte, 36)
		outs[8], outs[13], outs[18], outs[23] = '-', '-', '-', '-'
		hex.Encode(outs[0:8], outb[0:4])    // 4
		hex.Encode(outs[9:13], outb[4:6])   // 2
		hex.Encode(outs[14:18], outb[6:8])  // 2
		hex.Encode(outs[19:23], outb[8:10]) // 2
		hex.Encode(outs[24:], outb[10:])    // 6
		return outs
	}
}

// UUID will 16 bytes of UUID
func UUIDb() (uuidBytes []byte) {
	// This is standalone, simple UUIDs generation
	outb := make([]byte, 16)

	found := false
	for i := 0; i < 30; i++ { // up to 30 tries if error
		n, _ := io.ReadFull(rand.Reader, outb)
		if n == 16 {
			found = true
			break
		}
	}
	if !found { // tried 30 times to get rand.Reader but failed; return nil.
		return nil
	}

	outb[8] = outb[8]&^0xc0 | 0x80 // variant bits; see section 4.1.1
	outb[6] = outb[6]&^0xf0 | 0x40 // version 4 (pseudo-random); see section 4.1.3

	return outb
}
