import moment from 'moment';

function formatDate(timestamp) {
  return moment.unix(timestamp).format('MMMM Do YYYY, h:mm:ss a');
};

export default formatDate;
