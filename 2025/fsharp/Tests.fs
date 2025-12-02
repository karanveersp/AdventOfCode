module Tests

open FsUnit.Xunit
open Xunit

open Day1
open Helpers

[<Fact>]
let ``Sample input test`` () =
    read 1 true
    |> parse 
    |> getPassword 100
            
[<Fact>]
let ``Part 1 test`` () =
    read 1 false
    |> parse
    |> getPassword 100
    |> should equal 1084
 
[<Fact>]
let ``Sample input test part 2`` () =
    read 1 true
    |> parse
    |> getPasswordPart2 100
    |> should equal 6
                
 
[<Fact>]
let ``Part 2 test`` () =
    read 1 false
    |> parse
    |> getPasswordPart2 100
    |> should equal 6475
                