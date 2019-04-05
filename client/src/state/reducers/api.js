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
    containers: { $set: data }
  }),
  [ActionsTypes.LIST_BACKUPS_IN_CONTAINER_SUCCESS]: (state, { response: { data }, containerName }) => {
    let found = false;
    const updatedContainers = state.containers.map(c => {
      if (c.name === containerName) {
        found = true;
        return data;
      }

      return c;
    });

    if (!found) {
      updatedContainers.push(data);
    }

    return update(state, {
      containers: { $set: updatedContainers },
    });
  },
};

export default api;
