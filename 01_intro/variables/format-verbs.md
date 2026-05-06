# types notes and formattings verbs

## types 

## inteigers signed 

| Type  | Size                                                                         | Range                                                                                      |
| ----- | ---------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------ |
| int   | Depends on platform:<br>32 bits (32-bit systems)<br>64 bits (64-bit systems) | -2147483648 to 2147483647 (32-bit)<br>-9223372036854775808 to 9223372036854775807 (64-bit) |
| int8  | 8 bits / 1 byte                                                              | -128 to 127                                                                                |
| int16 | 16 bits / 2 bytes                                                            | -32768 to 32767                                                                            |
| int32 | 32 bits / 4 bytes                                                            | -2147483648 to 2147483647                                                                  |
| int64 | 64 bits / 8 bytes                                                            | -9223372036854775808 to 9223372036854775807                                                |

## Unsigned 

| Type   | Size                                                                         | Range                                                          |
| ------ | ---------------------------------------------------------------------------- | -------------------------------------------------------------- |
| uint   | Depends on platform:<br>32 bits (32-bit systems)<br>64 bits (64-bit systems) | 0 to 4294967295 (32-bit)<br>0 to 18446744073709551615 (64-bit) |
| uint8  | 8 bits / 1 byte                                                              | 0 to 255                                                       |
| uint16 | 16 bits / 2 bytes                                                            | 0 to 65535                                                     |
| uint32 | 32 bits / 4 bytes                                                            | 0 to 4294967295                                                |
| uint64 | 64 bits / 8 bytes                                                            | 0 to 18446744073709551615                                      |


## Float 

| Type    | Size    | Range                  |
| ------- | ------- | ---------------------- |
| float32 | 32 bits | -3.4e+38 to 3.4e+38    |
| float64 | 64 bits | -1.7e+308 to +1.7e+308 |


## formattings verbs

| Verb   | Description                              |
| ------ | ---------------------------------------- |
| %v	 | Prints the value in the default format   |
| %#v	 | Prints the value in Go-syntax format     |
| %T	 | Prints the type of the value             |
| %%	 | Prints the % sign                        |

## inteigers
| Verb | Description                                |
| ---- | ------------------------------------------ |
| %b   | Base 2                                     |
| %d   | Base 10                                    |
| %+d  | Base 10 and always show sign               |
| %o   | Base 8                                     |
| %O   | Base 8, with leading 0o                    |
| %x   | Base 16, lowercase                         |
| %X   | Base 16, uppercase                         |
| %#x  | Base 16, with leading 0x                   |
| %4d  | Pad with spaces (width 4, right justified) |
| %-4d | Pad with spaces (width 4, left justified)  |
| %04d | Pad with zeroes (width 4)                  |

## String

| Verb | Description                                                 |
| ---- | ----------------------------------------------------------- |
| %s   | Prints the value as plain string                            |
| %q   | Prints the value as a double-quoted string                  |
| %8s  | Prints the value as plain string (width 8, right justified) |
| %-8s | Prints the value as plain string (width 8, left justified)  |
| %x   | Prints the value as hex dump of byte values                 |
| % x  | Prints the value as hex dump with spaces                    |

## Boolean

| Verb | Description                                  |
| ---- | -------------------------------------------- |
| %t   | Value of the boolean in true or false format |

## Float

| Verb  | Description                               |
| ----- | ----------------------------------------- |
| %e    | Scientific notation with 'e' as exponent  |
| %f    | Decimal point, no exponent                |
| %.2f  | Default width, precision 2                |
| %6.2f | Width 6, precision 2                      |
| %g    | Exponent as needed, only necessary digits |
