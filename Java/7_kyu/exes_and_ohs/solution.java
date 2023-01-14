public class XO {
  
    public static boolean getXO (String str) {
      
      int xAmount = 0, oAmount = 0;
      
      for(int i = 0; i < str.length(); ++i) {
      
        if(str.charAt(i) == 'x' || str.charAt(i) == 'X') {
        
          ++xAmount;
        
        }
        
        else if(str.charAt(i) == 'o' || str.charAt(i) == 'O') {
        
          ++oAmount;
        
        }
        
      }
      
      return xAmount == oAmount;
      
    }
  }