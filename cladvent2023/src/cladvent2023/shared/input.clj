(ns cladvent2023.shared.input
  (:require [clojure.java.io :as io]
            [clojure.string :as str]))

(defn read-from-file [filename]
  (->> filename
       (io/resource)
       (slurp)
       (str/split-lines)))
