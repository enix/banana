import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { createStore } from 'redux-saga-wrapper';

import initialState from './state/initialState';
import actions from './state/actions';
import App from './containers/App';
import { extendArray } from './helpers';
import registerServiceWorker from './registerServiceWorker';

const store = createStore(initialState, actions);

extendArray();
ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);

registerServiceWorker();
