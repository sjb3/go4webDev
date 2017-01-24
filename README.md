# go4webDev

adopted from Larry Price' tut from PACKT

middleware: negroni is from https://github.com/urfave/negroni

## also, a note from golang meetup re: middleware

#### golang middlewre | negroni:  

very similar to expressJS

not a full framework/a library familiar API

can be used with routing pkg

lots of 3rd party middleware

github.com/urfave/negroni

#### golang middleware | Interpose

anther lib

nesting of FIFO middleware

github.com/carbonation/interpose

#### golang middleware: alice

very simple lib

"TINY"

builds chain of middleware

supports any handler - must be handler fun

transforming syntax from doing it urself

github.com//justinas/alice

#### glang middleware: RYE

easy to configure

built in stated per middleware

supports 1.7 Golang

one of the box middleware

jet concept

extensible

set-up)

cars: allow cross origin calls easily

      allow specific headers, methods, and origins

      defaults for working w dev

origins: methods/ CRUD

JWT)

easy JWT validation

looks 4 authorization header w/ a bearer prefix

strips prefix, checks for jwt(returns 400 if not found

uses “jwt-go”

route-logging)

super simple route logging

uses logos for logging routes

addy, methods, uri, protocol(http 1.0, 2.0 etc)

using context)

context is automatically added to the req scope

req 1.7 context - built for future w/ go

unobstructive

example JWT middleware adds JWT to the context  - retrieval is easy
