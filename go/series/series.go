package series

func All(n int, s string) []string {
    var result []string

    if n > len(s) {
        return result
    }

    for i := 0; i <= len(s)-n; i++ {
        result = append(result, s[i:i+n])
    }

    return result
}

func UnsafeFirst(n int, s string) string {
    if n > len(s) {
        return ""
    }
    return s[:n]
}

func First(n int, s string) (first string, ok bool) {
    if n > len(s) {
        return "", false
    }
    return s[:n], true
}