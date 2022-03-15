package utils

func Filter(vs []string, f func(string) bool) []string {
    filtered := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            filtered = append(filtered, v)
        }
    }
    return filtered
}
