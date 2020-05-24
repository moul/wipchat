# wipchat

:smile: wipchat

![CI](https://github.com/moul/wipchat/workflows/CI/badge.svg)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/wipchat)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/wipchat/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/wipchat.svg)](https://github.com/moul/wipchat/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/wipchat)](https://goreportcard.com/report/moul.io/wipchat)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/wipchat/badge)](https://www.codefactor.io/repository/github/moul/wipchat)
[![codecov](https://codecov.io/gh/moul/wipchat/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/wipchat)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/wipchat.svg)](https://microbadger.com/images/moul/wipchat)
[![GolangCI](https://golangci.com/badges/github.com/moul/wipchat.svg)](https://golangci.com/r/github.com/moul/wipchat)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)


## Usage

```console
$ export WIPCHAT_KEY=XXX
```

```console
$ wipchat todo 👋 hello world
{
  "CreateTodo": {
    "ID": "150865",
    "CreatedAt": "2020-05-24T23:43:14Z",
    "UpdatedAt": "2020-05-24T23:43:14Z",
    "Body": "👋 hello world",
    "User": {
      "ID": "1780",
      "URL": "https://wip.chat/@moul"
    }
  }
}
```

```console
$ wipchat done "👋 hello world #oss"wipchat
{
  "CreateTodo": {
    "ID": "150866",
    "CreatedAt": "2020-05-24T23:44:15Z",
    "CompletedAt": "2020-05-24T23:44:15Z",
    "UpdatedAt": "2020-05-24T23:44:15Z",
    "Body": "👋 hello world #oss",
    "Product": {
      "ID": "3493",
      "Hashtag": "oss",
      "URL": "https://wip.chat/products/oss"
    },
    "User": {
      "ID": "1780",
      "URL": "https://wip.chat/@moul"
    }
  }
}
```

```console
$ wipchat me  | jq . | head -11
{
  "Viewer": {
    "ID": "1780",
    "URL": "https://wip.chat/@moul",
    "Username": "moul",
    "Firstname": "Manfred",
    "Lastname": "Touron 🇫🇷",
    "AvatarURL": "https://wip.imgix.net/cache/user/1780/avatar/acb8193abf5d107a8d644da41b071494.png?ixlib=rb-3.2.1&w=64&h=64&fit=crop&s=1b2e60b141240f8b3f975a2dc42f5cb5",
    "CompletedTodosCount": 128,
    "BestStreak": 16,
    "Streaking": true,
```

```console
$ wipchat me | jq '.Viewer.Todos[].Body'
"🟨 join WIP #life"
"🟨 add #berty on WIP"
"🟨 add #oss on WIP"
"♻️ find or make an integration to have my todos on WIP without leaving trello #life"
"🐛 fix an AMP bug in a hugo-template, that I missed for multiple weeks on my personal website #life"
```

```console
$ wipchat me | jq '.Viewer.Products[].Name'
"Missions"
"protoc-gen-gotemplate"
"Alfred TOTP"
"Scaleway"
"gRPCb.in"
"FranceP2P / Paris P2P"
"Ultreme"
"Wulo"
"WIP Tools"
"Berty"
"Manfred's Life"
"Open-source stuff"
"Millipede"
"multiarch"
"Pathwar"
"depviz"
"sshportal"
"assh"
```

## Install

### Using go

```console
$ go get -u moul.io/wipchat
```

### Releases

See https://github.com/moul/wipchat/releases

## License

© 2020 [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
