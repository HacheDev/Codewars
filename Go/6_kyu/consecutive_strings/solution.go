package kata

import "strings"

func LongestConsec(strarr []string, k int) string {
  
  var longest string
  
  for i := 0; i < len(strarr) - (k - 1); i++  {
    
    var subArray [] string
    for j := 0; j < k; j++  {
       subArray = append(subArray, strarr[i + j])
    }
    
    concatString := strings.Join(subArray, "")
    
    if len(concatString) > len(longest) {
      longest = concatString
    }
  }
  
  return longest
}