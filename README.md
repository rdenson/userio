[![Build Status](https://travis-ci.org/rdenson/userio.svg?branch=master)](https://travis-ci.org/rdenson/userio)

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
There are three basic components here:
* display - output something
* prompt - take in user input
* interpretation - still under conceptualization
 * a furthering of "prompt" to encapsulate what I'm calling "a user request"

Colors and emojis employed here are expressed in ANSI escape sequences.
Color Palette below 👇

| ANSI Sequence | Color | Usage |
| --- | --- | --- |
| `\u001b[32m` | <span style="color:rgb(37,188,36)"> green </span> | default or standard |
| `\u001b[33m` | <span style="color:rgb(173,173,39)"> yellow </span> | information |
| `\u001b[32m` | <span style="color:rgb(51,187,200)"> cyan </span> | list element |
| `\u001b[35m` | <span style="color:rgb(211,56,211)"> magenta </span> | instruction |
