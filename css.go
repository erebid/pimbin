package pimbin

const defaultCSS = ` /* Pimbin's default CSS */

body { font-family: monospace; background-color: #ffffff }

a, a:visited {color: #000000;}

.filename, body > h1 {font-size: 1.25em;}
.filename {display: inline;}

img {display: block; max-width: 30vw;}

#file-index {list-style-type: none; padding-left: 2ch;}
/* Background */ .chroma { background-color: #ffffff }
/* Error */ .chroma .err { color: #a61717; background-color: #e3d2d2 }
/* LineTableTD */ .chroma .lntd { vertical-align: top; padding: 0; margin: 0; border: 0; }
/* LineTable */ .chroma .lntable { border-spacing: 0; padding: 0; margin: 0; border: 0; width: auto; overflow: auto; display: block; }
/* LineHighlight */ .chroma .hl { display: block; width: 100%;background-color: #e5e5e5 }
/* LineNumbersTable */ .chroma .lnt { margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
/* LineNumbers */ .chroma .ln { margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
/* Keyword */ .chroma .k { color: #000000; font-weight: bold }
/* KeywordConstant */ .chroma .kc { color: #000000; font-weight: bold }
/* KeywordDeclaration */ .chroma .kd { color: #000000; font-weight: bold }
/* KeywordNamespace */ .chroma .kn { color: #000000; font-weight: bold }
/* KeywordPseudo */ .chroma .kp { color: #000000; font-weight: bold }
/* KeywordReserved */ .chroma .kr { color: #000000; font-weight: bold }
/* KeywordType */ .chroma .kt { color: #445588; font-weight: bold }
/* NameAttribute */ .chroma .na { color: #008080 }
/* NameBuiltin */ .chroma .nb { color: #0086b3 }
/* NameBuiltinPseudo */ .chroma .bp { color: #999999 }
/* NameClass */ .chroma .nc { color: #445588; font-weight: bold }
/* NameConstant */ .chroma .no { color: #008080 }
/* NameDecorator */ .chroma .nd { color: #3c5d5d; font-weight: bold }
/* NameEntity */ .chroma .ni { color: #800080 }
/* NameException */ .chroma .ne { color: #990000; font-weight: bold }
/* NameFunction */ .chroma .nf { color: #990000; font-weight: bold }
/* NameLabel */ .chroma .nl { color: #990000; font-weight: bold }
/* NameNamespace */ .chroma .nn { color: #555555 }
/* NameTag */ .chroma .nt { color: #000080 }
/* NameVariable */ .chroma .nv { color: #008080 }
/* NameVariableClass */ .chroma .vc { color: #008080 }
/* NameVariableGlobal */ .chroma .vg { color: #008080 }
/* NameVariableInstance */ .chroma .vi { color: #008080 }
/* LiteralString */ .chroma .s { color: #dd1144 }
/* LiteralStringAffix */ .chroma .sa { color: #dd1144 }
/* LiteralStringBacktick */ .chroma .sb { color: #dd1144 }
/* LiteralStringChar */ .chroma .sc { color: #dd1144 }
/* LiteralStringDelimiter */ .chroma .dl { color: #dd1144 }
/* LiteralStringDoc */ .chroma .sd { color: #dd1144 }
/* LiteralStringDouble */ .chroma .s2 { color: #dd1144 }
/* LiteralStringEscape */ .chroma .se { color: #dd1144 }
/* LiteralStringHeredoc */ .chroma .sh { color: #dd1144 }
/* LiteralStringInterpol */ .chroma .si { color: #dd1144 }
/* LiteralStringOther */ .chroma .sx { color: #dd1144 }
/* LiteralStringRegex */ .chroma .sr { color: #009926 }
/* LiteralStringSingle */ .chroma .s1 { color: #dd1144 }
/* LiteralStringSymbol */ .chroma .ss { color: #990073 }
/* LiteralNumber */ .chroma .m { color: #009999 }
/* LiteralNumberBin */ .chroma .mb { color: #009999 }
/* LiteralNumberFloat */ .chroma .mf { color: #009999 }
/* LiteralNumberHex */ .chroma .mh { color: #009999 }
/* LiteralNumberInteger */ .chroma .mi { color: #009999 }
/* LiteralNumberIntegerLong */ .chroma .il { color: #009999 }
/* LiteralNumberOct */ .chroma .mo { color: #009999 }
/* Operator */ .chroma .o { color: #000000; font-weight: bold }
/* OperatorWord */ .chroma .ow { color: #000000; font-weight: bold }
/* Comment */ .chroma .c { color: #999988; font-style: italic }
/* CommentHashbang */ .chroma .ch { color: #999988; font-style: italic }
/* CommentMultiline */ .chroma .cm { color: #999988; font-style: italic }
/* CommentSingle */ .chroma .c1 { color: #999988; font-style: italic }
/* CommentSpecial */ .chroma .cs { color: #999999; font-weight: bold; font-style: italic }
/* CommentPreproc */ .chroma .cp { color: #999999; font-weight: bold; font-style: italic }
/* CommentPreprocFile */ .chroma .cpf { color: #999999; font-weight: bold; font-style: italic }
/* GenericDeleted */ .chroma .gd { color: #000000; background-color: #ffdddd }
/* GenericEmph */ .chroma .ge { color: #000000; font-style: italic }
/* GenericError */ .chroma .gr { color: #aa0000 }
/* GenericHeading */ .chroma .gh { color: #999999 }
/* GenericInserted */ .chroma .gi { color: #000000; background-color: #ddffdd }
/* GenericOutput */ .chroma .go { color: #888888 }
/* GenericPrompt */ .chroma .gp { color: #555555 }
/* GenericStrong */ .chroma .gs { font-weight: bold }
/* GenericSubheading */ .chroma .gu { color: #aaaaaa }
/* GenericTraceback */ .chroma .gt { color: #aa0000 }
/* GenericUnderline */ .chroma .gl { text-decoration: underline }
/* TextWhitespace */ .chroma .w { color: #bbbbbb }
`
