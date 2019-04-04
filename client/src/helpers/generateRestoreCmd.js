
function generateRestoreCmd(config) {
  let baseCmd = `bananagent restore ${config.name} ${config.time} ${config.target}`;

  if (config.only) {
    baseCmd += `/${config.only} --file-to-restore=${config.only}`;
  }

  return baseCmd;
}

export default generateRestoreCmd;
