
function trimSlash(str) {
  if (str[str.length - 1] === '/') {
    return str.slice(0, str.length - 1);
  }

  return str;
}

function generateRestoreCmd({ command: { name, target, opaque_id } }) {
  if (!target) {
    return;
  }

  return `bananactl restore ${name} ${opaque_id} ${trimSlash(target)}.bak`;
}

export default generateRestoreCmd;
