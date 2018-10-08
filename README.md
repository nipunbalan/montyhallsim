# Monty Hall Problem Simulator in Go

## What is the problem?

[Monty Hall problem](https://en.wikipedia.org/wiki/Monty_Hall_problem) is a probability puzzle loosely based on the American television game show _[Let's Make a Deal](https://en.wikipedia.org/wiki/Let%27s_Make_a_Deal "Let's Make a Deal")_ and named after its original host, [Monty Hall](https://en.wikipedia.org/wiki/Monty_Hall "Monty Hall").

"Suppose you're on a game show, and you're given the choice of three doors: Behind one door is a car; behind the others, goats. You pick a door, say No. 1, and the host, who knows what's behind the doors, opens another door, say No. 3, which has a goat. He then says to you, "Do you want to pick door No. 2?" Is it to your advantage to switch your choice?"

It can be proved that the winning strategy is always to switch your choice. It can be proved using [various methods](https://en.wikipedia.org/wiki/Monty_Hall_problem)


## Simulator
This simulator program runs the scenario many times and shows you how many times the each strategy wins
```
go run monty.go <number of trials>

go run monty.go 100000
```

## References
[21](https://www.imdb.com/title/tt0478087/) movie
[Wikipedia page on Monty Hall Problem](https://en.wikipedia.org/wiki/Monty_Hall_problem)
[Wolfram Mathworld](http://mathworld.wolfram.com/MontyHallProblem.html)
[Prof. Lisa Goldberg's explanation](https://www.youtube.com/watch?v=4Lb-6rxZxx0)
[Prof. Lisa Goldberg's explanation with math proof](https://www.youtube.com/watch?v=ugbWqWCcxrg)
[Monty Hall Problem for Dummies - Numberphile](https://www.youtube.com/watch?v=7u6kFlWZOWg)
