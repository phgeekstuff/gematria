// Copyright 2018 PH Geek Stuff. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
    Package table/standard provides the most known transformation table, and 
    related code.
*/
package standard

import "fmt"

const (
    a = 1      // aleph
    b = 2      // bet
    g = 3      // guimel
    d = 4      // dalet
    h = 5      // hey
    v = 6      // vav
    z = 7      // zain
    ch = 8     // chet 
    t = 9      // tet
    y = 10     // yud
    k = 20     // kaf
    l = 30     // lamed
    m = 40     // mem and mem sofit
    n = 50     // nun and nun sofit
    s = 60     // samech
    ayin = 70  // ayin
    p = 80     // pe
    tz = 90    // tzadik
    ku = 100   // kuf
    r = 200    // resh
    sh = 300   // shin
    ta = 400   // taf
)

func PrintStandardTable() {
    fmt.Println(a)
}
