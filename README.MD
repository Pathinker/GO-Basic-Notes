# GO

<p align="justify">   
Go is another compiled language created by Google and three engineers to improve software development and reduce complexity from another high-level languages like C++ or Java, tend to improve parallel programming and reduce hardware use.
</p>

# How to install

## Ubuntu

```
sudo snap install go --classic
```

>[!Note]
> The argument --classic is required to allow Go to install without a virtual environment and have access to the whole machine. This is secure since it uses the default and official Go repository.

# Summary

The following exercises folder develop the next assigments.

- **1.-** $${\color{royalblue}\space Basic \space Structure \space (Hello World)}$$

- **2.-** $${\color{royalblue}\space Package \space Management}$$

- **3.-** $${\color{royalblue}\space Variables \space (Syntax)}$$

- **4.-** $${\color{royalblue}\space Conditions \space (Syntax)}$$

- **5.-** $${\color{royalblue}\space Datatypes \space (Dynamic \space Arrays \space and \space Maps)}$$

- **6.-** $${\color{royalblue}\space Function \space Structure}$$

- **7.-** $${\color{royalblue}\space Pointers}$$

- **8.-** $${\color{royalblue}\space ASCII \space and \space Unicode}$$

- **9.-** $${\color{royalblue}\space Concurrency}$$

## $${\color{royalblue}Basic \space Structure \space}$$

<p align="justify">   
Go is made to be efficient it only allows to compile if all dependencies and variables are used, if not it will not allow to compile and run, the import process can be done with external packages from Go, to accomplish this it is need to type the next commands.
</p>

```
go mod init {name-proyect}
```

```
go get {github-or-repository-address}
```

<p align="justify">   
In order to accomplish a hello world on Go it is necessary to declare a package this namespace will be used later to import local libraries and files then import the modules to use fmt is the standard IO functions The execution of a Go file can be direct on terminal by run or generate an executable such as GCC commands on C which can be later used to create an instance of the program.
</p>

```
go run {go-file}
```

```
go build -o {go-output-executable.o go-file.go}
```