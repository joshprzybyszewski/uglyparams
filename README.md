# uglyparams
A golang linter to make lots of params ugly

## Problem to Solve

Sometimes functions/methods will grow over time and end up with "just one more input param."
Before you know it, there might be 12 or 15 input parameters. This is too many for multiple
reasons. First, it makes it difficult to know what the functions' relevant inputs are. Second,
it causes a lot of copying when making the function call. Third, it signals that your code is
smelling bad and should be re-thought.

However, there's no reason for devs to avoid adding more and more input parameters. One way
to increase the pain of seeing a 15 input parameter function/method is to make it take up a
ton of real-estate on screen and to put each input parameter on its own line.

## How to solve it

Eventually, this may end up in golangci-lint. For now, you should be able to use this with
the following:

```bash
go install github.com/joshprzybyszewski/uglyparams

uglyparams -fix ./myProject/...
```