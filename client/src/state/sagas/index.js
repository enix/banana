import { fork, all } from 'redux-saga/effects';

import app from './app';

const sagas = function* () {
  yield all([
    fork(app),
  ]);
}

export default sagas;
