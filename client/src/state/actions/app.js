import ActionsTypes from '../../constants/ActionsTypes';

const app = {
  setupApp: () => ({
    type: ActionsTypes.SETUP_APP,
    payload: {},
  }),
};

export default app;
