(ns cladvent2023.day_01.core
  (:require [cladvent2023.shared.input :as aoc]
            [clojure.string :as str]))

(def text-digits ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"])

(defn first-and-last [input]
  [(first input), (last input)])

(defn replace-at [s [idx replacement]]
  (if (some? idx)
    (str (subs s 0 idx) replacement (subs s (inc idx))) (identity s)))

(defn replace-several [line translations]
  (reduce #(replace-at %1 %2) line translations))

(defn get-stringified-locs [input]
  (let [full (map (fn [txt]
                    [(str/index-of input txt), (+ (.indexOf text-digits txt) 1)]) text-digits)
        filtered (filter some? full)
        res (replace-several input filtered)]
    res))

(defn calibrate [input]
  (->> input
       (map str)
       (map parse-long)
       (filter some?)
       (first-and-last)
       (reduce str)
       (parse-long)))

(defn solve_1 [filename]
  (->> filename
       (aoc/read-from-file)
       (map calibrate)
       (reduce +)))

(defn solve_2 [filename]
  (->> filename
       (aoc/read-from-file)
       (map get-stringified-locs)
       (map calibrate)
       (reduce +)))

(comment
  (= 142 (solve_1 "samples/day_one_part_one.txt"))
  (= 281 (solve_2 "samples/day_one_part_two.txt"))

  (solve_1 "inputs/day_one.txt")
  (solve_2 "inputs/day_one.txt"))


