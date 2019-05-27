
function getTagColor(tag) {
  if (/.*_start/.test(tag)) {
    return 'orange';
  }

  if (/.*_done/.test(tag)) {
    return 'green';
  }

  if (/(.*_failed)|(.*_crashed)/.test(tag)) {
    return 'volcano';
  }

  return 'grey';
}

function getTypeTagColor(type) {
  if (/(I|i)nc.*/.test(type)) {
    return 'blue';
  }

  if (/(F|f)ull/.test(type)) {
    return 'green';
  }

  return 'grey';
}

export { getTagColor, getTypeTagColor };
