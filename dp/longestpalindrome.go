func longestPalindrome(s string) string {
    m := make(map[string]bool)
    result := "" 
    for i :=0;i<len(s)-1;i++ {
        for j:=i;j<len(s);j++{
            if(isPalindrome(i,j,s,m)){
                c := s[i:j]
                if(len(result) < len(c)){
                    result = c    
                }    
            }
        }
    }
    
    return result
}

func isPalindrome(i int,j int, s string, m map[string]bool) bool {
    if i <= j {
        
    buffer1 := bytes.NewBufferString(strconv.Itoa(i))
    buffer1.WriteString("$")
    buffer1.WriteString(strconv.Itoa(j))
    lkey :=  buffer1.String()    
    _,ok := m[lkey];
        
    if ok {
        return m[lkey]
    }
        a := i+1
        b := j+1 
        buffer := bytes.NewBufferString(strconv.Itoa(a))
        buffer.WriteString("$")
        buffer.WriteString(strconv.Itoa(b))
        key :=  buffer.String()
        subPalindrome := false
        _,keyExists := m[key]
        if(keyExists){
            subPalindrome = m[key] 
        } else {
            isPalindrome(a,b,s,m)
        } 
        
        result := (s[i]==s[j]) && subPalindrome
        
        m[lkey] = result
        
        return result 
    }
    
    return true
}
