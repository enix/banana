import { select, put, takeLatest } from 'redux-saga/effects';
import axios from 'axios';

import ActionsCreators from '../actions';
import ActionsTypes from '../../constants/ActionsTypes';
import { getAppTitle } from '../selectors';

function* fireAjaxFailure({ payload: { error } }) {
  yield console.log(error);
}

function* fireAjaxSuccess({ payload: { response } }) {
  yield console.log(response);
}

function* fireAjax({ payload }) {
  Object.assign(payload.headers, {
    'content-type': 'application/json',
    'accept': 'application/json',
  });

  try {
    const response = yield axios(payload);

    yield put(ActionsCreators.fireAjaxSuccess(response, payload));
    if (typeof payload.onSuccess === 'function') {
      yield put(payload.onSuccess(response, payload));
    }
  } catch (e) {
    yield put(ActionsCreators.fireAjaxFailure(e, payload));
    if (typeof payload.onFailure === 'function') {
      yield put(payload.onFailure(e, payload));
    }
  }
}

function* setupApp() {
  document.title = yield select(getAppTitle);
}

const sagas = function* () {
  yield takeLatest(ActionsTypes.SETUP_APP, setupApp);
  yield takeLatest(ActionsTypes.FIRE_AJAX, fireAjax);
  yield takeLatest(ActionsTypes.FIRE_AJAX_SUCCESS, fireAjaxSuccess);
  yield takeLatest(ActionsTypes.FIRE_AJAX_FAILURE, fireAjaxFailure);
};

export default sagas;
