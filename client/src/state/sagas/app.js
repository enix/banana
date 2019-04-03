import { select, takeLatest } from 'redux-saga/effects';

import ActionsTypes from '../../constants/ActionsTypes';
import { getAppTitle } from '../selectors';

function* setupApp(action) {
  document.title = yield select(getAppTitle);
}

const sagas = function* () {
  yield takeLatest(ActionsTypes.SETUP_APP, setupApp);
};

export default sagas;
