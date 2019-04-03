import { fork, all } from 'redux-saga/effects';

import app from './app';
import api from './api';

const sagas = function* () {
  yield all([
    fork(app),
    fork(api),
  ]);
}

export default sagas;
