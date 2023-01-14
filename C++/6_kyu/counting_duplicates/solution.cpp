size_t duplicateCount(const std::string& in); // helper for tests

size_t duplicateCount(const char* in)  {

    int count = 0;
    size_t length = strlen(in);
    int arr[256] {};
    
    std::cout << length;//@TESTING: for testing
    
    for(int i = 0; i < length; ++i)  {
    
      ++arr[in[i]];
    
    }
    
  for(int i = 0; i < 256; ++i)  {//@TESTING:for testing
    
    std::cout << arr[i];
    
  }
    
    for(int i = 65; i < 91; ++i)  {
    
      if(arr[i]+arr[i+32] > 1) ++count;
    
    }
    
    return count;
    
}