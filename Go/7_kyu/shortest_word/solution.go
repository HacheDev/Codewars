package kata

import "strings"

func FindShort(s string) int {
  
  words := strings.Split(s, " ")
  shortestLength := len(words[0])
  
  for i := 1; i < len(words); i++ {
    currLength := len(words[i])
    if currLength < shortestLength  {
      shortestLength = currLength
    }
  }

  return shortestLength
}