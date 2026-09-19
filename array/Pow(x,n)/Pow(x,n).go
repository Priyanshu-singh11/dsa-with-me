func myPow(x float64, n int) float64 {
    ans := 1.0;
    if n==0{ return 1.0}
    if x==0{return 0.0}
    if x==1{return 1.0}
    if x==-1 && n%2==0{return 1.0}
    if x==-1 && n%2!=0{ return -1.0}
    if n<0{
        x = 1/x
        n = -n
    }
    for n>0{
        if n%2 == 1 {
        ans = ans * x
    }
    x = x*x
    n = n/2
    }
    return ans 
}
