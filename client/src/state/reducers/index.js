import app from './app';
import api from './api';
import initialState from '../initialState';

const reducersMap = {
  ...app,
  ...api,
  leaveStateUnchanged: state => state,
};

const reducers = function reducers(state = initialState, action) {
  const reducer = reducersMap[action.type] || reducersMap.leaveStateUnchanged;
  const newState = reducer(state, action.payload, action.meta);
  return newState;
};

export default reducers;
