import { createSelector } from 'reselect';

export const getAppTitle = createSelector(
  state => state.app.title,
  title => title,
);
