namespace Triangle
{
  bool isTriangle(long int a, long int b, long int c)
  {
    return (a+b)>c && (a+c)>b && (b+c)>a;
  }
};