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

This is a Go implementation of The Mystery of Silver Mountain, a text adventure game published by Usborne Publishing in 1984. It was published in a book in the form of terse BASIC code that you had to type out and save to tape or disc before being able to play. Obfuscation in the code helped avoid the player from knowing too much about the game before playing. 

A PDF copy of the book is still available from Usborne's website here: https://usborne.com/gb/books/computer-and-coding-books.

All credit must go to the authors: Chris Oxlade and Judy Tatchell, and copyright holders Usborne Publishing.

Re-implementing the BASIC code was nothing much more than a trip down memory lane, and a way to replay the game without requiring a BASIC interpreter. The Go code honours the original BASIC and makes no attempt to refactor to a more readable or maintainable form. All the obfuscation in the original is preserved, including the spaghetti of GOTOs and GOSUBs.

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