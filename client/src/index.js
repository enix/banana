import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { createStore, applyMiddleware } from 'redux';
import createSagaMiddleware from 'redux-saga';

import App from './containers/App';
import storeApp from './state/reducers';
import sagas from './state/sagas';
import registerServiceWorker from './registerServiceWorker';

import './assets/vendor/bootstrap/css/bootstrap.css';
import './assets/vendor/font-awesome/css/fontawesome-all.min.css';
import './assets/css/front.css';

const sagaMiddleware = createSagaMiddleware();
const store = createStore(storeApp, applyMiddleware(sagaMiddleware));

sagaMiddleware.run(sagas);
ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);

registerServiceWorker();
