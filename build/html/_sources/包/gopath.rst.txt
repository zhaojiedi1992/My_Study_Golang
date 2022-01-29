gopath
=================================

GOPATH 是 Go语言中使用的一个环境变量，它使用绝对路径提供项目的工作目录。

.. code-block:: go 

    ➜  ~ go env
    GO111MODULE=""
    GOARCH="amd64"
    GOBIN=""
    GOCACHE="/Users/dxm/Library/Caches/go-build"
    GOENV="/Users/dxm/Library/Application Support/go/env"
    GOEXE=""
    GOFLAGS=""
    GOHOSTARCH="amd64"
    GOHOSTOS="darwin"
    GONOPROXY=""
    GONOSUMDB=""
    GOOS="darwin"
    GOPATH="/Users/dxm/go"
    GOPRIVATE=""
    GOPROXY="https://mirrors.aliyun.com/goproxy/"
    GOROOT="/usr/local/go"
    GOSUMDB="sum.golang.org"
    GOTMPDIR=""
    GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
    GCCGO="gccgo"
    AR="ar"
    CC="clang"
    CXX="clang++"
    CGO_ENABLED="1"
    GOMOD=""
    CGO_CFLAGS="-g -O2"
    CGO_CPPFLAGS=""
    CGO_CXXFLAGS="-g -O2"
    CGO_FFLAGS="-g -O2"
    CGO_LDFLAGS="-g -O2"
    PKG_CONFIG="pkg-config"
    GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/3p/jc3w8_dn5dd0r5_5ccp3nrqm0000gn/T/go-build619420410=/tmp/go-build -gno-record-gcc-switches -fno-common"

使用gopath工程结构
--------------------------------------------

.. code-block:: go 

    export GOPATH=`pwd`
    mkdir -p src/hello

    填写代码

    go install hello