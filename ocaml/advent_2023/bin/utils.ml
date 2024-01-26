let rec printlines data =
  match data with
  | [] -> None
  | x :: tl ->
    print_endline x;
    printlines tl
;;

let print_tuple (a, b, c) = Printf.printf "(%d, %d, %d)\n" a b c
let print_2_tuple (a, b) = Printf.printf "(%s, %d)\n" a b

let read_lines file =
  let contents = In_channel.with_open_bin file In_channel.input_all in
  String.split_on_char '\n' contents |> List.filter (fun line -> String.length line > 0)
;;
