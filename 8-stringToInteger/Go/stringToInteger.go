package stringToInteger


import (
	"math"
	"strings"
	"unicode"
)
func myAtoi(s string) int {
    s = strings.TrimSpace(s);
    if s == "" {	return 0; }
    negFlag := false;
    n := 0;
    nLen := 0;

    
    if s[0] == '-' || s[0] == '+'{
     if s[0] == '-'{
           negFlag = true;}

        s = s[1:]
    }
    for _, l := range s{
        if nLen > 12 {break}

        if unicode.IsDigit(l){
			l -= '0'
            n = (n * 10) + int(l)
            
            if(n > 0){nLen++}
        } else {
            break
        }
    }
    if negFlag {
        n = n * (-1)
    }
	if n > math.MaxInt32{
		n = math.MaxInt32
	}
	if n < -math.MaxInt32{
		n = -math.MaxInt32-1
	}
    return n 
}