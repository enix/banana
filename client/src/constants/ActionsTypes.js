import keyMirror from 'key-mirror';

const ActionsTypes = keyMirror({
  SETUP_APP: null,

  PING_API: null,
  PING_API_FAILURE: null,

  LIST_AGENTS: null,
  LIST_AGENTS_SUCCESS: null,

  GET_AGENT: null,
  GET_AGENT_SUCCESS: null,
});

export default ActionsTypes;
