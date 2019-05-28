
function getTagColor(tag) {
  if (/start/gi.test(tag)) {
    return 'orange';
  }

  if (/done/gi.test(tag)) {
    return 'green';
  }

  if (/stopped|failed|crashed/gi.test(tag)) {
    return 'volcano';
  }

  return 'grey';
}

function getTypeTagColor(type) {
  if (/inc/gi.test(type)) {
    return 'blue';
  }

  if (/full/gi.test(type)) {
    return 'green';
  }

  return 'grey';
}

export { getTagColor, getTypeTagColor };
