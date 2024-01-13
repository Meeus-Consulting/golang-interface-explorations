# Partial Function Application and Composition Root in Go

This repository explores the concepts of partial function application and the use of a composition root in Go, as an
alternative to traditional single-purpose interfaces and composition roots.

## Overview

In many Go applications, interfaces are used extensively to define contracts for various components, allowing for easy
testing and decoupling. However, this can sometimes lead to bloated interfaces, rigid structures, and an overemphasis on
object-oriented principles. This project demonstrates an alternative approach, using partial function application and a
simplified composition root.

## Key Concepts

1. **Partial Function Application**: This technique involves creating a function that is preloaded with some arguments
   and returns another function. In Go, this is achieved using closures. It allows for greater modularity and
   reusability of code.

2. **Composition Root**: The place in an application where the wiring of different components takes place. In this
   approach, the composition root is simplified, focusing on function-based dependencies rather than interface-based
   ones.

## Opinion

I'm not entirely sure this is idiomatic, but it does feel cleaner than going the interface route. It's also slightly more
terse and sidesteps the need for dinky little methods. It also enforces through the type-system that we're not going
to incur bloat.
