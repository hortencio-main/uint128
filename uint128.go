package Uint128

import (
    "math/bits"
    "fmt"
)

type Uint128 struct {
    H uint64
    L uint64
}

func (a Uint128) Add128(b Uint128) Uint128 {
    l:= a.L + b.L
    
    carry := uint64(0)
    
    if l < a.L {
        carry = 1
    }
    
    return Uint128{
        H: a.H + b.H + carry,
        L: l,
    }
}

func (a Uint128) Sub128(b Uint128) Uint128 {
    l := a.L - b.L
    h := a.H - b.H
    
    if a.L < b.L {
        h -= 1
    }
    
    return Uint128{H: h, L: l}
}


func (a Uint128) Add32(x int32) Uint128 {
    var b Uint128
    b.L = uint64(uint32(x))
    
    if x < 0 {
        b.H = ^uint64(0) // (0xFFFFFFFFFFFFFFFF)
    }

    return a.Add128(b)
}

func (a Uint128) Sub32(x int32) Uint128 {
    return a.Add32(-x)
}


func (a Uint128) Add64(x int64) Uint128 {
    var b Uint128
    b.L = uint64(x)
    
    if x < 0 {
        b.H = ^uint64(0) // (0xFFFFFFFFFFFFFFFF)
    }

    return a.Add128(b)
}

func (a Uint128) Sub64(x int64) Uint128 {
    return a.Add64(-x)
}

func (a Uint128) Mul128(b Uint128) Uint128 {
    
    mul64 := func (x, y uint64) Uint128 {
        h, l := bits.Mul64(x, y)
        return Uint128{H: h, L: l}
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

    return Uint128{H: hi, L: lo}
}

func (a Uint128) ShiftLeft(n uint) Uint128 {
    if n >= 128 {
        return Uint128{0, 0}
    } else if n == 0 {
        return a
    } else if n < 64 {
        return Uint128{
            H: (a.H << n) | (a.L >> (64 - n)),
            L: a.L << n,
        }
    } else { // 64 <= n < 128
        return Uint128{
            H: a.L << (n - 64),
            L: 0,
        }
    }
}

func (a Uint128) ShiftRight(n uint) Uint128 {
    if n >= 128 {
        return Uint128{0, 0}
    } else if n == 0 {
        return a
    } else if n < 64 {
        return Uint128{
            H: a.H >> n,
            L: (a.L >> n) | (a.H << (64 - n)),
        }
    } else { // 64 <= n < 128
        return Uint128{
            H: 0,
            L: a.H >> (n - 64),
        }
    }
}

func (u Uint128) String() (s string) {
    if u.H > 0 {
        s = fmt.Sprintf("0x%x%016x", u.H, u.L)
    } else {
        s = fmt.Sprintf("0x%x", u.L)
    }
    
    return s
}
