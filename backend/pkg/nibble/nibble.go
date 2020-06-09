package nibble

// Takes a byte and breaks the bits into two nibbles
func Break(input byte) (byte, byte) {
	return input >> 4, (input & 0x0F)
}

// Takes two bytes and assembles the nibbles together
func Assemble(high byte, low byte) byte {
	return (high<<4 | low)
}
