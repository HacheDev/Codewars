export class G964 {

    public static digPow = (n: number, p: number) => {
      let num: string = n.toString(10);
      let res: number = 0;
      let _i: number;//contador _i
      for(_i = 0; _i < num.length; ++_i)  {  //iterate n, do  (a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) 
        let aux: number = +num[_i];
        res += aux**p;
        ++p;
      }
      return Number.isInteger(res/n) ? res/n : -1;
    }
}