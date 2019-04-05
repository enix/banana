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
  [ActionsTypes.LIST_TREES_IN_CONTAINER_SUCCESS]: (state, { response: { data }, containerName }) => {
    const updatedContainers = state.containers.upsert(data, c => c.name === containerName);
    return update(state, {
      containers: { $set: updatedContainers },
    });
  },
  [ActionsTypes.LIST_BACKUPS_FOR_TREE_SUCCESS]: (state, { response: { data }, containerName, treeName }) => {
    let container = { ...state.containers.find(c => c.name === containerName) };

    if (!container || !container.contents) {
      container = {
        name: containerName,
        contents: [data]
      };
    }
    else {
      container.contents.upsert(data, t => t.name === treeName);
    }

    const updatedContainers = state.containers.upsert(container, c => c.name === containerName);
    return update(state, {
      containers: { $set: updatedContainers },
    });
  },
};

export default api;
