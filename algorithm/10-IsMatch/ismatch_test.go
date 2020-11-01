package main

import (
    "testing"
)

// 对于s和p，逐一匹配. 当匹配s[i]和p[j]是，在匹配过程中，对于p有三种情况：
// 1. p[j]为正常字符
// 2. p[j]为字符'.'
// 3. p[j]为字符'*'
// 对于第一种情况，直接检查s[i]和p[j]是否相等即可
// 对于第二种情况，由于'.'可以匹配任意字符，直接匹配s[i+1]和p[j+1]即可
// 对于第三种情况，当p[j+1]为'*'是，说明p[j]的字符可以出现0次或多次：
// 出现0次时，p直接跳过两个字符，匹配s[i]和p[j+2]即可
// 出现至少一次时，检查s[i]是否等于p[j]，如果不相等匹配失败，如果相等，匹配s[i+1]和p[j]即可
// 结束条件：当p能匹配完s时，表示完全匹配
func IsMatchByRecursive(s, p string) bool {
    if len(p) <= 0 {
        return len(s) == 0
    }

    match := (len(s) > 0) && (p[0] == '.' || s[0] == p[0])
    if (len(p) > 1) && (p[1] == '*') {
        return IsMatchByRecursive(s, p[2:]) || (match && IsMatchByRecursive(s[1:], p))
    }
    return match && IsMatchByRecursive(s[1:], p[1:])
}

func IsMatchByMemo(s, p string) bool {
    memo := make([][]int, len(s)+1)
    for i := 0; i < len(s)+1; i++ {
        memo[i] = make([]int, len(p)+1)
    }
    return isMatchByMemo(0, 0, s, p, memo)
}

func isMatchByMemo(i, j int, s, p string, memo [][]int) bool {
    if memo[i][j] != 0 {
        return memo[i][j] == 1
    }

    if j >= len(p) {
        return i >= len(s)
    }
        match := (i < len(s)) && (p[j] == '.' || s[i] == p[j])
    if (j+1 < len(p)) && (p[j+1] == '*') {
        match = isMatchByMemo(i, j+2, s, p, memo) || (match && isMatchByMemo(i+1, j, s, p, memo))
    } else {
        match = match && isMatchByMemo(i+1, j+1, s, p, memo)
    }

    if match {
        memo[i][j] = 1
    } else {
        memo[i][j] = 0
    }
    return match
}

func IsMatchByDp(s, p string) bool {
    slen := len(s)
    plen := len(p)
    dp := make([][]bool, slen+1)
    for i := 0; i < slen+1; i++ {
        dp[i] = make([]bool, plen+1)
    }

    dp[slen][plen] = true
    for i := slen; i >= 0; i-- {
        for j := plen - 1; j >= 0; j-- {
            match := (i < slen) && (p[j] == '.' || s[i] == p[j])
            if (j+1 < plen) && (p[j+1] == '*') {
                dp[i][j] = dp[i][j+2] || (match && dp[i+1][j])
            } else {
                dp[i][j] = match && dp[i+1][j+1]
            }
        }
    }
    return dp[0][0]
}

func TestIsMatch(t *testing.T) {
    testIsMatch(IsMatchByRecursive, t)
    testIsMatch(IsMatchByMemo, t)
    testIsMatch(IsMatchByDp, t)
}

func testIsMatch(fn func(s, p string) bool, t *testing.T) {
    s := "abc"
    p := "abc"
    ans := fn(s, p)
    if !ans {
        t.Errorf("test s=%s p=%s fail", s, p)
    }

    s = "abc"
    p = "abd"
    ans = fn(s, p)
    if ans {
        t.Errorf("test s=%s p=%s fail", s, p)
    }

    s = "abc"
    p = "a.c"
    ans = fn(s, p)
    if !ans {
        t.Errorf("test s=%s p=%s fail", s, p)
    }

    s = "abc"
    p = "a.d"
    ans = fn(s, p)
    if ans {
        t.Errorf("test s=%s p=%s fail", s, p)
    }

    s = "abc"
    p = "d*a*bc"
    ans = fn(s, p)
    if !ans {
        t.Errorf("test s=%s p=%s fail", s, p)
    }

    s = ""
    p = ""
    ans = fn(s, p)
    if !ans {
        t.Errorf("test s=%s p=%s fail", s, p)
    }

    s = "abc"
    p = ""
    ans = fn(s, p)
    if ans {
        t.Errorf("test s=%s p=%s fail", s, p)
    }

    s = ""
    p = ".*"
    ans = fn(s, p)
    if !ans {
        t.Errorf("test s=%s p=%s fail", s, p)
    }
}
