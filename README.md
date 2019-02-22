[![JadeTrans Usage][11]][9]

A very geeky command line translation tool.


[![Build Status][1]][2] [![MIT licensed][3]][4] [![Coverage Status][5]][6] [![Go Report Card][7]][8]

![JadeTrans Usage][10]

[1]: https://travis-ci.org/archervanderwaal/jadetrans.svg?branch=master
[2]: https://travis-ci.org/archervanderwaal/jadetrans
[3]: https://img.shields.io/dub/l/vibe-d.svg
[4]: https://github.com/archervanderwaal/jadetrans/blob/master/LICENSE
[5]: https://coveralls.io/repos/github/archervanderwaal/jadetrans/badge.svg
[6]: https://coveralls.io/github/archervanderwaal/jadetrans
[7]: https://goreportcard.com/badge/github.com/archervanderwaal/jadetrans
[8]: https://goreportcard.com/report/github.com/archervanderwaal/jadetrans
[9]: https://github.com/archervanderwaal/jadetrans
[10]: https://github.com/archervanderwaal/jadetrans/blob/master/snapshots/jadetrans-usage.gif
[11]: https://github.com/archervanderwaal/jadetrans/blob/master/snapshots/jadetrans-logo.png


## Features

[x] Chinese <=> English
[] Speech
[] Others language support
[] Google translation engine

## Installation

### Go

```bash

go get -u github.com/archervanderwaal/jadetrans

jadetrans -h
```

## Usage

### Configuration

```bash
cd $HOME/.jadetrans
vim jadetrans.yaml
# add appKey and appSecret to jadetrans.yaml
# exp:
# youdao:
#   appKey: 
#   appSecret: 
```

If you don't have appKey and appSecret translated by youdao, go [here](https://blog.csdn.net/qq_38288606/article/details/80522233) and learn how to get appkey and appsecret translated by youdao.


### Query

```bash
jadetrans <words>
```

### Select engine

```bash
jadetrans <words> -e=youdao
```

### Speech

```bash
jadetrans <words> -voice=1
#or
jadetrans <words> -voice=0
```