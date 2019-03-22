#include <string>

std::string highestScoringWord(const std::string &str)
{
  
  int val = 0;
  std::string word = "";
  
  for(int i = 0; i < str.length(); )  {
  
    int valAux = 0;
    int j = 0;
    
    for(j = i; j < str.length() && (int)str[j] != 32; ++j)  {
      valAux += (int)str[j] - 96;
    }
    
    if(valAux > val){
      val = valAux;
      word = str.substr(i,j-i);
    }
    i = j + 1;

  }
  return word;
}
