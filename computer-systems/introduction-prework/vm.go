package vm

import (
	"fmt"
)

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
//
func compute(memory []byte) {

	registers := [3]byte{8, 0, 0} // PC, R1 and R2
	// Keep looping, like a physical computer's clock
	for {
		fmt.Printf("memory: %v\n", memory)
		fmt.Printf("registers: pc: %d r1: %d r2: %d\n", registers[0], registers[1], registers[2])

		// Fetch (extract instruction at memory address of PC and additonal op args)
		op := memory[registers[0]]
		oparg1 := memory[registers[0]+1]
		oparg2 := memory[registers[0]+2]
		// incremement pc
		registers[0] += 3
		// decode and execute
		switch op {
		case Load:
			fmt.Printf("Load %d %d\n", oparg1, oparg2)
			registers[oparg1] = memory[oparg2]
		case Store:
			fmt.Printf("Store %d %d\n", oparg1, oparg2)
			memory[oparg2] = registers[oparg1]
		case Add:
			fmt.Printf("Add %d %d\n", oparg1, oparg2)
			registers[oparg1] += registers[oparg2]
		case Addi:
			fmt.Printf("Addi %d %d\n", oparg1, oparg2)
			registers[oparg1] += oparg2
		case Sub:
			fmt.Printf("Sub %d %d\n", oparg1, oparg2)
			// Possible bug? Could oparg1 ever be r2 for some reason?
			registers[oparg1] -= registers[oparg2]
		case Subi:
			fmt.Printf("Subi %d %d\n", oparg1, oparg2)
			// Possible bug? Could oparg1 ever be r2 for some reason?
			registers[oparg1] -= oparg2
		case Jump:
			fmt.Printf("Jump %d\n", oparg1)
			registers[0] = oparg1
		case Beqz:
			fmt.Printf("Beqz %d %d\n", oparg1, oparg2)
			if registers[oparg1] == 0 {
				registers[0] += oparg2
			}
		case Halt:
			fmt.Println("Halt")
			return
		}
	}
}
