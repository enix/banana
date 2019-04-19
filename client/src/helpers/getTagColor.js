
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

export default getTagColor;
