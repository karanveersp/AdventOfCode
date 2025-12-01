module Solution

open System.IO

// pw is count of how many times we end up at 0 during the sequence.
type Rotation =
    | Left of int
    | Right of int

let getPassword (rotations: Rotation list) n : int =
    // helper function to rotate in a direction one by one
    let rec rotate currLoc r =
        match r with
        | Left 0
        | Right 0 -> currLoc
        | Left times ->
            let loc = if currLoc - 1 < 0 then n - 1 else currLoc - 1 // wraparound
            rotate loc (Left(times - 1))
        | Right times ->
            let loc = if currLoc + 1 = n then 0 else currLoc + 1 // wraparound
            rotate loc (Right(times - 1))


    let rec aux rs currIdx count =
        match rs with
        | [] -> count
        | h :: t ->
            let newIndex = rotate currIdx h
            let count = if newIndex = 0 then count + 1 else count
            aux t newIndex count

    aux rotations (n / 2) 0


let getPasswordPart2 rotations n =
    let rec rotate currLoc zeroCount r =
        match r with
        | Left 0 -> (currLoc, zeroCount)
        | Right 0 -> (currLoc, zeroCount)
        | Left times ->
            let loc = if currLoc - 1 < 0 then n - 1 else currLoc - 1
            let zeroCount = if loc = 0 then zeroCount + 1 else zeroCount
            rotate loc zeroCount (Left(times - 1))
        | Right times ->
            let loc = if currLoc + 1 = n then 0 else currLoc + 1
            let zeroCount = if loc = 0 then zeroCount + 1 else zeroCount
            rotate loc zeroCount (Right(times - 1))


    let rec aux rs currIdx count =
        match rs with
        | [] -> count
        | h :: t ->
            let newIndex, zeroCount = rotate currIdx 0 h
            aux t newIndex (count + zeroCount)

    aux rotations (n / 2) 0



let parseInput (fileName: string) =
    let text = File.ReadAllText(fileName)

    let lineToRotation (line: string) : Rotation =
        if line.StartsWith("L") then
            Left(int (line.Substring(1)))
        else
            Right(int (line.Substring(1)))

    text.Split("\n") |> Array.map lineToRotation |> List.ofArray
