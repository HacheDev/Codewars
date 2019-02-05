#include <math.h>

bool is_square(int n)
{
  
  if (n < 0) return false;
  
  int root(round(sqrt(n)));
  return n == root * root;
  
}