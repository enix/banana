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

class AjaxError extends Error {
  constructor(ajax, ...args) {
    super(...args);
    this.ajax = ajax;
  }
}

async function fireAjax(config) {
  try {
    return await axios({
      ...defaultAjaxConfig,
      ...config,
      url: `https://api.banana.enix.io${config.uri}`,
    });
  }
  catch (error) {
    throw new AjaxError(config, error.message);
  }
}

export default fireAjax;
