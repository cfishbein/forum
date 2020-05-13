_Note: Putting development on hold until Go 2 is released. More specifically, when generics and cleaner error handling
is implemented. I find the verbosity of constantly error checking every tiny function tiring, unkempt, and just an
overall unejoyable experience, despite its correctness. I also find that writing reusable code without generics is
difficult at best, not possible at worst. Defining functions is one way to do it, but that just doesn't make for clean
code in most cases._

# Forum

A basic web forum to explore the Go programming language.

# Installation

Run the appropriate database creation script for your operating system under the `config` directory.

# TODO

[ ] Cleaner Form validation for incoming requests

[ ] Render Go HTML templates on server side or maybe keep the whole thing as a REST API and use a front end framework

[ ] Implement interfaces for the db layer to plugin something other than SQLite3
