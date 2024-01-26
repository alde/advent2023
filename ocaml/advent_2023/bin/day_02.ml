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

let max_dice color =
  match color with
  | "red" -> 12
  | "green" -> 13
  | "blue" -> 14
  | _ -> failwith "invalid color"
;;

let rec all lst =
  match lst with
  | [] -> true
  | h :: t -> if h then all t else false
;;

let to_rgb lst =
  let rec aux list (r, g, b) =
    match list with
    | [] -> r, g, b
    | (color, value) :: t ->
      (match color with
       | "red" -> (r + value, g, b) |> aux t
       | "green" -> (r, g + value, b) |> aux t
       | "blue" -> (r, g, b + value) |> aux t
       | _ -> failwith "invalid color")
  in
  aux lst (0, 0, 0)
;;

let process_round_part_one round =
  round
  |> String.split_on_char ','
  |> List.map @@ String.split_on_char ' '
  |> List.map (fun l -> last l, List.nth l 1 |> int_of_string)
  |> List.map (fun (color, value) -> max_dice color >= value)
  |> all
;;

let process_round_part_two round =
  round
  |> String.split_on_char ','
  |> List.map @@ String.split_on_char ' '
  |> List.map (fun l -> last l, List.nth l 1 |> int_of_string)
  |> to_rgb
;;

let extract_game_number line =
  line
  |> String.split_on_char ':'
  |> first
  |> String.split_on_char ' '
  |> last
  |> int_of_string
;;

let extract_rounds line ~processor =
  line
  |> String.split_on_char ':'
  |> last
  |> String.split_on_char ';'
  |> List.map processor
;;

let parse_game line =
  let game_number = extract_game_number line in
  let rounds = extract_rounds line ~processor:process_round_part_one in
  if rounds |> all then game_number else 0
;;

let max_dice_used_in_game dice_set =
  let rec aux list (r, g, b) =
    match list with
    | [] -> r, g, b
    | (r1, g1, b1) :: t -> aux t (max r r1, max g g1, max b b1)
  in
  aux dice_set (0, 0, 0)
;;

let parse_game_part_two line =
  extract_rounds line ~processor:process_round_part_two
  |> max_dice_used_in_game
  |> fun (r, g, b) -> r * g * b
;;

let solve input_file =
  let input = Utils.read_lines input_file in
  input
  |> List.map @@ parse_game
  |> List.fold_left ( + ) 0
  |> Printf.printf "Day Two - Part One: %d\n";
  input
  |> List.map parse_game_part_two
  |> List.fold_left ( + ) 0
  |> Printf.printf "Day Two - Part Two: %d\n"
;;
