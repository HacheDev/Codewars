package kata


func ArrayDiff(a, b []int) []int {
  
  for i := 0; i < len(a); i++ {
    if contains(b, a[i]) {
      a = remove(a, i)
      i--
    }
  }
  
  return a
  
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}

func contains(slice []int, elem int) bool  {
  
  isContained := false
  for i := 0; i < len(slice) && !isContained; i++ {
    if slice[i] == elem {
      isContained = true
    }
  }
  
  return isContained
}