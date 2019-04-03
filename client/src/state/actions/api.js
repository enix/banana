import ActionsTypes from '../../constants/ActionsTypes';

const api = {
  pingApi: () => ({
    type: ActionsTypes.PING_API,
    payload: {},
  }),
  pingApiFailure: (error, ajax) => ({
    type: ActionsTypes.PING_API_FAILURE,
    payload: { error },
  }),
};

export default api;
