import { put, takeLatest } from 'redux-saga/effects';

import ActionsCreators from '../actions';
import ActionsTypes from '../../constants/ActionsTypes';
import { fireAjax } from '../../helpers';

function* pingApi() {
  try {
    yield fireAjax({ url: 'http://localhost:8080/ping' });
  }
  catch (error) {
    yield put(ActionsCreators.pingApiFailure(error))
  };
}

function* listBackupContainers() {
  try {
    const response = yield fireAjax({ uri: '/containers' });
    yield put(ActionsCreators.listBackupContainersSuccess(response))
  }
  catch (_) {}
}

function* listTreesInContainer({ payload: { containerName } }) {
  try {
    const response = yield fireAjax({ uri: `/containers/${containerName}` });
    yield put(ActionsCreators.listTreesInContainerSuccess(containerName, response))
  }
  catch (_) {}
}

function* listBackupsForTree({ payload: { containerName, treeName } }) {
  try {
    const response = yield fireAjax({ uri: `/containers/${containerName}/tree/${treeName}` });
    yield put(ActionsCreators.listBackupsForTreeSuccess(containerName, treeName, response))
  }
  catch (_) {}
}

const sagas = function* () {
  yield takeLatest(ActionsTypes.PING_API, pingApi);
  yield takeLatest(ActionsTypes.LIST_BACKUP_CONTAINERS, listBackupContainers);
  yield takeLatest(ActionsTypes.LIST_TREES_IN_CONTAINER, listTreesInContainer);
  yield takeLatest(ActionsTypes.LIST_BACKUPS_FOR_TREE, listBackupsForTree);
};

export default sagas;
