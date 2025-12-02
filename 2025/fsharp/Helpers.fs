module Helpers

let read day sample =
    let filename =
        if sample then $"day%i{day}-sample.txt"
        else $"day%i{day}.txt"
    let path = System.IO.Path.Combine(__SOURCE_DIRECTORY__, "inputs", filename)
    System.IO.File.ReadAllText path
