var fs = require('fs');
var crypto = require('crypto');

const MACRO = /^(\s*\/\/\/)\s*(\w+)\s*(.*?)(?:\s*\[([0-9a-fA-F]+)\])?\s*$/; //(?:\[([0-9A-Fa-f])\])?$/;

const BLOCKS = {
  id: { body: 'lines', md5: '123abc' },
};

const TWEAKS = {
  id: { edits: [{ regex: /abc/, replace: '$1' }], md5: '' },
};

const USES = {
  blocks: [], //{ file: file, id: macro.arg, md5: macro.md5 }
  tweaks: [],
};

const DIRTY = {};

const files = process.argv.slice();
files.shift();
files.shift();

for (let file of files) {
  scan(file);
}

for (let use of USES.blocks) {
  const block = BLOCKS[use.id];
  if (!block) {
    console.error('block not found: ' + use.id);
    process.exit(4);
  }
  if (block.md5 != use.md5) {
    DIRTY[use.file] = use.id;
  }
}

for (let use of USES.tweaks) {
  const tweak = TWEAKS[use.id];
  if (!tweak) {
    console.error('tweak not found: ' + use.id);
    process.exit(5);
  }
  if (tweak.md5 != use.md5) {
    DIRTY[use.file] = use.id;
  }
}

const DIRTIES = Object.keys(DIRTY);

//console.log("BLOCKS: "+JSON.stringify(BLOCKS,4));
//console.log("TWEAKS: "+JSON.stringify(TWEAKS));
//console.log("USES: "+JSON.stringify(USES));
//console.log("DIRTY: "+JSON.stringify(DIRTY));

for (let dirty of DIRTIES) {
  console.log('editing ' + dirty);
  edit_file(dirty);
}

process.exit(0);

// ------------

function parse_line(line) {
  MACRO.lastIndex = 0;
  const m = MACRO.exec(line);
  if (!m) {
    return { cmd: 'line', arg: line, md5: '', indent: '' };
  } else {
    return { cmd: m[2], arg: m[3], md5: m[4] || '', indent: m[1] };
  }
}

function md5(data) {
  return crypto
    .createHash('md5')
    .update(data)
    .digest('hex')
    .substr(0, 8);
}

function scan(file) {
  const body = fs.readFileSync(file, { encoding: 'utf-8' });
  const lines = body.split(/\n/).reverse();
  scan_file(file, lines);
}

function scan_file(file, lines) {
  while (lines.length) {
    let line = lines.pop();
    const macro = parse_line(line);
    switch (macro.cmd) {
      case 'copy':
        scan_block([macro.arg], lines);
        break;
      case 'tweak':
        scan_tweak(macro.arg, lines);
        break;
      case 'paste':
        USES.blocks.push({ file: file, id: macro.arg, md5: macro.md5 });
        break;
      case 'use':
        USES.tweaks.push({ file: file, id: macro.arg, md5: macro.md5 });
        break;
      default:
    }
  }
}

function scan_block(stack, lines) {
  const id = stack[stack.length - 1];
  const block = (BLOCKS[id] = { body: '', md5: '' });
  while (lines.length) {
    let line = lines.pop();
    const macro = parse_line(line);
    switch (macro.cmd) {
      case 'copy':
        const nestedId = macro.arg;
        stack.push(nestedId);
        scan_block(stack, lines);
        stack.pop();
        break;
      case 'end':
        block.md5 = md5(block.body);
        return; // return!
      case 'line':
        for (let id of stack) {
          BLOCKS[id].body += line + '\n';
        }
        break;
    }
  }
}

function scan_tweak(id, lines) {
  const tweak = (TWEAKS[id] = { edits: [], md5: '' });
  let body = id;
  while (lines.length) {
    let line = lines.pop();
    const macro = parse_line(line);
    switch (macro.cmd) {
      case 're':
        body += macro.arg;
        var m = /^\/(.*)\/(.*)\/$/.exec(macro.arg);
        if (m == null) {
          console.error('cant parse: ' + macro.arg);
          process.exit(2);
        }
        const deesc = m[2].replace(/\\n/g, '\n');
        tweak.edits.push({ regex: new RegExp(m[1], 'mg'), replace: deesc });
        break;
      case 'fn':
        body += macro.arg;
        var m = /^\/(.*)\/\s*\((\w+(?:,\w+)*)\)\s*=>\s*{(.*)}\s*$/.exec(macro.arg);
        if (m == null) {
          console.error('cant parse: ' + macro.arg);
          process.exit(5);
        }
        const re = m[1],
          args = m[2],
          fn = m[3];
        tweak.edits.push({ regex: new RegExp(re, 'mg'), replace: new Function(args, fn) });
        break;
      case 'end':
        tweak.md5 = md5(body);
        return;
    }
  }
}

function edit_file(file) {
  const body = fs.readFileSync(file, { encoding: 'utf-8' });
  const lines = body.split(/\n/).reverse();
  const edited = [];

  while (lines.length) {
    let line = lines.pop();
    const macro = parse_line(line);
    switch (macro.cmd) {
      case 'paste':
        const id = macro.arg;
        let block = BLOCKS[id];
        edited.push(macro.indent + ' paste ' + id + ' [' + block.md5 + ']');
        edit_block(block, lines, edited);
        break;
      default:
        edited.push(line);
    }
  }
  fs.writeFileSync(file, edited.join('\n'));
}

function edit_block(block, lines, edited) {
  let body = block.body;
  while (lines.length) {
    let line = lines.pop();
    const macro = parse_line(line);
    switch (macro.cmd) {
      case 'use':
        const tweak = TWEAKS[macro.arg];
        for (let edit of tweak.edits) {
          body = body.replace(edit.regex, edit.replace);
        }
        edited.push(macro.indent + ' use ' + macro.arg + ' [' + tweak.md5 + ']');
        break;
      case 'end':
        edited.push(body);
        edited.push(line);
        return;
    }
  }
}
