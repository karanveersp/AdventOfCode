module Tests
open System
open System.IO
open Xunit
open Solution

[<Fact>]
let ``Sample input test`` () =
    let rotations = parseInput "input-sample.txt"
    let res = getPassword rotations 100
    Assert.Equal(3, res)
    
            
[<Fact>]
let ``Part 1 test`` () =
    let rotations = parseInput "input.txt"
    let res = getPassword rotations 100
    Assert.Equal(1084, res)
 
[<Fact>]
let ``Sample input test part 2`` () =
    let rotations = parseInput "input-sample.txt"
    let res = getPasswordPart2 rotations 100
    Assert.Equal(6, res)
                
 
[<Fact>]
let ``Part 2 test`` () =
    let rotations = parseInput "input.txt"
    let res = getPasswordPart2 rotations 100
    Assert.Equal(6475, res)
                