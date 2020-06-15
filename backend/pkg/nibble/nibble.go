package nibble

// Takes a byte and breaks the bits into two nibbles
func BreakSlice(input []byte) []byte {
	var result []byte
	for i := range input {
		high, low := Break(input[i])
		result = append(result, high, low)
	}
	return result
}

// Takes a byte and breaks the bits into two nibbles
func AssembleSlice(input []byte) []byte {
	var result []byte
	for i := 0; i < len(input); i = i + 2 {
		assembled := Assemble(input[i], input[i+1])
		result = append(result, assembled)
	}
	return result
}

// Takes a byte and breaks the bits into two nibbles
func Break(input byte) (byte, byte) {
	return input >> 4, (input & 0x0F)
}

// Takes two bytes and assembles the nibbles together
func Assemble(high byte, low byte) byte {
	return (high<<4 | low)
}
