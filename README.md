# GO-Study

go programming language self study

Go, or Golang as it is often called, was developed by Google employees—chiefly longtime Unix guru and Google distinguished engineer Rob Pike—but it’s not strictly speaking a “Google project.” Rather, Go is developed as a community-led open source project, spearheaded by leadership that has strong opinions about how Go should be used and the direction the language should take.

Go is meant to be simple to learn, straightforward to work with, and easy to read by other developers. Go does not have a large feature set, especially when compared to languages like C++. Go is reminiscent of C in its syntax, making it relatively easy for longtime C developers to learn. That said, many features of Go, especially its concurrency and functional programming features, harken back to languages such as Erlang.
As a C-like language for building and maintaining cross-platform enterprise applications of all sorts, Go has much in common with Java. And as a means of enabling rapid development of code that might run anywhere, you could draw a parallel between Go and Python, though the differences are far greater than the similarities.

Go has been compared to scripting languages like Python in its ability to satisfy many common programming needs. Some of this functionality is built into the language itself, such as “goroutines” for concurrency and threadlike behavior, while additional capabilities are available in Go standard library packages, like Go’s http package. Like Python, Go provides automatic memory management capabilities including garbage collection.
Unlike scripting languages such as Python, Go code compiles to a fast-running native binary. And unlike C or C++, Go compiles extremely fast—fast enough to make working with Go feel more like working with a scripting language than a compiled language. Further, the Go build system is less complex than those of other compiled languages. It takes few steps and little bookkeeping to build and run a Go project.


                                                                  Hardware Limitations

We have observed that in a decade, the hardware and processing configuration is changing at a very slow rate. In 2004, P4 was having the clock speed of 3.0 GHz and now in 2018, Macbook pro has the clock speed of Approx (2.3Ghz v 2.66Ghz). To speed up, the functionality we use more processors, but using more processor the cost also increases. And due to this we use limited processors and using limited processor we have a heavy programming language whose threading takes more memory and slows down the performance of our system. Hence, to overcome such problem Golang has been designed in such a way that instead of using threading it uses Goroutine, which is similar to threading but consumes very less memory. 
Like threading consumes 1MB whereas Goroutine consumes 2KB of memory, hence at the same time, we can have millions of goroutine triggered. 
So the above-discussed point makes golang a strong language that handles concurrency like C++ and Java. 


Advantages and Disadvantages of Go Language

Advantages:  

Flexible- It is concise, simple and easy to read.
Concurrency- It allows multiple process running simultaneously and effectively.
Quick Outcome- Its compilation time is very fast.
Library- It provides a rich standard library.
Garbage collection- It is a key feature of go. Go excels in giving a lot of control over memory allocation and has dramatically reduced latency in the most recent versions of the garbage collector.
It validates for the interface and type embedding.
Concurrency: Go provides excellent support for concurrency, making it easy to write code that can run multiple tasks simultaneously. This is achieved through Goroutines and Channels, which allow you to write code that can run multiple operations at the same time.
Performance: Go is designed to be fast and efficient, with a focus on performance and low memory usage. This makes it well-suited for building high-performance network services, as well as for solving complex computational problems.
Simplicity: Go has a straightforward syntax and a simple type system, making it easy to learn and use, even for people with no prior programming experience.
Garbage Collection: Go has built-in garbage collection, which automatically manages memory for you. This eliminates the need for manual memory management, reducing the likelihood of memory leaks and other bugs that can arise from manual memory management.
Statically Typed: Go is a statically typed language, which means that types are determined at compile time. This provides stronger type safety and makes 
it easier to catch type-related bugs before they occur.



Disadvantages:  

It has no support for generics, even if there are many discussions about it.
The packages distributed with this programming language is quite useful but Go is not so object-oriented in the conventional sense.
There is absence of some libraries especially a UI tool kit.
Limited Object-Oriented Features: Go does not have full-fledged object-oriented features like inheritance and polymorphism. This can make it more difficult to write complex programs, especially for developers who are used to traditional object-oriented languages.
No Generics: Go does not have built-in support for generics, which makes it difficult to write reusable code.
Immature Standard Library: Go’s standard library is relatively new and still maturing, which can make it difficult to find the tools you need for a particular task.


