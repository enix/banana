import update from 'immutability-helper';

import extendUpdate from '../../helpers/extendUpdate';
import ActionsTypes from '../../constants/ActionsTypes';

extendUpdate(update);

const api = {
  [ActionsTypes.PING_API_FAILURE]: state => update(state, {
    app: {
      isSetup: { $set: false },
    },
  }),
  [ActionsTypes.LIST_BUCKETS_SUCCESS]: (state, { response: { data }}) => update(state, {
    buckets: { $set: data }
  }),
};

export default api;
