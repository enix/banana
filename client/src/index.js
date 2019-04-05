import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import createSagaMiddleware from 'redux-saga';

import App from './containers/App';
import storeApp from './state/reducers';
import sagas from './state/sagas';
import registerServiceWorker from './registerServiceWorker';
import { extendArray } from './helpers';

const sagaMiddleware = createSagaMiddleware();
const store = createStore(storeApp, applyMiddleware(sagaMiddleware));

extendArray();
sagaMiddleware.run(sagas);
ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);

registerServiceWorker();
