# GO

Go is another compiled language created by Google and three engineers to improve software development and reduce complexity from another high-level languages like C++ or Java, tend to improve parallel programming and reduce hardware use.

# How to install

## Ubuntu

sudo snap install go --classic

>[!Note]
> The argument --classic is required to allow Go to install without a virtual environment and have access to the whole machine. This is secure since it uses the default and official Go repository.

# Summary

The following exercises folder develop the next assigments.

1.- Basic Structure (Hello World)
2.- Package Management
3.- Variables (Syntaxis)
4.- Conditions (Syntacis)
5.- Datatypes (Dinamic Arrays & Maps)
6.- Function Structure
7.- Pointers
8.- ASCII and Unicode
9.- Concurrence

## Basic Structure

Go is made to be efficient it only allows to compile if all dependencies and variables are used, if not it will not allow to compile and run, the import process can be done with external packages from Go, to accomplish this it is need to type the next commands.

```
go mod init {name-proyect}
```

```
go get {github-or-repository-address}
```