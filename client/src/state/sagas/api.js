import { put, takeLatest } from 'redux-saga/effects';

import ActionsCreators from '../actions';
import ActionsTypes from '../../constants/ActionsTypes';

function* pingApi() {
  yield put(ActionsCreators.fireAjax({
    url: 'http://localhost:8080/ping',
    onFailure: ActionsCreators.pingApiFailure,
  }));
}

const sagas = function* () {
  yield takeLatest(ActionsTypes.PING_API, pingApi);
};

export default sagas;
