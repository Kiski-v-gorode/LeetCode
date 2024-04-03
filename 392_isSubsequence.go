package main

// The isSubsequence function takes two strings, s and t, and returns true
// if string s is a subsequence of string t, and false otherwise.
func isSubsequence(s string, t string) bool {
    // If the length of string s is greater than the length of string t,
    // s cannot be a subsequence of t, return false.
    if len(s) > len(t) {
        return false
    }
    
    // Use two pointers to compare characters in strings s and t.
    sPointer, tPointer := 0, 0
    for sPointer < len(s) && tPointer < len(t) {
        // If the current characters of s and t are equal, increase both pointers.
        if s[sPointer] == t[tPointer] {
            sPointer++
            tPointer++
        } else {
            // If the current characters are not equal, increase the t pointer
            // to check the next character in t.
            tPointer++
        }
    }
    // If all characters of string s were found in t, then s is a subsequence of t, return true.
    return sPointer == len(s)
}
