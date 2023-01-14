function sumOfDigits(from,to){
  
  let digitString : string = "";
  for(from; from <= to; ++from)  {  //all nums in range
    digitString += from.toString();
  }
  
  let res : number = 0;
  for(let _i = 0; _i < digitString.length; ++_i)  {  //sum all nums
    res += Number.parseInt(digitString[_i],10);
  }
  
  return res;
  
}