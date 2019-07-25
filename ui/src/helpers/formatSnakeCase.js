
function formatSnakeCase(string, fullUpperCase = false) {
  if (!string) {
    return '';
  }

  const strWithSpaces = string.replace(/_/g, ' ')

  if (fullUpperCase) {
    return strWithSpaces.toUpperCase();
  }

  return strWithSpaces[0].toUpperCase() + strWithSpaces.slice(1);
}

export default formatSnakeCase;
