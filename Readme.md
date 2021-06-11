# Go restful 개발환경 구축
> vscode에서 golang으로 web, restful 개발환경 구축을 시작하자

[![NPM Version][npm-image]][npm-url]
[![Build Status][travis-image]][travis-url]
[![Downloads Stats][npm-downloads]][npm-url]

![](../header.png)
## 사전준비

vscode 설치 : https://code.visualstudio.com/
git 설치 : https://git-scm.com/downloads

golang 설치 (path 설정):

```sh
윈도우키 + R > systdm.cpl ,3
```
![image](https://user-images.githubusercontent.com/16375921/121686096-b52a9180-cafb-11eb-9cab-2ee8eb1e72f1.png)

vscode framework 설치:

```sh
go get -u github.com/labstack/echo/...
```

## Troubleshooting

"github.com/... " 패키지 import 에러 발생
```sh
go env -w GO111MODULE=auto
```

## vscode 플러그인
```sh
GitLens : git 관리
Git History : commit/push 이력 관리
```

<!-- Markdown link & img dfn's -->
[npm-image]: https://img.shields.io/npm/v/datadog-metrics.svg?style=flat-square
[npm-url]: https://npmjs.org/package/datadog-metrics
[npm-downloads]: https://img.shields.io/npm/dm/datadog-metrics.svg?style=flat-square
[travis-image]: https://img.shields.io/travis/dbader/node-datadog-metrics/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/dbader/node-datadog-metrics
[wiki]: https://github.com/yourname/yourproject/wiki
