
function generateRestoreCmd({ name, time, target, only }) {
  let baseCmd = `./bananagent restore ${name} ${time} ${target}`;

  if (only) {
    baseCmd += `/${only} --file-to-restore=${only}`;
  }

  return baseCmd;
}

export default generateRestoreCmd;
