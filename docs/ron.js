"use strict";
function resolve (rule_name, rules, patterns) {
    if (rule_name in patterns) {
        return patterns[rule_name]
    }
    const rule = rules[rule_name];
    return patterns[rule_name] = rule.replace(/\$(\w+)/g, (match, name) => {
        const parser = resolve(name, rules, patterns);
        return parser;
    });
}

const RULES = {

    BASE64:     "[0-9A-Za-z_~]",
    UNICODE:    "\\\\u[0-9a-fA-F]\\{4}",
    INT:        "\\([(\\[{}\\])]\\)\\{0,1}$BASE64\\{0,10}",
    UUID:       "[`]\\{0,1}\\($BASE64\\/\\)\\{0,1}\\($INT\\)\\{0,1}\\([-+$%]\\)\\{0,1}\\($INT\\)\\{0,1}",

    INT_ATOM:   "=\\s*[+-]\\{0,1}[0-9]\\{1,17}",
    UUID_ATOM:  ">\\s*$UUID",
    STRING_ATOM:"'\\($UNICODE\\|\\\\[^\\n\\r]\\|[^'\"\\\\\\n\\r]\\)*'",
    FLOAT_ATOM: "\\^\\s*[+-]\\{0,1}[0-9]\\{0,19}[.][0-9]\\{1,19}\\([Ee][+-]\\{0,1}[0-9]\\{1,3}\\)\\{0,1}",
    OPTERM:     "[!?,;]",
    FRAMETERM:  "[.]",

    ATOM:       "=($INT_ATOM)|'($STRING_ATOM)'|\\^($FLOAT_ATOM)|>($UUID)",
    TYPE_ID:    "*\\s*$UUID",
    OBJECT_ID:  "#\\s*$UUID",
    EVENT_ID:   "@\\s*$UUID",
    REF_ID:     ":\\s*$UUID",
    OP:         "(?:\\s*\\*\\s*($UUID_ATOM))?(?:\\s*#\\s*($UUID_ATOM))?(?:\\s*@\\s*($UUID_ATOM))?(?:\\s*:\\s*($UUID_ATOM))?\\s*((?:\\s*$ATOM)*)\\s*($OPTERM)?",
    FRAME:      "($OP)+$FRAMETERM?",

};

const PATTERNS = {};


console.log('" vim syntax file');
console.log('" Language: Replicated Object Notation');
console.log('"');
console.log('if exists("b:current_syntax")');
console.log('    finish');
console.log('endif');
console.log('');

const COLORS = {
    INT_ATOM: "LightGreen",
    STRING_ATOM: "DarkGreen",
    UUID_ATOM: "DarkGreen",
    FLOAT_ATOM: "Green",
    OBJECT_ID: "Blue",
    TYPE_ID: "DarkMagenta",
    EVENT_ID: "DarkBlue",
    REF_ID: "Cyan",
    OPTERM: "Yellow",
    FRAMETERM: "Red",
};

Object.keys(COLORS).forEach((id) => {
    console.log( "syn match "+id+" /" + resolve(id, RULES, PATTERNS) + "/" );
    console.log( "hi "+id+" ctermfg="+COLORS[id]);
});

console.log('');
console.log('let b:current_syntax = "ron"');
