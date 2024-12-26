(def data (slurp "../input-1.txt"))
(def rows  (clojure.string/split-lines data))
(def cols  (map #(clojure.string/split % #"\s+") rows))
;; (println (flatten cols))
(def values  (map (fn [x] (map #(Integer/parseInt %) x)) cols))
;; (print values)

(def first_col (sort (map first values)))
(def second_col (sort (map second values)))
;; (print "\n========= First column: ==========\n" first_col)
;; (print "\n========= Second column: ==========\n" second_col)

(def distance (reduce + (map (fn [a b] (abs (- a b))) first_col second_col)))
(println "Distance: " distance) ; 1319616

(def arr2-counts
    (reduce (fn [counts num]
        (update counts num (fnil inc 0)))
        {}
        second_col))

;; (print arr2-counts)
;; (print (type arr2-counts)) # clojure.lang.PersistentHashMap

(def similarity
    (->> first_col
        (filter #(> (get arr2-counts % 0) 0))
        (map #(* % (arr2-counts %)))
        (reduce +)))

(println "Similarity: " similarity) ; 27267728
