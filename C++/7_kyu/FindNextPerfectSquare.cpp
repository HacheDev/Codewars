#include <math.h>

long int findNextSquare(long int sq) {

 long int rounded_root(round(sqrt(sq)));
 
  if(sq != rounded_root * rounded_root) return -1;
  
  long int root(sqrt(sq));
  ++root;

return root*root;

}