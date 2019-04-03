import { put, takeLatest } from 'redux-saga/effects';

import ActionsCreators from '../actions';
import ActionsTypes from '../../constants/ActionsTypes';

function* pingApi() {
  yield put(ActionsCreators.fireAjax({
    url: 'http://localhost:8080/ping',
    onFailure: ActionsCreators.pingApiFailure,
  }));
}

function* listBuckets() {
  yield put(ActionsCreators.fireAjax({
    url: 'http://localhost:8080/buckets',
    onSuccess: ActionsCreators.listBucketsSuccess,
  }));
}

const sagas = function* () {
  yield takeLatest(ActionsTypes.PING_API, pingApi);
  yield takeLatest(ActionsTypes.LIST_BUCKETS, listBuckets);
};

export default sagas;
