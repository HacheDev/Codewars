#include<vector>

unsigned int number(const std::vector<std::pair<int, int>>& busStops){

  unsigned int people(0);

  for(int i = 0; i < busStops.size(); ++i) {
  
    people += busStops[i].first - busStops[i].second;
      
  }
  
  return people;

}