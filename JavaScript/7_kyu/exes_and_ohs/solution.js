function XO(str) {

    var x = 0, o = 0;
    for(let i = 0; i < str.length; ++i)  {
    
      if (str.charAt(i) == "x" || str.charAt(i) == "X")  {
      
        ++x;
      
      }
      
      else if (str.charAt(i) == "o" || str.charAt(i) == "O")  {
      
        ++o;
      
      }
      
    }
    
    return x == o;
    
}