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

## $${\color{royalblue}Basic \space Structure}$$

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

## $${\color{royalblue}Package \space Management}$$

<p align="justify">  
Working with packages in Go requires naming all the files with the same package. To import other libraries, it is necessary to specify the GitHub repository path. Once this is done, a go.mod file is created containing the Go version and the desired functions.
</p>

```
go get {github-address-directory}
```

## $${\color{royalblue}Variables \space (Syntax)}$$

<p align="justify">  
Go allows multiple formats for variable declaration it can use standard C or inferred like Python accepted syntax:
</p>

```
var number int = 0
```

```
var number = 0
```

```
number := 0
```

[>!Note]
> The operator ":=" is special for the Go language and is used to declare the implicit type, there is actually a curious way to use it for example encapsulating functions from libraries inside a dynamic data type and then use it as a regular function similar to a macro declared as a header in other languages.

## $${\color{royalblue}\space Conditions \space (Syntax)}$$

<p align="justify">  
As the variable rules, the syntax for conditions can be declared in different ways including the short assignment operator ":=" to create the data type inside the condition as for example
</p>

```
if random := rand.IntN(10); random > 5
```

```
for l := range 1
```

# $${\color{royalblue}\space Datatypes \space (Dynamic \space Arrays \space and \space Maps)}$$

<p align="justify">  
The use of dynamic arrays or maps requires the make function standard arrays have noticeable interesting ways to be declared Go can infer the size of an array with [...] declaration Since variables in Go are not objects like Python the assignment of new memory allocations is not part of a method or function itself, it is needed to make a new assignment.
</p>

```
dinamic_vector_python = append(dinamic_vector_python, i) 
```

# $${\color{royalblue}\space Function \space Structure}$$

<p align="justify">  
Functions can return multiple data types at the same time, making this a relevant feature due to the limitations from other languages and the need of using arrays or tuples to get the exact behaviour, the arguments do not need to be specified individually if all match the same data type.
</p>

```
func overload(n ... int)int
```

```
func overload(a, b, c int)int int int
```

# $${\color{royalblue}\space Pointers}$$

<p align="justify">  
Memory management can be used to solve memory addressing.
</p>

```
ptx := &array[3]
```

```
*ptx = 0
```

# $${\color{royalblue}\space ASCII \space and \space Unicode}$$

<p align="justify">  
ASCII code is extended with the use of runes accepting the Unicode/UTF-8 standard for string manipulation.
</p>

# $${\color{royalblue}\space Concurrency}$$

<p align="justify">  
Go routines are a lightweight thread employed in the same memory address of the current thread, it is a feature intended in Go to follow its philosophy "Do not communicate by sharing memory; instead, share memory by communicating." This resembles the tool called channels to extract information and reduce the use of more complex tools and methodologies to avoid deadlocks and race conditions, channels hold the execution and wait until the Go routine has finished.
</p>

<p align="justify">  
Channels can also be buffered stopping only when the buffer is full, Go internally has a certain amount of validation that will interrupt execution if it detects a deadlock instead of just waiting for a never end.
</p>

More information on <a href = "https://go.dev/blog/codelab-share"> Go </a> documentation.

```
func counter(ch chan bool){

	for i := 0; i < 10; i++{

		fmt.Println("We are couting with this one! 🗣️🔥 ", i)
	}

	ch <- true // If not returned this value a deadlock happens, since it does not return nothing.
}
```

```
go counter(ch)

```

[>!Tip]
> Timeouts and custom error messages with panic are used to manage workflow between Go routines.