// Cue has syntactic sugar for writing numbers too:
//
//	- 0x, 0o, 0b for hex, octal, and binary
//	- K, M, G, T, P for sizing with optional i
//	- e/E for decimal exponents
//	- Underscores for visual readability for large numbers

hex:  0xdeadbeef
oct:  0o755
bin:  0b0101_0001
cpu:  0.5Mi
mem:  4Gi
mill: 1M
bill: 1G
zero: 0.0
half: 0.5
trim: 01.23
mole: 6.022_140_76e+23
tiny: 1.2345e-12
long: 23_456_789_000_000000