export class G964 {

    public static digPow = (n: number, p: number) => {
      let num: string = n.toString(10);
      let res: number = 0;
      let _i: number;//contador _i
      for(_i = 0; _i < num.length; ++_i)  {  //iterate n, do  (a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) 
        let aux: number = +num[_i];
        res += Math.pow(aux,p);
        ++p;
      }
      let ksol: number = -1;//final solution
      let _k: number;//contador _k
      for(_k = 0; ksol == -1 && _k <= Math.sqrt(res); ++_k)  {  //compare n*_k with res
        if(res/_k == n) ksol = _k;
      }
      return ksol;
    }
}