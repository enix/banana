import ActionsTypes from '../../constants/ActionsTypes';

const api = {
  pingApi: () => ({
    type: ActionsTypes.PING_API,
    payload: {},
  }),
  pingApiFailure: (error, ajax) => ({
    type: ActionsTypes.PING_API_FAILURE,
    payload: { error, ajax },
  }),

  listBuckets: () => ({
    type: ActionsTypes.LIST_BUCKETS,
    payload: {},
  }),
  listBucketsSuccess: (response, ajax) => ({
    type: ActionsTypes.LIST_BUCKETS_SUCCESS,
    payload: { response, ajax },
  }),
};

export default api;
