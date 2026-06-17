# EDI CLI

>The main website for the EDI project can be visited [here](https://ediproject.org). 

This project is a command-line tool for using the core [EDI module](https://github.com/Chad-Glazier/edi) to analyze the programs that play the [Game of Amazons](https://en.wikipedia.org/wiki/Game_of_the_Amazons). Amazons has been historically studied and used for computer tournaments, but most of the existing research focuses on justifying and improving the authors' individual programs. In contrast, EDI is an effort to implement a variety of programs to directly compare them in terms of both raw performance (i.e., who wins more often), but also the more specific questions regarding the algorithms such as:
- What is the ideal tradeoff in terms of search depth versus evaluation strength?
- Which move ordering heuristics actually matter?
- Can strong Monte Carlo models be beaten by Alpha-Beta?

The core [`edi`](https://github.com/Chad-Glazier/edi) module implements the actual game-playing programs and the means to collect certain analytics, while this CLI tool is meant to run games between programs to collect and visualize statistics.

## Installation

If you want to watch some robots play the game in your terminal and help collect data, feel free to install this tool. At the time of writing, it can only be installed via the `go` CLI which you can download from [here](https://go.dev/). Once you have Go installed, you can get the program from here:

```sh
go install https://github.com/edi_cli/edi
```

Then you're ready to use the program:

```sh
edi
```
