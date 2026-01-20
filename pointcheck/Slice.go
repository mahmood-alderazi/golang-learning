package piscine

func Slice(a []string, nbrs ...int) []string {
    n := len(a)
    start := 0
    end := n
    
    // Step 1: Get start
    if len(nbrs) > 0 {
        start = nbrs[0]
        if start < 0 {
            start = n + start
        }
    }
    
    // Step 2: Get end
    if len(nbrs) > 1 {
        end = nbrs[1]
        if end < 0 {
            end = n + end
        }
    }
    
    // Step 3: Fix and return
    if start < 0 {
        start = 0
    }
    if end < 0 {
        end = 0
    }
    if start > n {
        start = n
    }
    if end > n {
        end = n
    }
    if start > end {
        return nil
    }
    
    return a[start:end]
}
