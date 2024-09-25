package cipher

import (
    "strings"
    "unicode"
)

// Shift cipher type
type shift struct {
    distance int
}

// Vigenere cipher type
type vigenere struct {
    key string
}

// NewCaesar returns a shift cipher with a fixed distance of 3.
func NewCaesar() Cipher {
    return NewShift(3)
}

// NewShift returns a shift cipher with the specified distance.
func NewShift(distance int) Cipher {
    if distance == 0 || distance < -25 || distance > 25 {
        return nil
    }
    return shift{distance: distance}
}

// Encode for the shift cipher.
func (c shift) Encode(input string) string {
    return shiftText(input, c.distance)
}

// Decode for the shift cipher.
func (c shift) Decode(input string) string {
    return shiftText(input, -c.distance)
}

// Helper function to shift text for Shift Cipher
func shiftText(input string, distance int) string {
    var result strings.Builder
    for _, r := range input {
        if unicode.IsLetter(r) {
            shifted := shiftRune(unicode.ToLower(r), distance)
            result.WriteRune(shifted)
        }
    }
    return result.String()
}

// Helper function to shift a single rune
func shiftRune(r rune, distance int) rune {
    offset := int(r - 'a')
    return rune(((offset + distance + 26) % 26) + 'a')
}

// NewVigenere returns a Vigenere cipher with the given key.
func NewVigenere(key string) Cipher {
    if !isValidKey(key) || strings.Trim(key, "a") == "" {
        return nil
    }
    return vigenere{key: key}
}

// Encode for the Vigenere cipher.
func (v vigenere) Encode(input string) string {
    return vigenereText(input, v.key, true)
}

// Decode for the Vigenere cipher.
func (v vigenere) Decode(input string) string {
    return vigenereText(input, v.key, false)
}

// Helper function for Vigenere Cipher to handle encoding and decoding
func vigenereText(input, key string, encode bool) string {
    var result strings.Builder
    keyIndex := 0
    keyLength := len(key)
    for _, r := range input {
        if unicode.IsLetter(r) {
            shift := int(key[keyIndex] - 'a')
            if !encode {
                shift = -shift
            }
            result.WriteRune(shiftRune(unicode.ToLower(r), shift))
            keyIndex = (keyIndex + 1) % keyLength
        }
    }
    return result.String()
}

// Helper function to validate a Vigenere key
func isValidKey(key string) bool {
    for _, r := range key {
        if !unicode.IsLower(r) {
            return false
        }
    }
    return true
}
