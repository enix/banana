import keyMirror from 'key-mirror';

const ActionsTypes = keyMirror({
  SETUP_APP: null,

  PING_API: null,
  PING_API_FAILURE: null,

  LIST_BACKUP_CONTAINERS: null,
  LIST_BACKUP_CONTAINERS_SUCCESS: null,

  LIST_BACKUPS_IN_CONTAINER: null,
  LIST_BACKUPS_IN_CONTAINER_SUCCESS: null,
});

export default ActionsTypes;
