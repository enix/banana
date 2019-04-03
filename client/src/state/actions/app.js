import ActionsTypes from '../../constants/ActionsTypes';

const app = {
  setupApp: () => ({
    type: ActionsTypes.SETUP_APP,
    payload: {},
  }),
  fireAjax: ({ url, method = 'get', headers = {}, data = null, onSuccess = null, onFailure = null }) => ({
    type: ActionsTypes.FIRE_AJAX,
    payload: { url, method, headers, data, onSuccess, onFailure },
  }),
  fireAjaxSuccess: (response, ajax) => ({
    type: ActionsTypes.FIRE_AJAX_SUCCESS,
    payload: { response, ajax },
  }),
  fireAjaxFailure: (error, ajax) => ({
    type: ActionsTypes.FIRE_AJAX_FAILURE,
    payload: { error, ajax },
  }),
};

export default app;
