import { put, takeLatest } from 'redux-saga/effects';
import { notification } from 'antd';

import ActionsCreators from '../actions';
import ActionsTypes from '../../constants/ActionsTypes';
import { fireAjax } from '../../helpers';

function* pingApi() {
  const request = { uri: '/ping' };

  try {
    const response = yield fireAjax(request);
    yield put(ActionsCreators.pingApiSuccess(request, response))
  }
  catch (error) {
    yield put(ActionsCreators.pingApiFailure(error))
    yield put(ActionsCreators.ajaxFailure(request, error))
  };
}

function* listAgents() {
  const request = { uri: '/agents' };

  try {
    const response = yield fireAjax(request);
    yield put(ActionsCreators.listAgentsSuccess(request, response))
  }
  catch (error) {
    yield put(ActionsCreators.ajaxFailure(request, error))
  }
}

function* getAgent({ payload: { org, cn } }) {
  const request = { uri: `/agents/${org}:${cn}`, org, cn };
  
  try {
    const response = yield fireAjax(request);
    yield put(ActionsCreators.getAgentSuccess(request, response))
  }
  catch (error) {
    yield put(ActionsCreators.ajaxFailure(request, error))
  }
}

function* getAgentMessages({ payload: { org, cn } }) {
  const request = { uri: `/agents/${org}:${cn}/messages`, org, cn };

  try {
    const response = yield fireAjax(request);
    yield put(ActionsCreators.getAgentMessagesSuccess(request, response))
  }
  catch (error) {
    yield put(ActionsCreators.ajaxFailure(request, error))
  }
}

function ajaxDidFail({ payload: { request, error } }) {
  notification.error({
    message: error.message,
    description: `${request.method || 'GET'} ${request.uri}`,
  });
}

const sagas = function* () {
  yield takeLatest(ActionsTypes.PING_API, pingApi);
  yield takeLatest(ActionsTypes.LIST_AGENTS, listAgents);
  yield takeLatest(ActionsTypes.GET_AGENT, getAgent);
  yield takeLatest(ActionsTypes.GET_AGENT_MESSAGES, getAgentMessages);
  yield takeLatest(ActionsTypes.AJAX_FAILURE, ajaxDidFail);
};

export default sagas;
