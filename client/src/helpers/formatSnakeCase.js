
function formatSnakeCase(string, fullUpperCase = false) {
  const strWithSpaces = string.replace(/_/g, ' ')

  if (fullUpperCase) {
    return strWithSpaces.toUpperCase();
  }

  return strWithSpaces[0].toUpperCase() + strWithSpaces.slice(1);
}

export default formatSnakeCase;
