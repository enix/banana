import keyMirror from 'key-mirror';

const ActionsTypes = keyMirror({
  SETUP_APP: null,

  FIRE_AJAX: null,
  FIRE_AJAX_SUCCESS: null,
  FIRE_AJAX_FAILURE: null,

  PING_API: null,
  PING_API_FAILURE: null,

  LIST_BUCKETS: null,
  LIST_BUCKETS_SUCCESS: null,
});

export default ActionsTypes;
