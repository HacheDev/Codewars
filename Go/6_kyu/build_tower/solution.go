package kata

import "strings"

func TowerBuilder(nFloors int) []string {
  
  var tower []string
  maxWidth := 1 + (nFloors - 1) * 2
  
  for i := 0; i < nFloors; i++  {
    floor := strings.Repeat(" ", (maxWidth - (1 + i * 2)) / 2)
    floor += strings.Repeat("*", 1 + i * 2)
    floor += strings.Repeat(" ", (maxWidth - (1 + i * 2)) / 2)
    tower = append(tower, floor)
  }
  
  return tower
}