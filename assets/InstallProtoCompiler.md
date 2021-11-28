# Mac OS Install

1(a) Install wget via brew if you don't already have it
```sh
brew install wget
```


1(b) Brew install & upgrade
```sh
brew install protobuf
brew upgrade protobuf

```
2. Specific to the project and installs the GO protobuf runtime. If we wanted to compile in more langues like python, javascript, java...etc we would need to install their runtimes as well
```sh
go get google.golang.org/protobuf/...@v1.25.0
```

3. `zsh` is my default terminal command line tool. I want to make sure that `protoc-gen-go` can be accessed. You should be able to run `grep -R error $HOME/go/bin/protoc-gen-go` and get some output like `Binary file /Users/bmoore/go/bin/protoc-gen-go matches`. If it fails it should look something like `grep: /Users/bmoore/go/bin/protoc-gen-go: No such file or directory`

```sh
echo 'export PATH="$PATH:$HOME/go/bin"' >> ~/.zshenv
```

4(a). Test Install was successful
```sh
protoc --version
```

4(b) Test Makefile can compile protos
```sh
make compile-protos
```



Refer to the [Official Protoc Installation Link](https://grpc.io/docs/protoc-installation/) if these steps didn't help


