import axios from 'axios';

const defaultAjaxConfig = {
  method: 'get',
  url: null,
  data: null,
  onSuccess: null,
  onFailure: null,
  headers: {
    'content-type': 'application/json',
    'accept': 'application/json',
  },
};

async function fireAjax(config) {
  return await axios({
    ...defaultAjaxConfig,
    ...config,
    url: `http://localhost:8080${config.uri}`
  });
}

export default fireAjax;
