
function trimSlash(str) {
  if (str[str.length - 1] === '/') {
    return str.slice(0, str.length - 1);
  }

  return str;
}

function generateRestoreCmd({ command: { name, target }, timestamp}) {
  if (!target) {
    return;
  }

  return `bananactl restore ${name} ${timestamp} ${trimSlash(target)}.bak`;
}

export default generateRestoreCmd;
