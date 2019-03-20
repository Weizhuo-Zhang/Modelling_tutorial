#!/bin/bash
javac ConQuick.java
javac SeqQuick.java

# Iteration 1
for iteration_time in {1..10}
do
    echo "Iteration $iteration_time"
    echo "ConQuick"
    time java ConQuick
    echo "SeqQuick"
    time java SeqQuick
done

