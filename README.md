# Daily Coding Problem: Problem #1553 [Hard]

I'm pretty sure this problem has shown up before.

## Problem Statement

This problem was asked by Airbnb.

You come across a dictionary of sorted words in a language you've never seen before.
Write a program that returns the correct order of letters in this language.

For example,
given `['xww', 'wxyz', 'wxyw', 'ywx', 'ywz']`,
you should return `['x', 'z', 'w', 'y']`.

## Analysis

The only way I can see to do this is to compare pairs of words.

If the first character of a pair of words is different,
`word[n][0]`  is lexically less than `word[n+1][0]`

If the first character of a pair of words is identical,
compare the second characters.
If the second characteres differ, the first word's second character
is lexically less than the second word's second character.
if  the second characters are identical, compare third characters,
and so on.

The first character of the first word has to be the "lexically least" letter of alle

The data structure to keep the characters in will need to allow arbitrary insertions,
but not deletions.
Also need to keep track of which characters have already been seen
The trick here is that even when a pair of characters have a lexical relationship,
we don't know if there are characters that are lexically between them.
We have to pick a data structure that allows a "less than" character to have
multiple characters "greater than".

We also need a hashtable or Go `map` for the characters seen so far.

The Go `list` standard package is too general, has too many methods.
I'll write my own, which will have minimal methods.

### Evolution of data struct

```
type node struct {
    character rune
    next *node
    prev *node
}
```

```
type node struct {
    character rune
    children  []*node
}
```

```
type node struct {
    character rune
    parent    *node
    children  []*node
}
```

```
type node struct {
    character rune
    parent    *node
    children  map[rune]*node
}
```

## Interview Analysis
