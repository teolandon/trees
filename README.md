# treeutils

A Go package to print trees in a readable way.

# Usage

None yet, you could clone it and try running it.

# TODO

- Make it work with any type that implements Stringer. Should be fairly simple

- Define a new Tree type, the one used here is just for Tour of Go. Add useful
    methods to it.

- Consider concurrency. Left-right recursion need not happen sequentially.

- Consider dynamically changing the size of tree branches. Currently a single
    long string can make the whole tree bulky.
