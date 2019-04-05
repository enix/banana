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

  listTreesInContainer: (containerName) => ({
    type: ActionsTypes.LIST_TREES_IN_CONTAINER,
    payload: { containerName },
  }),
  listTreesInContainerSuccess: (containerName, response) => ({
    type: ActionsTypes.LIST_TREES_IN_CONTAINER_SUCCESS,
    payload: { containerName, response },
  }),

  listBackupsForTree: (containerName, treeName) => ({
    type: ActionsTypes.LIST_BACKUPS_FOR_TREE,
    payload: { containerName, treeName },
  }),
  listBackupsForTreeSuccess: (containerName, treeName, response) => ({
    type: ActionsTypes.LIST_BACKUPS_FOR_TREE_SUCCESS,
    payload: { containerName, treeName, response },
  }),
};

export default api;
