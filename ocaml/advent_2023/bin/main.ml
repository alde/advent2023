let solutions day =
  match day with
  | "one" -> Day_01.solve
  | "two" -> Day_02.solve
  | _ -> failwith @@ day ^ " is not a valid day"
;;

let num_of day =
  match day with
  | "one" -> "01"
  | "two" -> "02"
  | _ -> failwith @@ day ^ " is not a valid day"
;;

let () =
  let num_args = Sys.argv |> Array.length in
  let day, input =
    match num_args with
    | 1 -> "all", "input"
    | 2 -> Sys.argv.(1), "input"
    | _ -> Sys.argv.(1), Sys.argv.(2)
  in
  match day with
  | "all" ->
    [ "one"; "two" ]
    |> List.iter (fun d ->
      "./inputs/day_" ^ num_of d ^ "/" ^ input ^ ".txt" |> solutions d)
  | day -> "./inputs/day_" ^ num_of day ^ "/" ^ input ^ ".txt" |> solutions day
;;
