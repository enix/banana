
function generateRestoreCmd({ command: { name, target }, timestamp}) {
  return `bananactl restore ${name} ${timestamp} ${target}`;
}

export default generateRestoreCmd;
