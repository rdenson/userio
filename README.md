# USERIO - Colorized Command-Line Input/Output
General Purpose I/O For Go CLI projects (in color!).

## Summary
This submodule is meant to be a wrapper for formatting screen output and ingesting
user input. The goal here is easy portability into a Go CLI program and for providing
some semblance of readability against a dark background.

### Design Philosophy
Simplistic colors, forming brief predictable output. If you need something trendy
or eye-catching look elsewhere. This package is for applying simple formatting to
CLI programs.

### Technical
basic components:
* display - output something
* prompt - take in user input
* choice - menu style selection (_in testing_)

colors and emojis employed here are expressed in ANSI escape sequences
