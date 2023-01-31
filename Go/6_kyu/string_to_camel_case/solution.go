package kata

import(
  "strings"
  "regexp"
  )

func ToCamelCase(s string) string {
	// your code
	
  camelString := ""
  reg := regexp.MustCompile("[_-][a-zA-Z0-9]")
  rep := strings.NewReplacer("_", "", "-", "")
  camelString = reg.ReplaceAllStringFunc(s, strings.ToUpper)
  return rep.Replace(camelString)

}