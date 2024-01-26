let is_number c =
  match c with
  | '0' .. '9' -> true
  | _ -> false
;;

let string_to_char_list s = s |> String.to_seq |> List.of_seq

let rec process_line line ~ef ~left ~right =
  match line with
  | [] -> (left * 10) + right
  | head :: tail ->
    let num = ef (head :: tail) in
    if num > 0
    then
      if left == 0
      then process_line tail ~ef ~left:num ~right:num
      else process_line tail ~ef ~left ~right:num
    else process_line tail ~ef ~left ~right
;;

let get_num x =
  match x with
  | [] -> 0
  | x :: _ when is_number x -> Char.code x - 48
  | _ -> 0
;;

let get_num_with_word (line : char list) =
  match line with
  | [] -> 0
  | 'o' :: 'n' :: 'e' :: _ -> 1
  | 't' :: 'w' :: 'o' :: _ -> 2
  | 't' :: 'h' :: 'r' :: 'e' :: 'e' :: _ -> 3
  | 'f' :: 'o' :: 'u' :: 'r' :: _ -> 4
  | 'f' :: 'i' :: 'v' :: 'e' :: _ -> 5
  | 's' :: 'i' :: 'x' :: _ -> 6
  | 's' :: 'e' :: 'v' :: 'e' :: 'n' :: _ -> 7
  | 'e' :: 'i' :: 'g' :: 'h' :: 't' :: _ -> 8
  | 'n' :: 'i' :: 'n' :: 'e' :: _ -> 9
  | x :: _ when is_number x -> Char.code x - 48
  | _ -> 0
;;

let rec process_lines data ~extractor_func ~acc =
  match data with
  | [] -> acc
  | head :: tail ->
    process_lines
      tail
      ~extractor_func
      ~acc:
        (head |> string_to_char_list |> process_line ~ef:extractor_func ~left:0 ~right:0)
    + acc
;;

let solve input_file =
  let input = Utils.read_lines input_file in
  input
  |> process_lines ~extractor_func:get_num ~acc:0
  |> Printf.printf "Day One - Part One: %d\n";
  input
  |> process_lines ~extractor_func:get_num_with_word ~acc:0
  |> Printf.printf "Day One - Part Two: %d\n"
;;
