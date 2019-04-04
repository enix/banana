import { put, takeLatest } from 'redux-saga/effects';

import ActionsCreators from '../actions';
import ActionsTypes from '../../constants/ActionsTypes';
import { fireAjax } from '../../helpers';

function* pingApi() {
  try {
    yield fireAjax({
      url: 'http://localhost:8080/ping',
      onFailure: ActionsCreators.pingApiFailure,
    });
  }
  catch (error) {
    yield put(ActionsCreators.pingApiFailure(error))
  };
}

function* listBackupContainers() {
  try {
    const response = yield fireAjax({ uri: '/backups' });
    yield put(ActionsCreators.listBackupContainersSuccess(response))
  }
  catch (_) {}
}

const sagas = function* () {
  yield takeLatest(ActionsTypes.PING_API, pingApi);
  yield takeLatest(ActionsTypes.LIST_BACKUP_CONTAINERS, listBackupContainers);
};

export default sagas;
