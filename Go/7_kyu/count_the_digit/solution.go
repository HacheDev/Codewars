package kata

import "strconv"
import "strings"

func NbDig(n int, d int) int {
  
  var count int
  digitString := strconv.Itoa(d)
  
  for i := 0; i <= n; i++ {
    numString := strconv.Itoa(i*i)
    count += strings.Count(numString, digitString)
  }
  
  return count
}