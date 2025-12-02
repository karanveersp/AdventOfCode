module Tests

open FsUnit.Xunit
open Xunit

open Day1
open Day2
open Helpers

[<Fact>]
let ``Day1 Sample`` () =
    read 1 true
    |> parseDay1
    |> getPassword 100
    |> should equal 3
            
[<Fact>]
let ``Day1 Part 1`` () =
    read 1 false
    |> parseDay1
    |> getPassword 100
    |> should equal 1084
 
[<Fact>]
let ``Day1 Part 2 Sample`` () =
    read 1 true
    |> parseDay1
    |> getPasswordPart2 100
    |> should equal 6
                
 
[<Fact>]
let ``Day1 Part 2`` () =
    read 1 false
    |> parseDay1
    |> getPasswordPart2 100
    |> should equal 6475
                
[<Fact>]
let ``Day2 Sample`` () =
    read 2 true
    |> parseDay2
    |> findInvalidProductIDs
    |> should equal 1227775554L
     
[<Fact>]
let ``Day2 Part 1`` () =
    read 2 false
    |> parseDay2
    |> findInvalidProductIDs
    |> should equal 9188031749L

[<Fact>]
let ``Day2 Part 2 Sample`` () =
    read 2 true
    |> parseDay2
    |> findInvalidProductIDsP2
    |> should equal 4174379265L


[<Fact>]
let ``Day2 Part 2`` () =
    read 2 false
    |> parseDay2
    |> findInvalidProductIDsP2
    |> should equal 11323661261L
    
