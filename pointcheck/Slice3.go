package piscine

func Slice(a []string, nbrs ...int) []string {
    length := len(a)
    
    start := 0
    if len(nbrs) > 0 {
        start = nbrs[0]
        if start < 0 {
            start += length
        }
    }
    
    end := length
    if len(nbrs) > 1 {
        end = nbrs[1]
        if end < 0 {
            end += length
        }
    }
    
    // Validation
    if start < 0 {
        start = 0
    }
    if end < 0 {
        end = 0
    }
    if start > length {
        start = length
    }
    if end > length {
        end = length
    }
    if start > end {
        return nil
    }
    
    return a[start:end]
}
