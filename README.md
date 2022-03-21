# USERIO - Colorized Command-Line Input/Output
General Purpose I/O For Go CLI projects (in color!).

## Summary
This submodule is meant to be a wrapper for printing messages and ingesting user
input. The goal here is portability into any basic cli program and for providing
some sort of standardization. Colorized for your pattern recognition pleasure,
it's a reliable boiler plate Go import.  
_"designed for black backgrounds"_

### Design Philosophy
Simplistic colors, forming brief predictable output. If you need something trendy
or eye-catching look elsewhere. This package is for applying simple colors to CLI
programs.

### Technical
Basic components here:
* display - output something
* prompt - take in user input

_colors and emojis employed here are expressed in ANSI escape sequences_
