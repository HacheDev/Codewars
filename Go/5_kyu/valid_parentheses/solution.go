package kata

import "strings"

func ValidParentheses(parens string) bool {
  
  if len(parens) < 0 || len(parens) > 100 {
    panic("Input length must be between 0 and 100")
  }
  
  if len(parens) % 2 != 0 || strings.Count(parens, "(") != strings.Count(parens, ")") {
    return false
  }
  
  openedCounter := 0
  valid := true
  
  for i := 0; i < len(parens) && valid; i++  {
    if string(parens[i]) == "(" {
      openedCounter++
      
    } else if string(parens[i]) == ")"  {
      if openedCounter > 0  {
        openedCounter--
        
      } else  {
        valid = false
      }
      
    } else  {
      valid = false
    }
  }
  
  return valid
}