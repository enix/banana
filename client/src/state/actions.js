import { notification } from 'antd';
import update from 'immutability-helper';

import { fireAjax } from '../helpers';

export default {
  setupApp: {
    *saga() {
      yield document.title = 'Banana';
    },
  },

  pingApi: {
    *saga() {
      return yield fireAjax({ uri: '/ping' });
    },
  },
  pingApiSuccess: {
    reducer: (state, { data }) => update(state, {
      user: {
        $set: {
          name: data.issuer,
          organization: data.organization,
        },
      },
    }),
  },

  listAgents: {
    *saga() {
      return yield fireAjax({ uri: '/agents' });
    },
  },
  listAgentsSuccess: {
    reducer(state, { data }) {
      return update(state, {
        agentList: { $set: data }
      });
    },
  },

  getAgent: {
    *saga(org, cn) {
      return yield fireAjax({ uri: `/agents/${org}:${cn}`, org, cn });
    },
  },
  getAgentSuccess: {
    reducer: (state, { data }, { payload: [ org, cn ] }) => update(state, {
      agents: {
        $set: {
          [org]: {
            [cn]: data,
          },
        },
      },
    }),
  },

  getAgentMessages: {
    *saga(org, cn) {
      return yield fireAjax({ uri: `/agents/${org}:${cn}/messages`, org, cn });
    },
  },
  getAgentMessagesSuccess: {
    reducer: (state, { data }, { payload: [ org, cn ] }) => update(state, {
      agentsMessages: {
        $set: {
          [org]: {
            [cn]: data,
          },
        },
      },
    }),
  },

  actionFailed: {
    *saga(error) {
      yield notification.error({
        message: error.message,
        description: `${error.ajax.method || 'GET'} ${error.ajax.uri}`,
      });
    },
  },
};
