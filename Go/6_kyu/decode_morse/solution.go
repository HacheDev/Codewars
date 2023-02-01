package kata

import "strings"

//@TODO: check when there are more than 3 whitespaces in a row

func DecodeMorse(morseCode string) string {
  
  morseWords := strings.Split(morseCode, "   ")
  for i := 0; i < len(morseWords); i++ {
    morseChars := strings.Split(morseWords[i], " ")
    morseWords[i] = MORSE_CODE[morseChars[0]]
    
    for j := 1; j < len(morseChars); j++  {
      morseWords[i] += MORSE_CODE[morseChars[j]]
    }
  }
  return strings.Join(morseWords, " ")
}
