package main

import (
    "math/bits"
    "fmt"
)

type uint128 struct {
    H uint64
    L uint64
}

func (a uint128) Add128(b uint128) uint128 {
    l:= a.L + b.L
    
    carry := uint64(0)
    
    if l < a.L {
        carry = 1
    }
    
    return uint128{
        H: a.H + b.H + carry,
        L: l,
    }
}

func (a uint128) Sub128(b uint128) uint128 {
    l := a.L - b.L
    h := a.H - b.H
    
    if a.L < b.L {
        h -= 1
    }
    
    return uint128{H: h, L: l}
}


func (a uint128) Add32(x int32) uint128 {
    var b uint128
    b.L = uint64(uint32(x))
    
    if x < 0 {
        b.H = ^uint64(0) // (0xFFFFFFFFFFFFFFFF)
    }

    return a.Add128(b)
}

func (a uint128) Sub32(x int32) uint128 {
    return a.Add32(-x)
}


func (a uint128) Add64(x int64) uint128 {
    var b uint128
    b.L = uint64(x)
    
    if x < 0 {
        b.H = ^uint64(0) // (0xFFFFFFFFFFFFFFFF)
    }

    return a.Add128(b)
}

func (a uint128) Sub64(x int64) uint128 {
    return a.Add64(-x)
}

func (a uint128) Mul128(b uint128) uint128 {
    
    mul64 := func (x, y uint64) uint128 {
        h, l := bits.Mul64(x, y)
        return uint128{H: h, L: l}
    }

    a0, a1 := a.L, a.H
    b0, b1 := b.L, b.H

    loLo := mul64(a0, b0)
    loHi := mul64(a0, b1)
    hiLo := mul64(a1, b0)
    hiHi := mul64(a1, b1)

    // Add up with carries
    lo := loLo.L
    carry := loLo.H

    carry += loHi.L
    carry += hiLo.L

    hi := hiHi.L + loHi.H + hiLo.H + a1*b1 + carry

    return uint128{H: hi, L: lo}
}

func (a uint128) ShiftLeft(n uint) uint128 {
    if n >= 128 {
        return uint128{0, 0}
    } else if n == 0 {
        return a
    } else if n < 64 {
        return uint128{
            H: (a.H << n) | (a.L >> (64 - n)),
            L: a.L << n,
        }
    } else { // 64 <= n < 128
        return uint128{
            H: a.L << (n - 64),
            L: 0,
        }
    }
}

func (a uint128) ShiftRight(n uint) uint128 {
    if n >= 128 {
        return uint128{0, 0}
    } else if n == 0 {
        return a
    } else if n < 64 {
        return uint128{
            H: a.H >> n,
            L: (a.L >> n) | (a.H << (64 - n)),
        }
    } else { // 64 <= n < 128
        return uint128{
            H: 0,
            L: a.H >> (n - 64),
        }
    }
}

func (u uint128) String() (s string) {
    if u.H > 0 {
        s = fmt.Sprintf("0x%x%016x", u.H, u.L)
    } else {
        s = fmt.Sprintf("0x%x", u.L)
    }
    
    return s
}
