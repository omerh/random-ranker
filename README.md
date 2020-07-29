# random-ranker

Get random ranking when choosing the rank of the first position.
Based on [cobra](https://github.com/spf13/cobra)

Parameters are:

1. `--p1` the rank of the first position
2. `--count` the amount of positions to rank

## Build

In order to build this

```bash
git clone https://github.com/omerh/random-ranker.git
cd ./random-ranker
go build -ldflags "-s -w"
```

## Run

```bash
./random-ranker get random --p1 1000 --count 5
```

It will print out the following:

```log
2020/07/29 16:07:11 Setting ranking for 5 positions with P1 starting rank of 1000
P_1 is:  1000
P_2 is:  965
P_3 is:  919
P_4 is:  887
P_5 is:  854
```

### Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).

### Copyright

Copyright (c) 2019 Omer Haim, [@omerhaim](http://twitter.com/omerhaim).
See [LICENSE](LICENSE) for further details.
