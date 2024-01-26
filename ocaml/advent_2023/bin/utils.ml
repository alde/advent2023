let rec printlines data =
  match data with
  | [] -> None
  | x :: tl ->
    print_endline x;
    printlines tl
;;

let first l =
  match l with
  | [] -> failwith "empty list"
  | h :: _ -> h
;;

let last l =
  match List.rev l with
  | [] -> failwith "empty list"
  | h :: _ -> h
;;

let read_lines file =
  let contents = In_channel.with_open_bin file In_channel.input_all in
  String.split_on_char '\n' contents |> List.filter (fun line -> String.length line > 0)
;;
