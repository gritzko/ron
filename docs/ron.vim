" vim syntax file
" Language: Replicated Object Notation
"
if exists("b:current_syntax")
    finish
endif

syn match INT_ATOM /=\s*[+-]\{0,1}[0-9]\{1,17}/
hi INT_ATOM ctermfg=LightGreen
syn match STRING_ATOM /'\(\\u[0-9a-fA-F]\{4}\|\\[^\n\r]\|[^'"\\\n\r]\)*'/
hi STRING_ATOM ctermfg=DarkGreen
syn match UUID_ATOM />\s*[`]\{0,1}\(\([(\[{}\])]\)\{0,1}[0-9A-Za-z_~]\{0,10}\)\{0,1}\([-+$%]\)\{0,1}\(\([(\[{}\])]\)\{0,1}[0-9A-Za-z_~]\{0,10}\)\{0,1}/
hi UUID_ATOM ctermfg=DarkGreen
syn match FLOAT_ATOM /\^\s*[+-]\{0,1}[0-9]\{0,19}[.][0-9]\{1,19}\([Ee][+-]\{0,1}[0-9]\{1,3}\)\{0,1}/
hi FLOAT_ATOM ctermfg=Green
syn match OBJECT_ID /#\s*[`]\{0,1}\(\([(\[{}\])]\)\{0,1}[0-9A-Za-z_~]\{0,10}\)\{0,1}\([-+$%]\)\{0,1}\(\([(\[{}\])]\)\{0,1}[0-9A-Za-z_~]\{0,10}\)\{0,1}/
hi OBJECT_ID ctermfg=Blue
syn match TYPE_ID /*\s*[`]\{0,1}\(\([(\[{}\])]\)\{0,1}[0-9A-Za-z_~]\{0,10}\)\{0,1}\([-+$%]\)\{0,1}\(\([(\[{}\])]\)\{0,1}[0-9A-Za-z_~]\{0,10}\)\{0,1}/
hi TYPE_ID ctermfg=DarkMagenta
syn match EVENT_ID /@\s*[`]\{0,1}\(\([(\[{}\])]\)\{0,1}[0-9A-Za-z_~]\{0,10}\)\{0,1}\([-+$%]\)\{0,1}\(\([(\[{}\])]\)\{0,1}[0-9A-Za-z_~]\{0,10}\)\{0,1}/
hi EVENT_ID ctermfg=DarkBlue
syn match REF_ID /:\s*[`]\{0,1}\(\([(\[{}\])]\)\{0,1}[0-9A-Za-z_~]\{0,10}\)\{0,1}\([-+$%]\)\{0,1}\(\([(\[{}\])]\)\{0,1}[0-9A-Za-z_~]\{0,10}\)\{0,1}/
hi REF_ID ctermfg=Cyan
syn match OPTERM /[!?,;]/
hi OPTERM ctermfg=Yellow
syn match FRAMETERM /[.]/
hi FRAMETERM ctermfg=Red

let b:current_syntax = "ron"
