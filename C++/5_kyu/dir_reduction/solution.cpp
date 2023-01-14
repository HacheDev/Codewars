class DirReduction
{
public:
    static std::string oppDirection(std::string dir) {
    
    if(dir == "NORTH")  return "SOUTH";
    else if(dir == "EAST")  return "WEST";
    else if(dir == "WEST")  return "EAST";
    else if(dir == "SOUTH")  return "NORTH";
   
  }
    
public:
    static std::vector<std::string> dirReduc(std::vector<std::string> &arr)  {
    
    std::vector<std::string> &path = arr;
    
      for(int i = 0; i < path.size(); ++i)  {
      
         std::string dir = path[i];
         std::string opDir = oppDirection(dir);
         
         if(path[i+1] == opDir)  {
           path.erase(path.begin() + (i+1));
           path.erase(path.begin() + (i));
           dirReduc(path);
         }

      }
      
      return path;
    
    }
    
    
};