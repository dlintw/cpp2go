This utility is to let you quick map concept of C/C++ to Go.

The official site is http://github.com/dlintw/cpp2go

I need help to improve this utility.
Contact me by e-mail ( dlin.tw gmail )

Features
========

* search C/C++ keyword and mapping to GO
* display useful hints for C/C++ users

Install
=======

Dependency before install: `Google Go language <http://golang.org>`_

Note: support Go 1 only (you may require download on http://weekly.golang.org).

commands::

  // method 1
  git clone https://github.com/dlintw/cpp2go
  cd cpp2go
  make
  cp cpp2go <your_path>

  // method 2
  go get github.com/dlintw/cpp2go

Usage
=====

::
  cpp2go [options] <keyword_in_c/c++>

  cpp2go -h # list option help

  Usage of ./cpp2go:
    -n=false: is listing hint numbers
    -t=false: is testing

Example
=======

::

  cpp2go sprintf explicit long # search multiple keywords
  cpp2go -n # list hints
  cpp2go 12 # list No.12 hint

TODO
====

* add more c common function call
* add more c++ common stl
* add more c++ concept sample

.. vi:set et sw=2 ts=2:
