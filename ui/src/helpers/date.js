import moment from 'moment-timezone';

function formatDate(timestamp) {
  let tz = getSelectedTimezoneName();
  let date = moment(timestamp);

  if (tz === 'UTC') {
    date = date.utc();
  }
  else {
    tz = getLocalTimezoneAbbr();
  }

  return date.format(`MMMM Do YYYY, h:mm:ss a [${tz}]`);
};

function getSelectedTimezoneName() {
  const choice = localStorage.getItem('dateFormat');
  return choice === 'UTC' ? choice : getLocalTimezoneName();
}

function getLocalTimezoneName() {
  return moment.tz.guess();
}

function getLocalTimezoneAbbr() {
  return moment.tz(moment.tz.guess()).zoneAbbr();
}

export {
  formatDate,
  getSelectedTimezoneName,
  getLocalTimezoneName,
  getLocalTimezoneAbbr,
};
