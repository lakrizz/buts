# go-buts

This is an implementation of a bounded, unique, timeout Stack, which means that this is a simple Stack with the following properties:
- Limited Bounds (e.g., it has a capacity which will not be exceeded, overflowing items will be discarded from the bottom of the stack)
- Unique, items can of any kind but can't be contained n>1 times
- Timeout, items have a limited lifetime in the stack. Items timed out will be removed from the stack

