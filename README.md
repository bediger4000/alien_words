# Daily Coding Problem: Problem #1553 [Hard]

This has also appeared as problem 1203, and probably others.

## Problem Statement

This problem was asked by Airbnb.

You come across a dictionary of sorted words in a language you've never seen before.
Write a program that returns the correct order of letters in this language.

For example,
given `['xww', 'wxyz', 'wxyw', 'ywx', 'ywz']`,
you should return `['x', 'z', 'w', 'y']`.

## Analysis

The only way I can see to do this is to compare characters in adjacent pairs of words.
So compare word 0 and word 1, words 1 and 2, 2 and 3, etc.

If the first character of a pair of words is different,
`word[n][0]`  is lexically less than `word[n+m][0]`
Note that `n` is numerically less than `m`.
The initial letters of the sorted list of words have a "less than" relationship,
but there could be gaps.

If the first character of a pair of words is identical,
compare the second characters.
If the second characteres differ, the first word's second character
is lexically less than the second word's second character.
if  the second characters are identical, compare third characters,
and so on.

You get one lexically less than relationship of characters
between any pair of adjacent words,
but you can also say that the first character of a word is lexically less than
the first character of every word appearing before it in the sorted list.

### Data Structures

The data structure to keep the characters in will need to allow arbitrary insertions,
but not deletions.
The trick here is that even when a pair of characters have a lexical relationship,
we don't know if characters will show up that are lexically between them.
We also have to pick a data structure that allows a "less than" character to have
multiple characters "greater than".

I used a hashtable or Go `map` for the characters examined so far,
so access to a particular character's data structure is convenient.

### Evolution of data structure

At first glance, I thought a doubly-linked list would work:

```
type node struct {
    character rune
    next *node
    prev *node
}
```
After thinking a bit and writing a little code,
I realized that a character could have multiple children
during the procedure of examining pairs of words,
even if the final answer had one child character per character found.

I also thought I would need a root node, which I mistakenly
believed would be the first letter of the first word of the list.
That's not true, a word list of `['ba', 'bb', 'bc']` has a first letter
of first word that's not the first lexically sorted letter.

```
type node struct {
    character rune
    children  []*node
}
```

I also realized that to insert a newly encountered letter,
I'd need to have a "parent" pointer.

```
type node struct {
    character rune
    parent    *node
    children  []*node
}
```

Then I realized a Go `map` type would be easier coding all the slice manipulation
required to make a node a child and remove child nodes.

```
type node struct {
    character rune
    parent    *node
    children  map[rune]*node
}
```

In addition to a `map[rune]*node` to track all of the characters discovered,
I have a `

## Interview Analysis
