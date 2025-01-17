module SolutionTests

open FsUnit.Xunit
open Xunit

open Solution 

[<Fact>]
let ``hello world returns hello world`` () = hello "world" |> should equal "hello world"
