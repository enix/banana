import moment from 'moment';

function formatDate(timestamp) {
  let date = moment(timestamp);

  if (localStorage.getItem('dateFormat') === 'UTC') {
    date = date.utc();
  }

  return date.format('MMMM Do YYYY, h:mm:ss a');
};

export default formatDate;
