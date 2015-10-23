package com.gschat.rpctest;

using gslang.Package;

@Package(Lang:"golang",Name:"com.gschat.rpctest",Redirect:"github.com/gschat/rpctest")


contract Server {
    void push(byte[] content);
}
