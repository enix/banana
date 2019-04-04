import ActionsTypes from '../../constants/ActionsTypes';

const api = {
  pingApi: () => ({
    type: ActionsTypes.PING_API,
    payload: {},
  }),
  pingApiFailure: (error) => ({
    type: ActionsTypes.PING_API_FAILURE,
    payload: { error },
  }),

  listBackupContainers: () => ({
    type: ActionsTypes.LIST_BACKUP_CONTAINERS,
    payload: {},
  }),
  listBackupContainersSuccess: (response) => ({
    type: ActionsTypes.LIST_BACKUP_CONTAINERS_SUCCESS,
    payload: { response },
  }),
};

export default api;
