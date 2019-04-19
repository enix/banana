import update from 'immutability-helper';

import ActionsTypes from '../../constants/ActionsTypes';
import { extendUpdate } from '../../helpers';

extendUpdate(update);

const api = {
  [ActionsTypes.PING_API_SUCCESS]: (state, { response: { data }}) => update(state, {
    user: {
      $set: {
        name: data.issuer,
        organization: data.organization,
      },
    },
  }),
  [ActionsTypes.PING_API_FAILURE]: state => update(state, {
    app: {
      isSetup: { $set: false },
    },
  }),
  [ActionsTypes.LIST_AGENTS_SUCCESS]: (state, { response: { data } }) => update(state, {
    agentList: { $set: data }
  }),
  [ActionsTypes.GET_AGENT_SUCCESS]: (state, { request: { org, cn }, response: { data } }) => update(state, {
    agents: {
      $set: {
        [org]: {
          [cn]: data
        }
      }
    },
  }),
  [ActionsTypes.GET_AGENT_MESSAGES_SUCCESS]: (state, { request: { org, cn }, response: { data } }) => update(state, {
    agentsMessages: {
      $set: {
        [org]: {
          [cn]: data
        }
      }
    },
  }),

};

export default api;
