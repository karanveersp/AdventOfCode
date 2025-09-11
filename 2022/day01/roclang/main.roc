app [main!] { pf: platform "https://github.com/roc-lang/basic-cli/releases/download/0.19.0/Hj-J_zxz7V9YurCSTFcFdu6cQJie4guzsPMUi5kBYUk.tar.br" }

import pf.Stdout

main! =
    Stdout.line! (add_and_stringify 1 3)

add_and_stringify : Num a, Num a -> Str
add_and_stringify = |a, b|
    Num.to_str(a + b)

# Tests
expect
    res = add_and_stringify(1, 3)
    res == "4"
