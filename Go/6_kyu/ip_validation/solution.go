package kata

import "strings"
import "strconv"

func Is_valid_ip(ip string) bool {
  
  octets := strings.Split(ip, ".")
  
  if len(octets) != 4 {
    return false
  }
  
  valid := true
  
  for i := 0; i < len(octets) && valid; i++  {
    octet, err := strconv.Atoi(octets[i])
    
    if err != nil || len(octets[i]) > 1 && octets[i][0] == 48 || octet < 0 || octet > 255  {
      valid = false
    }
      
  }
	return valid
}

//We could also use ParseIP method from net package