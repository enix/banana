import axios from 'axios';

const defaultAjaxConfig = {
  method: 'get',
  url: null,
  data: null,
  headers: {
    'content-type': 'application/json',
    'accept': 'application/json',
  },
};

async function fireAjax(config) {
  return await axios({
    ...defaultAjaxConfig,
    ...config,
    url: `https://api.banana.enix.io${config.uri}`
  });
}

export default fireAjax;
