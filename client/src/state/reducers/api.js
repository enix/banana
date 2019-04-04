import update from 'immutability-helper';

import ActionsTypes from '../../constants/ActionsTypes';
import { extendUpdate } from '../../helpers';

extendUpdate(update);

const api = {
  [ActionsTypes.PING_API_FAILURE]: state => update(state, {
    app: {
      isSetup: { $set: false },
    },
  }),
  [ActionsTypes.LIST_BACKUP_CONTAINERS_SUCCESS]: (state, { response: { data }}) => update(state, {
    buckets: { $set: data }
  }),
};

export default api;
