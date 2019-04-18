import { put, takeLatest } from 'redux-saga/effects';

import ActionsCreators from '../actions';
import ActionsTypes from '../../constants/ActionsTypes';
import { fireAjax } from '../../helpers';

function* pingApi() {
  try {
    yield fireAjax({ uri: '/ping' });
  }
  catch (error) {
    yield put(ActionsCreators.pingApiFailure(error))
  };
}

function* listAgents() {
  try {
    const response = yield fireAjax({ uri: '/agents' });
    yield put(ActionsCreators.listAgentsSuccess({}, response))
  }
  catch (_) {}
}

function* getAgent({ payload: { org, cn } }) {
  try {
    const response = yield fireAjax({ uri: `/agents/${org}:${cn}` });
    yield put(ActionsCreators.getAgentSuccess({ org, cn }, response))
  }
  catch (_) {}
}

const sagas = function* () {
  yield takeLatest(ActionsTypes.PING_API, pingApi);
  yield takeLatest(ActionsTypes.LIST_AGENTS, listAgents);
  yield takeLatest(ActionsTypes.GET_AGENT, getAgent);
};

export default sagas;
