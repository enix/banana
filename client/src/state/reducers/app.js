import update from 'immutability-helper';

import ActionsTypes from '../../constants/ActionsTypes';
import { extendUpdate } from '../../helpers';

extendUpdate(update);

const app = {
  [ActionsTypes.SETUP_APP]: state => update(state, {
    app: {
      isSetup: { $set: true },
    },
  }),
};

export default app;
