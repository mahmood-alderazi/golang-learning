package piscine

func Slice(a []string, nbrs ...int) []string {
    size := len(a)
    
    start := 0
    if len(nbrs) > 0 {
        start = nbrs[0]
        if start < 0 {
            start += size
        }
    }
    
    end := size
    if len(nbrs) > 1 {
        end = nbrs[1]
        if end < 0 {
            end += size
        }
    }
    
    // Validation
    if start < 0 {
        start = 0
    }
    if end < 0 {
        end = 0
    }
    if start > size {
        start = size
    }
    if end > size {
        end = size
    }
    if start > end {
        return nil
    }
    
    return a[start:end]
}
