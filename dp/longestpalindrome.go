func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
    
    l := len(s) 
    m := make([][]bool,l)
    for i := 0;i<l;i++ {
        m[i] = make([]bool,l)
    }
    
	result := ""
    lr := 0
    for i := 0; i < l-1; i++ {
        for j := i+lr; j < l; j++ {
			if isPalindrome(i, j, s, m) {
				c := s[i : j+1]
				if len(result) < len(c) {
					result = c
                    lr = len(result)
				}
			}
		}
	}

	return result
}

func isPalindrome(i int, j int, s string, m [][]bool) bool {
	if i <= j {

		if m[i][j] {
            return m[i][j]
        }
        
		a := i + 1
		b := j - 1
		
        subPalindrome := false 
		
        if(a < 0 || b < 0 || a > b){
            subPalindrome = true
        } else {
            if m[a][b] {
                subPalindrome = m[a][b]
            } else {
                subPalindrome = isPalindrome(a, b, s, m)
            }
        }

		result := (s[i] == s[j]) && subPalindrome
        
        if(a >= 0 && b >= 0 && a <= b){
            m[a][b] = result
        } 

		return result
	}

	return true
}
