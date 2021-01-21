```
          MYSTERY OF SILVER
              MOUNTAIN
========================================


GOOD LUCK ON YOUR QUEST!. YOU ARE AT A
CROSSROADS AND YOU CAN GO E,W,

========================================


WHAT WILL YOU DO NOW 
?
```

This is a Go implementation of The Mystery of Silver Mountain, a text adventure game published by Usborne Publishing in 1984. It was published in a book in the form of terse BASIC code which you had to type out and save to tape or disc before being able to play. Obfuscation in the code helped avoid the player from knowing too much about the game before playing. 

A PDF copy of the book is still available from Usborne's website here: https://usborne.com/gb/books/computer-and-coding-books.

All credit must go to the authors: Chris Oxlade and Judy Tatchell, and copyright holders Usborne Publishing.

Re-implementing the BASIC code was nothing much more than a trip down memory lane, and a way to replay the game without requiring a BASIC interpreter. The Go code honours the original BASIC and makes no attempt to refactor to a more readable or maintainable form. All the obfuscation in the original is preserved, including the spaghetti of GOTOs and GOSUBs. That being said, some of the the subroutines and variable declarations have been reordered to account for Go's scoping rules.

Install via `git clone` and `go build`:
```
git clone https://github.com/fivegreenapples/go-mountain
cd go-mountain
go build -o silver-mountain
./silver-mountain
```

For Windows users:

```
go build -o silver-mountain.exe
silver-mountain
```

# Affordances for an Easier Life

Replicating the original program is all very well but does lead to some slightly frustrating game play. Perhaps that was part of the charm back in the eighties, but if memory serves, it was equally if not more frustrating back then.

For example, the `SAVE GAME` and `CONTINUE A SAVED GAME` operations both require typing out a filename plus a few other keystrokes. So just to save your progress and continue is a bit of a hassle. I've added a few niceties to make game play more pleasant:

1. Input commands are case insensitive.

   This just means there is a `strings.ToUpper(...)` wrapping any input as the BASIC only understands UPPERCASE COMMANDS.

2. You can specify a commands file.

    By writing a sequence of newline separated commands into a file, you can have the program run these before accepting manual input from STDIN. Specify the filename as the first command line argument to the program:

    ```
    ./silver-mountain commands.txt
    ```

    You can also comment your command file. Any line starting with a `#` will be ignored.

3. You can seed the RNG.

    Some elements of the game are generated at game start using a random number generator. If you use a command file as above, you'll want to seed the RNG so these game elements are predictable. Use a second command line argument:

    ```
    ./silver-mountain commands.txt 123
    ```

# Notes on Playing the Game

While trying not to give away any spoilers, it's worth noting that you can't complete this game without the book. The book contains some critical information that isn't available just from playing the game.

Good luck!
