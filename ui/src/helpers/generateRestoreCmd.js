
function trimSlash(str) {
  if (str[str.length - 1] === '/') {
    return str.slice(0, str.length - 1);
  }

  return str;
}

function generateRestoreCmd() {
  return 'restore helper not implemented yet, refer to doc to restore';
  // return `bananagent restore ${name} ${opaque_id} ${trimSlash(target)}.bak`;
}

export default generateRestoreCmd;
