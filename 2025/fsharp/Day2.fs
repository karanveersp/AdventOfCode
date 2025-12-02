module Day2

type Range = {
    Min: int64
    Max: int64
}

let parseDay2 (text: string) : Range list =
    text.Split(',')
    |> Array.map (fun s ->
        let parts = s.Split('-')
        { Min = int64 parts[0]; Max = int64 parts[1] })
    |> Array.toList

/// Day 2 Part 1
let findInvalidProductIDs (ranges: Range list) : int64 =
    let isInvalid (value: int64) =
        let s = string value
        let mid = s.Length / 2
        s[0..mid-1] = s[mid..]
    
    let checkRange r =
        seq { r.Min .. r.Max }
        |> Seq.filter isInvalid
        |> Seq.sum
    
    ranges
    |> Array.ofList
    |> Array.Parallel.sumBy checkRange

/// Day 2 Part 2
let findInvalidProductIDsP2 (ranges: Range list) : int64 =
    let isInvalid (value: int64) =
        // id is invalid if it has a repeating pattern
        let hasRepeatingPattern =
            let s = string value
            [1 .. s.Length - 1]
            |> List.exists (fun len ->
                let pattern = s.Substring(0, len)
                s.Replace(pattern , "") = "" && pattern <> s)
        hasRepeatingPattern
       
    let checkRange r =
        seq { r.Min .. r.Max }
        |> Seq.filter isInvalid
        |> Seq.sum
    
    ranges
    |> Array.ofList
    |> Array.Parallel.sumBy checkRange
 